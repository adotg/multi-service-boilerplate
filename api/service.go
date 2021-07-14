package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIResponse struct {
	LastAccessTime int64  `json:"lastAccessTime"`
	Data           string `json:"data"`
}

type Service struct {
	Router *mux.Router
	Server *http.Server
	cache  *Cache
}

func (serv *Service) health(w http.ResponseWriter, r *http.Request) {
	L.Info("/health is called and responded with ok")
	health := "ok"

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(health))
}

func (serv *Service) getLastAccessTime(userKey string) int64 {
	var lastAccessTime int64 = -1
	if !isEmpty(userKey) {
		timeStr := serv.cache.getAccessInf(userKey)
		if !isEmpty(timeStr) {
			if i, err := strconv.ParseInt(timeStr, 10, 32); err == nil {
				lastAccessTime = i
			} else {
				L.Errorf("Error while converting cached value %s", err.Error())
			}
		}
	}

	return lastAccessTime
}

func (serv *Service) updateLastAccessTime(userKey string) {
	if !isEmpty(userKey) {
		serv.cache.setAccessInf(userKey, strconv.FormatInt(time.Now().Unix(), 10))
	}
}

func (serv *Service) get(w http.ResponseWriter, r *http.Request) {
	env := GetEnvVars()
	pathParams := mux.Vars(r)
	key := pathParams["key"]
	userKey := r.Header.Get("user_key")
	L.Infof("/get is called with %s for user %s", key, userKey)

	// call data service from here
	res, err := http.Get("http://" + env.DataServiceHost + ":" + env.DataServicePort + "/key/" + key)
	if err != nil {
		L.Panicf("Not able to do GET to data service. Error=%s", err.Error())
	}
	resFromDataService, err := ioutil.ReadAll(res.Body)
	if err != nil {
		L.Panicf("Not able to do parse response from data service. Error=%s", err.Error())
	}

	resp := APIResponse{serv.getLastAccessTime(userKey), string(resFromDataService)}
	jsonStr, err := json.Marshal(resp)
	if err != nil {
		L.Errorf("/get error %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	serv.updateLastAccessTime(userKey)
	L.Infof("/get responds with %s", jsonStr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}

func (serv *Service) set(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	key := pathParams["key"]
	userKey := r.Header.Get("user_key")
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		L.Errorf("/set error while reading request body %s", err.Error())
	}
	L.Infof("/set is called with %s with value %s for user %s", key, body, userKey)

	// call the dataservice api here
	res, err := http.Post("http://"+env.DataServiceHost+":"+env.DataServicePort+"/key/"+key, "text/plain",
		bytes.NewBuffer(body))
	if err != nil {
		L.Panicf("Not able to do GET to data service. Error=%s", err.Error())
	}
	resFromDataService, err := ioutil.ReadAll(res.Body)
	if err != nil {
		L.Panicf("Not able to do parse response from data service. Error=%s", err.Error())
	}

	resp := APIResponse{serv.getLastAccessTime(userKey), string(resFromDataService)}
	jsonStr, err := json.Marshal(resp)
	if err != nil {
		L.Errorf("/get error %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	serv.updateLastAccessTime(userKey)
	L.Infof("/set responds with %s", jsonStr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}

func (serv *Service) initRoutes() {
	serv.Router.HandleFunc("/health", serv.health).Methods(http.MethodGet)
	serv.Router.HandleFunc("/get/{key}", serv.get).Methods(http.MethodGet)
	serv.Router.HandleFunc("/set/{key}", serv.set).Methods(http.MethodPost)
}

func (serv *Service) Init() {
	cache := Cache{}
	cache.Init()
	serv.cache = &cache

	serv.Router = mux.NewRouter()
	serv.initRoutes()
	L.Info("Service inited")
}

func (serv *Service) Run() {
	env := GetEnvVars()
	go func() {

		headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "user_key"})
		originsOK := handlers.AllowedOrigins([]string{"*"})
		methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

		serv.Server = &http.Server{
			Addr:    ":" + env.ServerPort,
			Handler: handlers.CORS(headersOK, originsOK, methodsOK)(serv.Router),
		}
		L.Infof("Listening to :%s", env.ServerPort)

		if err := serv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			L.Panicf("Error while starting the server: %s\n", err)
		}
	}()
}

func (serv *Service) Shutdown(ctx context.Context) {
	serv.cache.shutdown(ctx)
	if err := serv.Server.Shutdown(ctx); err != nil {
		L.Errorf("Server Shutdown Failed:%+v \n", err)
	}
	L.Info("Service is stopped")
}

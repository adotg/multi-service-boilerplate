package main

import (
	"log"
	"os"
	"sync"
)

type EnvVar struct {
	DataServiceHost string
	DataServicePort string
	RedisHost       string
	RedisPort       string
	ServerPort      string
}

var env *EnvVar
var once sync.Once

func isEmpty(str string) bool {
	return str == ""
}

func checkIfValidVar(varName string, varValue string) {
	if isEmpty(varValue) {
		log.Panicf("Env var %s value is empty\n", varName)
	}
}

func InitEnvVars() {
	envVarname := "DATA_SERVICE_HOST"
	dataServiceHost := os.Getenv(envVarname)
	checkIfValidVar(envVarname, dataServiceHost)

	envVarname = "DATA_SERVICE_PORT"
	dataServicePort := os.Getenv(envVarname)
	checkIfValidVar(envVarname, dataServicePort)

	envVarname = "REDIS_HOST"
	redisHost := os.Getenv(envVarname)
	checkIfValidVar(envVarname, redisHost)

	envVarname = "REDIS_PORT"
	redisPort := os.Getenv(envVarname)
	checkIfValidVar(envVarname, redisPort)

	envVarname = "SERVER_PORT"
	serverPort := os.Getenv(envVarname)
	checkIfValidVar(envVarname, serverPort)

	once.Do(func() {
		env = &EnvVar{
			DataServiceHost: dataServiceHost,
			DataServicePort: dataServicePort,
			RedisHost:       redisHost,
			RedisPort:       redisPort,
			ServerPort:      serverPort,
		}
	})
}

func GetEnvVars() *EnvVar {
	return env
}

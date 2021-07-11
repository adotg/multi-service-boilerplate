package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

var L = log.New()

func main() {
	InitEnvVars()

	L.SetLevel(log.DebugLevel)
	L.SetReportCaller(true)
	L.SetFormatter(GetLogFormatter())
	L.Info("Starting the API Service")

	serv := Service{}
	serv.Init()

	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	serv.Run()

	<-stopSig
	L.Warn("Server will be stopped gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	serv.Shutdown(ctx)
}

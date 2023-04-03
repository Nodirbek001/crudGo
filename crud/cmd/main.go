package main

import (
	"context"
	"crut/crud/api/controller"
	"crut/crud/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load()

	roor:=mux.NewRouter()
	roor.HandleFunc("/user", controller.Sign)
	controller.Init(roor)
	log.Println("main: Project is started on port: ", cfg.HTTPPort)
	errChan:= make(chan error, 1)
	osSignals:=make(chan os.Signal,1)
	signal.Notify(osSignals,os.Interrupt, syscall.SIGTERM)

	httpServer:=http.Server{
		Addr: cfg.HTTPPort,
		Handler: roor,
	}

	go func() {
		errChan<- httpServer.ListenAndServe()
	}()
	select {
	case err := <-errChan:
		log.Println(err)
	case <-osSignals:
		log.Println("main: recieved os signal, shutting down")
		_ = httpServer.Shutdown(context.Background())
		return
	}
}

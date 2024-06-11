package main

import (
	"lib"
	"log"
	"os"
	"os/signal"
	"parking-service/cmd/rest"
	"parking-service/internal/repository"
	"parking-service/internal/usecase"
	"syscall"
)

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")

	repository := repository.NewRepository(cfg.DataFile)
	service := usecase.NewService(repository)

	//setup rest server
	restChan := make(chan error, 1)
	go func() {
		restHandler := rest.NewHandler(service)
		restChan <- rest.RunServer(cfg, restHandler)
	}()

	interruption := make(chan os.Signal)
	go func() {
		signal.Notify(interruption, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	}()

	<-interruption

	select {
	case <-interruption:
		log.Println("interrupted")
	case err := <-restChan:
		log.Println("rest ran with an error", err)
	}
}

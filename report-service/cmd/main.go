package main

import (
	"github.com/opentracing/opentracing-go"
	"lib"
	"log"
	"os"
	"os/signal"
	"report-service/cmd/rest"
	"report-service/internal/repository"
	"report-service/internal/usecase"
	"syscall"
)

func main() {
	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")

	tracer, closer, err := lib.InitJaeger(cfg.App.Name)
	if err != nil {
		log.Fatalf("could not initialize jaeger tracer: %v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

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

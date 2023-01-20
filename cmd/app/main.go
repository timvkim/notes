package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timvkim/notes/internal/repository"
	"github.com/timvkim/notes/internal/service"
	"github.com/timvkim/notes/internal/transport/handlers"
)

func main() {
	log.Println("starting server ...")

	repository := repository.NewRepo()
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)

	server := &http.Server{
		Addr:    ":8090",
		Handler: handler.InitRouters(),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}

}

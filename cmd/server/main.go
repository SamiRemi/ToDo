package main

import (
	"log"
	"net/http"
	"todo/internal/handler"
	"todo/internal/repository"
	"todo/internal/service"
)

func main() {
	repo := repository.NewInMemoryRepo()
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

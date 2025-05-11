package main

import (
	"fmt"
	"log/slog"
	"myFirstGoProject/EncurtadorDeURL/api"
	"net/http"
	"time"
)


func main(){
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	fmt.Println("Step 0")

	db := make(map[string]string)

	handler := api.NewHandler(db)
	
	fmt.Println("Step 1")
	
	s := http.Server{
		ReadTimeout: 10 * time.Second,
		IdleTimeout: time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: handler,
	}
	
	fmt.Println("Step 2")
	
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	
	fmt.Println("Step 3")
	
	return nil	
}
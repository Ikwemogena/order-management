package main

import (
	"log"
	"github.com/ikwemogena/order-management/gateway/config"
	"github.com/ikwemogena/order-management/gateway/routes"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := routes.SetupRouter()
	
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
package main

import (
	"log"

	"github.com/Ikwemogena/order-management/config"
	"github.com/Ikwemogena/order-management/db"
	"github.com/Ikwemogena/order-management/routes"
)

func main() {
	log.Println("Starting server...")

	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log.Printf("Config loaded")

	db, err := db.Init(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	log.Println("Database initialized")

	r := routes.SetupRouter(db)

	log.Printf("Server started at %s", cfg.ServerAddress)

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
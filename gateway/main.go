package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ikwemogena/order-management-gateway/config"
)

func main() {
    fmt.Println("hello morgz")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Order Gateway bro!",
		})
	})
	
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ikwemogena/order-management-gateway/config"
	pb "github.com/ikwemogena/order-management/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	orderServiceAddr = "localhost:3000"
)

func main() {
    fmt.Println("hello morgz")

	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	defer conn.Close()

	log.Println("Successfully connected to order service", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	log.Println("Successfully created client", c)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OMS Gateway!",
		})
	})
	
	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
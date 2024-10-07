package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ikwemogena/order-management/gateway/models"
	pb "github.com/ikwemogena/order-management/stock/proto"
)

type StockHandler struct{}

func (h *StockHandler) connGrpc() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:8091",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, err
}

func (h *StockHandler) CreateStock(c *gin.Context){
	var stock models.Stock

	conn, err := h.connGrpc()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to stock service",
		})
		return
	}

	defer conn.Close()

	client := pb.NewStockServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	if err := c.ShouldBindJSON(&stock); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
        return
    }

	_, err = client.CreateStock(ctx, &pb.CreateStockRequest{
		ItemName:    stock.ItemName,
		ItemDescription: stock.Description,
		Quantity:    int32(stock.Quantity),
	})

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to add item to stock",
		})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

func (h *StockHandler) CheckStock(c *gin.Context){
	conn, err := h.connGrpc()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to stock service",
		})
		return
	}

	defer conn.Close()

	client := pb.NewStockServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

    itemID := c.Param("id")

	if itemID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Item ID is required",
        })
        return
    }

	r, err := client.CheckStock(ctx, &pb.CheckStockRequest{ItemId: itemID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": r.GetAvailable(),
	})
}
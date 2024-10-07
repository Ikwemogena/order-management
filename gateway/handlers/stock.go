package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ikwemogena/order-management/stock/proto"
)

type StockHandler struct{}

func (stock *StockHandler) connGrpc() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:8091",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, err
}

func CheckStock(c *gin.Context){
	stock := &StockHandler{}
	conn, err := stock.connGrpc()

	if err != nil {
		c.JSON(500, gin.H{
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
        c.JSON(400, gin.H{
            "error": "Item ID is required",
        })
        return
    }

	r, err := client.CheckStock(ctx, &pb.CheckStockRequest{ItemId: itemID})

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Item not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": r.GetAvailable(),
	})
}
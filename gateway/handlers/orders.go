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
	pb "github.com/ikwemogena/order-management/orders/proto"
)

type OrderHandler struct {}

func (order *OrderHandler) connGrpc() (*grpc.ClientConn, error) {

   conn, err := grpc.NewClient("orders:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))

   if err != nil {
        return nil, err
    }

	return conn, err
}

func(h *OrderHandler) CreateOrder(c *gin.Context){
	var order models.Order
	conn, err := h.connGrpc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Order Service Down!",
		})
		return
	}

	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
        return
    }

	var pbProducts []*pb.OrderItem
	for _, product := range order.Products {
		pbProducts = append(pbProducts, &pb.OrderItem{
			Id: product.ProductId,
			Quantity:  product.Quantity,
			UnitPrice: product.UnitPrice,
			TotalAmount:     product.TotalAmount,
		})
	}

	r, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
		Products: pbProducts,
		UserId: order.UserId,
		Amount: order.TotalAmount,
		ShippingAddress: order.ShippingAddress,
		PaymentMethod: order.PaymentMethod,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create order. Please try again.",
		})
		log.Printf("Failed to create order: %v", err)
		return
	}

	order.ID = r.GetOrderId()

	c.JSON(http.StatusCreated, order)
}
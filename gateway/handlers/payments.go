package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ikwemogena/order-management/gateway/models"
	pb "github.com/ikwemogena/order-management/payments/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentHandler struct{}

func (h *PaymentHandler) connGrpc() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:8094", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return conn, err
}

func (h *PaymentHandler) InitializePayment(c *gin.Context) {
	var payment models.Payment

	conn, err := h.connGrpc()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Payment Service Down!",
		})
		return
	}

	defer conn.Close()

	if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
        return
    }

	client := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := client.InitializePayment(ctx, &pb.InitializePaymentRequest{
		Provider:  payment.Provider,
		Email:     payment.Email,
		Amount:    payment.Amount,
		Currency:  payment.Currency,
	})

	log.Println(payment.Amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to initialize transaction. Please try again.",
		})
		log.Printf("Failed to initialize transaction: %v", err)
		return
	}

	c.JSON(http.StatusOK, r)
}

func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	paymentReference := c.Param("reference")

	if paymentReference == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Reference is required",
		})
	}

	conn, err := h.connGrpc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Payment Service Down!",
		})
		return
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := client.VerifyPayment(ctx, &pb.VerifyPaymentRequest{
		Reference:paymentReference,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to verify transaction. Please try again.",
		})
		log.Printf("Failed to verify transaction: %v", err)
		return
	}

	c.JSON(http.StatusOK, r)
}
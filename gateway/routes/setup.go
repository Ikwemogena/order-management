package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ikwemogena/order-management/gateway/handlers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	orderHandler := &handler.OrderHandler{}
	order := r.Group("/order")
	order.POST("/", orderHandler.CreateOrder)

	paymentHandler := &handler.PaymentHandler{}
	payment := r.Group("/payment")
	payment.POST("/initialize", paymentHandler.InitializePayment)
	payment.GET("/verify/:reference", paymentHandler.VerifyPayment)

	stockHandler := &handler.StockHandler{}
	stock := r.Group("/stock")
	stock.POST("/", stockHandler.CreateStock)
	stock.GET("/:id", stockHandler.CheckStock)

	return r
}
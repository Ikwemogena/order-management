package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ikwemogena/order-management/gateway/handlers"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	stock := r.Group("/stock")
	stock.GET("/:id", handler.CheckStock)

	return r
}
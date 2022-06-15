package handler

import (
	"github.com/duramash/constanta-emulator-task/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		transactions := api.Group("/transactions")
		{
			transactions.POST("/", h.createTransaction)
			transactions.PATCH("/:id", h.changeStatusOfTransactionById)
			transactions.GET("/:id", h.getStatusOfTransactionById)
			transactions.DELETE("/:id", h.cancelTransactionById)
			transactions.GET("/users/:user_attr", h.getTransactionsByUser)
		}
	}

	return router
}

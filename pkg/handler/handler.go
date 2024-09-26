package handler

import (
	"app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{service: srv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createItem)
			lists.GET("/", h.getAllItems)
			lists.GET("/:id", h.getItemById)
			lists.PUT("/:id", h.updateItemById)
			lists.DELETE("/:id", h.deleteItemById)
		}

		items := api.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItemById)
			items.DELETE("/:item_id", h.deleteItemById)
		}

	}
	return router
}

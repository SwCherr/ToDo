package handler

import (
	"app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		// auth.POST("/refresh", h.refreshToken)
	}

	// api := router.Group("/api", h.userIddentity)
	// {
	// 	// lists := api.Group("/lists")
	// 	// {
	// 	// 	lists.POST("/", h.createList)
	// 	// 	lists.GET("/", h.getAllLists)
	// 	// 	lists.GET("/:id", h.getListById)
	// 	// 	lists.PUT("/:id", h.updateListById)
	// 	// 	lists.DELETE("/:id", h.deleteListById)
	// 	// }
	// }
	return router
}

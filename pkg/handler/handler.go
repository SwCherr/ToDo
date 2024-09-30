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
		auth.POST("/up", h.signUp)
		// auth.POST("/in", h.signIn)
		auth.POST("/get", h.GetPareTokens)
		auth.POST("/refresh", h.refreshToken)
	}

	// api := router.Group("/api", h.userIddentity) {}
	return router
}

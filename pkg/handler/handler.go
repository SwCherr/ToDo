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
	auth := router.Group("/auth") // add check access autherization
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/get", h.GetPareTokens)
		auth.POST("/refresh", h.RefreshToken)
	}
	return router
}

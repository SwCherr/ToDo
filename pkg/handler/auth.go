package handler

import (
	todo "app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 6 video 3-47 | http.StatusBadRequest == 400 - ошибка со стороны клиента при запросе
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 6 video 3-47 | http.StatusBadRequest == 400 - ошибка со стороны клиента при запросе
		return
	}

	// отдаем клиенту
	newAccessToken, err := h.service.Authorization.GenerateAccessToken(input.Username, input.Password, c.ClientIP())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
		return
	}

	// сохраняем в бд в виде хэша
	newRefreshToken, err := h.service.Authorization.GenerateRefreshToken()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
		return
	}

	if err := h.service.CreateSession(input.Username, input.Password, newAccessToken, c.ClientIP()); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}

// type refreshToken struct {
// 	Token string `json:"token" binding:"required"`
// }

// func (h *Handler) refreshingToken(c *gin.Context) {
// 	var refreshTokenInpit refreshToken

// 	if err := c.BindJSON(&refreshTokenInpit); err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// }

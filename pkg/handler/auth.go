package handler

import (
	"app"
	b64 "encoding/base64"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input app.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.SignUp(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetPareTokens(c *gin.Context) {
	var input app.Sesion
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserIP = c.ClientIP()
	newAccessToken, newRefreshToken, err := h.service.GetPareToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// предача клиенту в формате base64
	newRefreshToken = b64.StdEncoding.EncodeToString([]byte(newRefreshToken))
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var input app.Sesion
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// декодирование RefreshToken из base64
	token, err := b64.StdEncoding.DecodeString(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.UserIP = c.ClientIP()
	input.RefreshToken = string(token)

	newAccessToken, newRefreshToken, err := h.service.RefreshToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// предача newRefreshToken в формате base64
	newRefreshToken = b64.StdEncoding.EncodeToString([]byte(newRefreshToken))
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}

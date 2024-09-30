package handler

import (
	b64 "encoding/base64"

	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	GUID int `json:"id" binding:"required"`
}

func (h *Handler) GetPareTokens(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.service.GeneratePareTokens(input.GUID, c.ClientIP())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// предача в формате base64
	newRefreshToken = b64.StdEncoding.EncodeToString([]byte(newRefreshToken))
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
		"user.UserIP":  c.ClientIP(),
	})
}

type refreshInput struct {
	GUID  int    `json:"id" binding:"required"`
	IP    string `json:"ip" binding:"required"`
	Token string `json:"token" binding:"required"`
}

func (h *Handler) refreshToken(c *gin.Context) {
	var input refreshInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// декодирование RefreshToken из формата base64
	token, err := b64.StdEncoding.DecodeString(input.Token)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.service.RefreshToken(input.GUID, c.ClientIP(), string(token))
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

package handler

import (
	"app"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input app.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.CreateUser(input)
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
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.UserIP = c.ClientIP()
	newAccessToken, newRefreshToken, err := h.service.GetPareToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var input app.Sesion
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.UserIP = c.ClientIP()
	newAccessToken, newRefreshToken, err := h.service.RefreshToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}

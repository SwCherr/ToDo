package handler

import (
	"app"
	b64 "encoding/base64"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input app.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 - ошибка со стороны клиента при запросе
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
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
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

// type signInInput struct {
// 	ID    int    `json:"id" binding:"required"`
// 	IP    string `json:"ip" binding:"required"`
// }

// func (h *Handler) signIn(c *gin.Context) {
// 	var input signInInput
// 	if err := c.BindJSON(&input); err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 - ошибка со стороны клиента при запросе
// 		return
// 	}
// 	fmt.Println(input)

// 	user, err := h.service.GetUserById(input.ID)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error()) //  500 - внутрення ошибка на сервере
// 		return
// 	}
// 	user.UserIP = c.ClientIP()

// 	newAccessToken, newRefreshToken, err := h.service.GeneratePareTokens(user)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
// 		return
// 	}

// 	// предача в формате base64
// 	newRefreshToken = b64.StdEncoding.EncodeToString([]byte(newRefreshToken))
// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"accessToken":  newAccessToken,
// 		"refreshToken": newRefreshToken,
// 	})
// }

type refreshInput struct {
	ID    int    `json:"id" binding:"required"`
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

	// newAccessToken, newRefreshToken, err := h.service.RefreshToken(input.ID, input.IP, string(token))
	newAccessToken, newRefreshToken, err := h.service.RefreshToken(input.ID, c.ClientIP(), string(token))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //  500 - внутрення ошибка на сервере
		return
	}

	// предача newRefreshToken в формате base64
	newRefreshToken = b64.StdEncoding.EncodeToString([]byte(newRefreshToken))
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}

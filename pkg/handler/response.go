package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct { // where error ?!!!!!
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) { // 6 video 3-47
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}

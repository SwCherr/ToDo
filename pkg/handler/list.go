package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userContext)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
}

func (h *Handler) getListById(c *gin.Context) {
}

func (h *Handler) updateListById(c *gin.Context) {
}

func (h *Handler) deleteListById(c *gin.Context) {
}

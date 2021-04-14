package handler

import (
	"net/http"

	GoArchitecture "github.com/Askaell/homework"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var input GoArchitecture.Item
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	newItem, err := h.services.Item.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          newItem.Id,
		"name":        newItem.Name,
		"description": newItem.Description,
		"price":       newItem.Price,
	})
}

func (h *Handler) getAllItem(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}

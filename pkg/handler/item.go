package handler

import (
	"net/http"
	"strconv"

	"github.com/Askaell/homework/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var input models.Item
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	newItem, err := h.repository.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(c, http.StatusOK, newItem)
}

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.repository.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(c, http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.repository.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(c, http.StatusOK, item)
}

func (h *Handler) deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.repository.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(c, http.StatusNoContent, nil)
}

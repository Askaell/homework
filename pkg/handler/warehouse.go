package handler

import (
	"net/http"
	"strconv"

	"github.com/Askaell/homework/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getWarehouseAllItems(c *gin.Context) {
	items, err := h.repository.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range items {
		items[i] = convertIntAmountToExpression(items[i])
	}

	writeResponse(c, http.StatusOK, items)
}

func (h *Handler) getWarehouseItemById(c *gin.Context) {
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

	writeResponse(c, http.StatusOK, convertIntAmountToExpression(item))
}

func (h *Handler) updateWarehouseItems(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	inputItem, err := h.repository.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	amount, err := strconv.Atoi(inputItem.Amount)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	summand, err := strconv.Atoi(c.Query("value"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	amount = amount + summand
	if amount < 0 {
		amount = 0
	}

	inputItem.Amount = strconv.Itoa(amount)

	if err := h.repository.Update(id, inputItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(c, http.StatusOK, nil)
}

// make convert item.Amount field from int to warehouse expression
func convertIntAmountToExpression(item models.Item) models.Item {
	amount, err := strconv.Atoi(item.Amount)
	if err != nil {
		item.Amount = ""
	}

	switch {
	case amount == 0:
		item.Amount = "Нет на складе"
	case amount > 0 && amount < 10:
		item.Amount = "На складе малое количество"
	case amount > 10:
		item.Amount = "На складе много"
	default:
		item.Amount = ""
	}

	return item
}

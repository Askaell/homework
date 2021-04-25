package handler

import (
	"github.com/Askaell/homework/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository repository.ItemRepository
}

func NewHandler(repository repository.ItemRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		items := api.Group("/items")
		{
			items.POST("", h.createItem)
			items.PUT("", h.updateItem)
			items.GET("", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}

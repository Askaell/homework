package service

import (
	"github.com/Askaell/homework/pkg/models"
	"github.com/Askaell/homework/pkg/repository"
)

type Item interface {
	Create(item models.Item) (*models.Item, error)
	GetAll() ([]models.Item, error)
	GetById(itemId int) (models.Item, error)
	Delete(itemId int) error
}

type Service struct {
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Item: NewItemService(repository.ItemRepository),
	}
}

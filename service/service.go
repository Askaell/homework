package service

import (
	GoArchitecture "github.com/Askaell/homework"
	"github.com/Askaell/homework/repository"
)

type Item interface {
	Create(item GoArchitecture.Item) (*GoArchitecture.Item, error)
	GetAll() ([]GoArchitecture.Item, error)
	GetById(itemId int) (GoArchitecture.Item, error)
	Delete(itemId int) error
}

type Service struct {
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Item: NewItemService(repository.Item),
	}
}

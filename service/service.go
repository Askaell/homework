package service

import (
	GoArchitecture "github.com/Askaell/homework"
	"github.com/Askaell/homework/repository"
)

type Item interface {
	Create(item GoArchitecture.Item) (*GoArchitecture.Item, error)
	GetAll() ([]GoArchitecture.Item, error)
}

type Service struct {
	Item
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Item: NewItemService(repository.Item),
	}
}

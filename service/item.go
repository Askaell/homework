package service

import (
	GoArchitecture "github.com/Askaell/homework"
	"github.com/Askaell/homework/repository"
)

type ItemService struct {
	repository repository.Item
}

func NewItemService(repository repository.Item) *ItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) Create(item GoArchitecture.Item) (newItem *GoArchitecture.Item, e error) {
	return s.repository.Create(item)
}

func (s *ItemService) GetAll() ([]GoArchitecture.Item, error) {
	return s.repository.GetAll()
}

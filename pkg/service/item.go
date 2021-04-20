package service

import (
	"github.com/Askaell/homework/pkg/models"
	"github.com/Askaell/homework/pkg/repository"
)

type ItemService struct {
	repository repository.ItemRepository
}

func NewItemService(repository repository.ItemRepository) *ItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) Create(item models.Item) (*models.Item, error) {
	return s.repository.Create(item)
}

func (s *ItemService) GetAll() ([]models.Item, error) {
	return s.repository.GetAll()
}

func (s *ItemService) GetById(itemId int) (models.Item, error) {
	return s.repository.GetById(itemId)
}

func (s *ItemService) Delete(itemId int) error {
	return s.repository.Delete(itemId)
}

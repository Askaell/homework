package fake_repository

import (
	"github.com/Askaell/homework/pkg/models"
)

func NewFakeRepository() *FakeRepo {
	return &FakeRepo{}
}

type FakeRepo struct {
	items []models.Item
}

func (f *FakeRepo) Create(item models.Item) (*models.Item, error) {
	f.items = append(f.items, item)
	return nil, nil
}

func (f *FakeRepo) Update(itemId int, item models.Item) error {
	for i := range f.items {
		if f.items[i].Id == itemId {
			f.items[i] = item
		}
	}

	return nil
}

func (f *FakeRepo) GetAll() ([]models.Item, error) {
	return f.items, nil
}

func (f *FakeRepo) GetById(itemId int) (models.Item, error) {
	for i := range f.items {
		if f.items[i].Id == itemId {
			return f.items[i], nil
		}
	}
	return models.Item{}, nil
}

func (f *FakeRepo) Delete(itemId int) error {
	return nil
}

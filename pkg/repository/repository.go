package repository

import (
	"github.com/Askaell/homework/pkg/models"
	"github.com/jmoiron/sqlx"
)

type ItemRepository interface {
	Create(item models.Item) (*models.Item, error)
	GetAll() ([]models.Item, error)
	GetById(itemId int) (models.Item, error)
	Delete(itemId int) error
}

type Repository struct {
	ItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ItemRepository: NewItemPostgres(db),
	}
}

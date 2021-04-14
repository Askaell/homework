package repository

import (
	GoArchitecture "github.com/Askaell/homework"
	"github.com/jmoiron/sqlx"
)

type Item interface {
	Create(item GoArchitecture.Item) (*GoArchitecture.Item, error)
	GetAll() ([]GoArchitecture.Item, error)
	GetById(itemId int) (GoArchitecture.Item, error)
}

type Repository struct {
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Item: NewItemPostgres(db),
	}
}

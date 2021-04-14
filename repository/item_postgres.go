package repository

import (
	"fmt"

	GoArchitecture "github.com/Askaell/homework"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(item GoArchitecture.Item) (newItem *GoArchitecture.Item, e error) {
	transaction, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (name, description, price) VALUES ($1, $2, $3) RETURNING id", itemTable)
	row := transaction.QueryRow(createItemQuery, item.Name, item.Description, item.Price)
	if err := row.Scan(&id); err != nil {
		transaction.Rollback()
		return nil, err
	}

	newItem = &GoArchitecture.Item{
		Id:          id,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}

	return newItem, transaction.Commit()
}

func (r *ItemPostgres) GetAll() ([]GoArchitecture.Item, error) {
	var items []GoArchitecture.Item

	query := fmt.Sprintf("SELECT * FROM %s", itemTable)
	err := r.db.Select(&items, query)

	return items, err
}

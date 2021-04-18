package repository

import (
	"fmt"

	"github.com/Askaell/homework/models"
	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(item models.Item) (*models.Item, error) {
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

	newItem := &models.Item{
		Id:          id,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}

	return newItem, transaction.Commit()
}

func (r *ItemPostgres) GetAll() ([]models.Item, error) {
	var items []models.Item

	query := fmt.Sprintf("SELECT * FROM %s", itemTable)
	err := r.db.Select(&items, query)

	return items, err
}

func (r *ItemPostgres) GetById(itemId int) (models.Item, error) {
	var item models.Item

	query := fmt.Sprintf("SELECT * FROM %s WHERE item.id = $1", itemTable)
	err := r.db.Get(&item, query, itemId)

	return item, err
}

func (r *ItemPostgres) Delete(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE item.id = $1", itemTable)
	_, err := r.db.Exec(query, itemId)

	return err
}

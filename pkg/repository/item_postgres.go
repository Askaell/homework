package repository

import (
	"fmt"
	"strings"

	"github.com/Askaell/homework/pkg/models"
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
	createItemQuery := fmt.Sprintf(
		"INSERT INTO %s (name, description, price, discountPrice, discount, dayItem, vendorCode, category) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		itemTable)

	row := transaction.QueryRow(createItemQuery, item.Name, item.Description,
		item.Price, item.DiscountPrice, item.Discount,
		item.DayItem, item.VendorCode, item.Category)

	if err := row.Scan(&id); err != nil {
		transaction.Rollback()
		return nil, err
	}

	newItem := &models.Item{
		Id:            id,
		Name:          item.Name,
		Description:   item.Description,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		Discount:      item.Discount,
		DayItem:       item.DayItem,
		VendorCode:    item.VendorCode,
		Category:      item.Category,
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

func (r *ItemPostgres) Update(itemId int, item models.Item) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if item.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, item.Name)
		argId++
	}

	if item.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, item.Description)
		argId++
	}

	if item.Price > 0 {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, item.Price)
		argId++
	}

	if item.DiscountPrice > 0 {
		setValues = append(setValues, fmt.Sprintf("discountPrice=$%d", argId))
		args = append(args, item.DiscountPrice)
		argId++
	}

	if item.Discount >= 0 {
		setValues = append(setValues, fmt.Sprintf("discount=$%d", argId))
		args = append(args, item.Discount)
		argId++
	}

	// item.DayItem
	setValues = append(setValues, fmt.Sprintf("dayItem=$%d", argId))
	args = append(args, item.DayItem)
	argId++

	if item.VendorCode != "" {
		setValues = append(setValues, fmt.Sprintf("vendorCode=$%d", argId))
		args = append(args, item.Discount)
		argId++
	}

	if item.Category != "" {
		setValues = append(setValues, fmt.Sprintf("category=$%d", argId))
		args = append(args, item.Category)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s it SET %s WHERE it.id = $%d`, itemTable, setQuery, argId)
	args = append(args, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}

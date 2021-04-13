package repository

import "github.com/jmoiron/sqlx"

type Item interface {
}

type Repository struct {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

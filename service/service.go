package service

import "github.com/Askaell/homework/repository"

type Item interface {
}

type Service struct {
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}

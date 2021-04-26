package service

import (
	"github.com/Askaell/homework/pkg/repository"
	"github.com/Askaell/homework/pkg/service/discount_service"
)

type DiscountService interface {
	Start(url string, activationTime string, timeLocation string)
}

func NewDiscountService(repository repository.ItemRepository) DiscountService {
	return discount_service.NewDiscountService(repository)
}

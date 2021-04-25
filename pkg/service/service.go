package service

import "github.com/Askaell/homework/pkg/service/discount_service"

type DiscountService interface {
	Start(time float32) error
	Stop()
}

func NewDiscountService() DiscountService {
	return discount_service.NewDiscountService()
}

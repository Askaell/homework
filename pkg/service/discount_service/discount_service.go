package discount_service

import (
	"log"
	"net/http"
	"time"

	"github.com/Askaell/homework/pkg/repository"
)

type discounts struct {
	category   map[string]float32
	vendorCode map[string]float32
	common     float32
}

type DiscountService struct {
	repository repository.ItemRepository
}

func NewDiscountService(repository repository.ItemRepository) *DiscountService {
	return &DiscountService{repository: repository}
}

func (s *DiscountService) Start(url string, activationTime string, timeLocation string) {
	go func() {
		for {
			timeNow, err := getTimeIn(timeLocation)
			if err != nil {
				log.Println("time error in discount_service: ", err)
				return
			}

			if timeNow.Format("15:04") != activationTime {
				time.Sleep(20 * time.Second)
				continue
			}

			discounts, err := s.getDiscounts(url)
			if err != nil {
				log.Println("getDiscounts fail: ", err)
				return
			}

			if err := s.applyDiscounts(discounts); err != nil {
				log.Println("applyDiscounts fail: ", err)
				return
			}

			time.Sleep(24 * time.Hour)
		}
	}()
}

func (s *DiscountService) getDiscounts(url string) (*discounts, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		log.Println(url, " fetch error: ", err)
		return nil, err
	}

	discounts, err := fetchDiscounts(resp.Body)
	resp.Body.Close()

	return discounts, err
}

func (s *DiscountService) applyDiscounts(d *discounts) error {
	items, err := s.repository.GetAll()
	if err != nil {
		return err
	}

	for _, item := range items {
		log.Println(item)
		item.DayItem = false
		item.Discount = d.common + d.category[item.Category]

		if vendorCodeDiscount := d.vendorCode[item.VendorCode]; vendorCodeDiscount != 0 {
			item.Discount += vendorCodeDiscount
			item.DayItem = true
		}

		item.DiscountPrice = item.Price - item.Price*(float32(item.Discount)/100.0)

		if err := s.repository.Update(item.Id, item); err != nil {
			return err
		}
		log.Println(item)
	}

	return nil
}

func getTimeIn(location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	t := time.Now().In(loc)
	return t, err
}

package tests

import (
	"testing"
	"time"

	"github.com/Askaell/homework/pkg/models"
	"github.com/Askaell/homework/pkg/service"
	"github.com/Askaell/homework/tests/fake_repository"
)

var repository = fake_repository.NewFakeRepository()
var discountService = service.NewDiscountService(repository)

var testItems = []models.Item{
	{
		Id:            0,
		Name:          "item 1",
		Description:   "",
		Price:         100,
		DiscountPrice: 0,
		Discount:      0,
		DayItem:       true,
		VendorCode:    "",
		Category:      "Sports",
	},
	{
		Id:            1,
		Name:          "item 2",
		Description:   "",
		Price:         1,
		DiscountPrice: 0,
		Discount:      0,
		DayItem:       false,
		VendorCode:    "123456",
		Category:      "Other",
	},
}

const (
	url            = "https://raw.githubusercontent.com/goarchitecture/lesson-2/feature/lesson-4/assets/discounts.csv"
	activationTime = "20:30"
	location       = "Europe/Moscow"
)

func init() {
	addTestItemsToRepo()
}

func addTestItemsToRepo() {
	for _, item := range testItems {
		repository.Create(item)
	}
}

func Test_discountService_cancelDayItem(t *testing.T) {
	testData := testItems[0]

	discountService.Start(
		url,
		time.Now().Format("15:04"),
		location)
	time.Sleep(1 * time.Second)

	result, _ := repository.GetById(testData.Id)

	if result.DayItem != false {
		t.Error("for ", testData, "\nexpected result: ", false, "\nreal result: ", result.DayItem)
	}
}

func Test_discountService_applyDayItem(t *testing.T) {
	testData := testItems[1]

	discountService.Start(
		url,
		time.Now().Format("15:04"),
		location)
	time.Sleep(1 * time.Second)

	result, _ := repository.GetById(testData.Id)

	if result.DayItem != true {
		t.Error("for ", testData, "\nexpected result: ", true, "\nreal result: ", result.DayItem)
	}
}

func Test_discountService_applyDiscounts(t *testing.T) {
	//getting from discount csv from urls
	expectedDiscounts := []int{
		15,
		17,
	}
	expectedPrices := []float32{}
	for i, d := range expectedDiscounts {
		expectedPrices = append(expectedPrices, testItems[i].Price-(testItems[i].Price*float32(d)/100))
	}

	discountService.Start(
		url,
		time.Now().Format("15:04"),
		location)
	time.Sleep(1 * time.Second)

	result, _ := repository.GetAll()

	for i, resultItem := range result {
		if resultItem.Discount != float32(expectedDiscounts[i]) && resultItem.DiscountPrice != expectedPrices[i] {
			t.Error("for ", resultItem, "\n expected result: ", expectedDiscounts[i], expectedPrices[i],
				"\nreal result: ", resultItem.Discount, resultItem.DiscountPrice)
		}
	}

}

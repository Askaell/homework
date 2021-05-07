package models

type Item struct {
	Id            int     `json:"id" db:"id"`
	Name          string  `json:"name" db:"name"`
	Description   string  `json:"description" db:"description"`
	Price         float32 `json:"price" db:"price"`
	DiscountPrice float32 `json:"discountPrice" db:"discountprice"`
	Discount      float32 `json:"discount" db:"discount"`
	DayItem       bool    `json:"dayItem" db:"dayitem"`
	VendorCode    string  `json:"vendorCode" db:"vendorcode"`
	Category      string  `json:"category" db:"category"`
	Amount        int     `json:"amount" db:"amount"`
}

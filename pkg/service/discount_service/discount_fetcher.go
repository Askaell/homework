package discount_service

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

//columns types numbers of discount table
const (
	colAmount = 3

	colDiscountType  = 0
	colDiscountName  = 1
	colDiscountValue = 2
)

//columns types of discount table
const (
	discountCategory   = "category"
	discountVendorCode = "item"
	discountCommon     = "-"
)

func fetchDiscounts(r io.Reader) (*discounts, error) {
	scanner := bufio.NewScanner(r)

	var commonDiscount float32 = 0
	categoryDiscount := make(map[string]float32)
	vendorCodeDiscount := make(map[string]float32)

	lineIdx := 1
	for scanner.Scan() {
		columns := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		if len(columns) != colAmount {
			continue
		}

		// scip table header
		if lineIdx == 1 {
			lineIdx++
			continue
		}

		discountValue, err := strconv.ParseFloat(columns[colDiscountValue], 32)
		if err != nil {
			log.Println("Wrong discount file format, at line ", lineIdx)
			return nil, err
		}
		// skip line with zero discount value
		if discountValue == 0 {
			continue
		}

		switch columns[colDiscountType] {
		case discountCategory:
			categoryDiscount[columns[colDiscountName]] = float32(discountValue)

		case discountVendorCode:
			vendorCodeDiscount[columns[colDiscountName]] = float32(discountValue)

		case discountCommon:
			commonDiscount = float32(discountValue)
		default:
			log.Println("Wrong discount type format, at line ", lineIdx)
			return nil, err
		}
		lineIdx++
	}

	return &discounts{categoryDiscount, vendorCodeDiscount, commonDiscount}, nil
}

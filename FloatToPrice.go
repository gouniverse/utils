package utils

import "strconv"

func FloatToPrice(priceFloat float64, currency string) string {
	price := strconv.FormatFloat(priceFloat, 'f', 2, 64)
	return ToCurrencySymbol(currency) + price
}

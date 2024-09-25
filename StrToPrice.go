package utils

import "strconv"

func StrToPrice(price string, currency string) string {
	priceFloat, errPrice := strconv.ParseFloat(price, 64)

	if errPrice != nil {
		return "n/a"
	}

	price = strconv.FormatFloat(priceFloat, 'f', 2, 64)
	return ToCurrencySymbol(currency) + price
}

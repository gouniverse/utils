package utils

import "strings"

func ToCurrencySymbol(currency string) string {
	if strings.ToUpper(currency) == "GBP" {
		return "&pound;"
	}
	if strings.ToUpper(currency) == "EUR" {
		return "&euro;"
	}
	if strings.ToUpper(currency) == "USD" {
		return "$"
	}
	return currency
}

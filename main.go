package main

import "fmt"

func main() {
	const USDinEUR = 0.84
	const USDinRUB = 84.92
	const EURinRUB = USDinRUB / USDinEUR
}

func getUserInput() (float64, string, string) {
	var value float64
	var originalCurrency string
	var targetCurrency string
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&originalCurrency)
	fmt.Print("Введите валюту для конвертации: ")
	fmt.Scan(&targetCurrency)
	fmt.Print("Введите колличество: ")
	fmt.Scan(&value)
	return value, originalCurrency, targetCurrency
}

func convert(value float64, originalCurrency string, targetCurrency string) {}

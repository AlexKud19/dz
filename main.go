package main

import (
	"errors"
	"fmt"
)

const USDinEUR = 0.84
const USDinRUB = 84.92
const EURinRUB = USDinRUB / USDinEUR
const USD = "USD"
const EUR = "EUR"
const RUB = "RUB"

func main() {
	currency := map[string]map[string]float64{
		USD: {USD: 1.0, EUR: USDinEUR, RUB: USDinRUB},
		EUR: {USD: 100 / USDinEUR, EUR: 1.0, RUB: EURinRUB},
		RUB: {USD: 100 / USDinRUB, EUR: 100 / EURinRUB, RUB: 1.0},
	}
	value, originalCurrency, targetCurrency := getUserInput()
	result := value * float32(currency[originalCurrency][targetCurrency])
	fmt.Printf("%.2f", result)
}

func getUserInput() (float32, string, string) {
	var value float32
	var originalCurrency string
	var targetCurrency string
	for {
		currency, err := getCurrencyInput(fmt.Sprintf("Введите исходную валюту (варианты: %v, %v, %v): ", USD, EUR, RUB))
		originalCurrency = currency
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	for {
		currencyCount, err := getCurrencyCountInput("Введите колличество: ")
		value = currencyCount
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	possibleOptions := map[string]string{
		USD: "EUR, RUB",
		EUR: "USD, RUB",
		RUB: "USD, EUR",
	}
	for {
		currency, err := getCurrencyInput(fmt.Sprintf("Введите валюту для конвертации: (варианты: %v): ", possibleOptions[originalCurrency]))
		targetCurrency = currency
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	return value, originalCurrency, targetCurrency
}

func getCurrencyInput(str string) (string, error) {
	var currency string
	fmt.Printf(str)
	fmt.Scan(&currency)
	if currency != USD && currency != EUR && currency != RUB {
		return "", errors.New("Введенное значение должно совпадать с предложенными вариантами")
	}
	return currency, nil
}

func getCurrencyCountInput(str string) (float32, error) {
	var currencyCount float32
	fmt.Printf(str)
	fmt.Scan(&currencyCount)
	if currencyCount <= 0 {
		return 0, errors.New("Введенное значение должно быть положительным числом")
	}
	return currencyCount, nil
}

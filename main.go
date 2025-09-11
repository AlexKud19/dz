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
	value, originalCurrency, targetCurrency := getUserInput()
	result := convert(value, originalCurrency, targetCurrency)
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
	possibleOptions := ""
	switch originalCurrency {
	case USD:
		possibleOptions = "EUR, RUB"
	case EUR:
		possibleOptions = "USD, RUB"
	default:
		possibleOptions = "USD, EUR"
	}
	for {
		currency, err := getCurrencyInput(fmt.Sprintf("Введите валюту для конвертации: (варианты: %v): ", possibleOptions))
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

func convert(value float32, originalCurrency string, targetCurrency string) float32 {
	var result float32
	switch originalCurrency {
	case USD:
		switch targetCurrency {
		case RUB:
			result = value * USDinRUB
		case EUR:
			result = value * USDinEUR
		default:
			result = value
		}
	case EUR:
		switch targetCurrency {
		case RUB:
			result = value * EURinRUB
		case USD:
			result = value / USDinEUR
		default:
			result = value
		}
	default:
		switch targetCurrency {
		case EUR:
			result = value / EURinRUB
		case USD:
			result = value / USDinRUB
		default:
			result = value
		}
	}
	return result
}

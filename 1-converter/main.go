package main

import (
	"fmt"
	"strings"
)

var rates = map[string]float64{
	"USD": 90.0,  // 1 USD = 90 RUB
	"EUR": 100.0, // 1 EUR = 100 RUB
	"RUB": 1.0,   // 1 RUB = 1 RUB (базовая валюта)
}

func main() {

	fromCurrency, amount, toCurrency, err := inputUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	convertAmount := convert(float64(amount), fromCurrency, toCurrency)

	fmt.Printf("%.2f %s → %.2f %s", amount, fromCurrency, convertAmount, toCurrency)
	// Вывод: 100.00 EUR → 9156.52 RUB
}

func inputUser() (string, float32, string, error) {

	var fromCurrency string
	var amount float32
	var toCurrency string

	keys := getKeyCurrency()
	keysImp := strings.Join(keys, "/")

	fmt.Printf("Введите изначальную валюту (%s): ", keysImp)
	fmt.Scan(&fromCurrency)
	if _, ok := rates[fromCurrency]; !ok {
		return "", 0, "", fmt.Errorf("неизвестная валюта: %s", fromCurrency)
	}

	keys = removeByValue(keys, fromCurrency)
	keysImp = strings.Join(keys, "/")

	fmt.Printf("Введите целевую валюту (%s): ", keysImp)
	fmt.Scan(&toCurrency)
	if _, ok := rates[toCurrency]; !ok {
		return "", 0, "", fmt.Errorf("неизвестная валюта: %s", toCurrency)
	}

	fmt.Printf("Введите сумму конвертации: ")
	fmt.Scan(&amount)

	return fromCurrency, amount, toCurrency, nil
}

func getKeyCurrency() []string {
	keys := make([]string, 0, len(rates))
	for k := range rates {
		keys = append(keys, k)
	}

	return keys
}

func removeByValue(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func convert(amount float64, fromCurrency, toCurrency string) float64 {
	result := (amount * rates[fromCurrency]) / rates[toCurrency]
	return result
}

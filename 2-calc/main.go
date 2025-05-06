package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	function, numberString, err := inputUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	numbers, err := getMassiveNumb(numberString)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := calcFunc(function, numbers)

	fmt.Println(result)
}

func inputUser() (string, string, error) {
	var inputFunc string
	var inputNumb string

	fmt.Print("Введите метод(AVG/SUM/MED): ")
	fmt.Scan(&inputFunc)

	switch inputFunc {
	case "AVG":
	case "SUM":
	case "MED":
		break
	default:
		return "", "", fmt.Errorf("неизвестная функция: %s", inputFunc)
	}

	fmt.Print("Введите числа: ")
	fmt.Scan(&inputNumb)

	return inputFunc, inputNumb, nil
}

func getMassiveNumb(numbString string) ([]int, error) {
	numbString = strings.ReplaceAll(numbString, " ", "")
	parts := strings.Split(numbString, ",")

	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		numb, err := strconv.Atoi(part)
		if err != nil {
			continue
		}

		numbers = append(numbers, numb)
	}

	if len(numbers) == 0 {
		return make([]int, 0, 0), fmt.Errorf("Введены не числа")
	}

	return numbers, nil

}

func calcFunc(nameFunc string, numbers []int) float64 {
	var result float64

	if nameFunc == "AVG" || nameFunc == "SUM" {
		for _, sum := range numbers {
			result += float64(sum)
		}

		if nameFunc == "SUM" {
			return result
		}

		return result / float64(len(numbers))
	}

	if nameFunc == "MED" {
		sort.Ints(numbers)
		l := len(numbers)
		if l%2 == 0 {
			result = float64((numbers[l/2-1] + numbers[l/2]) / 2)
		} else {
			result = float64(numbers[l/2])
		}

		return result
	}

	return 0
}

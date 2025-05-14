package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var calcOperation = map[string]func(...int) float64{
	"AVG": calcAvg,
	"SUM": calcSum,
	"MED": calcMed,
}

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

	calcFunc := calcOperation[function]

	result := calcFunc(numbers...)

	fmt.Println(result)
}

func inputUser() (string, string, error) {
	var inputFunc string
	var inputNumb string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите метод(AVG/SUM/MED): ")
	inputFunc, _ = reader.ReadString('\n')   // Считывает до Enter
	inputFunc = strings.TrimSpace(inputFunc) // Убираем пробелы и Enter

	_, ok := calcOperation[inputFunc]
	if ok == false {
		return "", "", fmt.Errorf("неизвестная функция: %s", inputFunc)
	}

	fmt.Print("Введите числа: ")
	inputNumb, _ = reader.ReadString('\n')   // Считывает до Enter
	inputNumb = strings.TrimSpace(inputNumb) // Убираем пробелы и Enter

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

func calcAvg(numbers ...int) float64 {
	var result float64

	result = calcSum(numbers...)
	return result / float64(len(numbers))
}

func calcSum(numbers ...int) float64 {
	var result float64
	for _, sum := range numbers {
		result += float64(sum)
	}
	return result
}

func calcMed(numbers ...int) float64 {
	var result float64
	sort.Ints(numbers)
	l := len(numbers)
	if l%2 == 0 {
		result = float64((numbers[l/2-1] + numbers[l/2]) / 2)
	} else {
		result = float64(numbers[l/2])
	}
	return result
}

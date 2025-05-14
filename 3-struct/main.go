package main

import (
	"fmt"
	"test/api/bins"
	"test/api/file"
	"test/api/storage"
)

func main() {
	// Создаем новый Bin
	bin := bins.CreateBin("12345", true, "My Bin")
	storage.SaveBin(bin)

	_, err := file.ReadFile("bins.json")

	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	} else {
		fmt.Println("Файл прочитан успешно")
	}

	isJson := file.IsJSON("bins.json")

	if !isJson {
		fmt.Println("Файл не является JSON-файлом")
	}
}

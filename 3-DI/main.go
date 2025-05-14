package main

import (
	"fmt"
	"test/api/api"
	"test/api/bins"
	"test/api/config"
	"test/api/file"
	"test/api/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := file.NewJsonDb("bins.json")
	api.NewApi(*config.NewConfig())
	binStorage := storage.NewBin(db)

	bin := bins.CreateBin("123", true, "My Bin")
	binStorage.AddBin(bin)
	binStorage.SaveBin()

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

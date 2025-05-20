package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"test/api/config"
	"test/api/file"
	"time"
)

const Url = "https://api.jsonbin.io/v3/b/"

type ApiStruct struct {
	ApiKey string
}

type IdBin struct {
	Id      string    `json:"id"`
	Created time.Time `json:"createdAt"`
	Private bool      `json:"private"`
}

type MetaData struct {
	MetaData IdBin `json:"metadata"`
}

func NewApi(conf config.Config) *ApiStruct {
	return &ApiStruct{
		ApiKey: conf.Key,
	}
}

func (a *ApiStruct) Create(nameBin string, fileName string) *IdBin {
	if nameBin == "" {
		return nil
	}

	isJson := file.IsJSON(fileName)
	if !isJson {
		fmt.Println("Выбран не JSON файл")
		return nil
	}

	jsonData, err := file.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка", err)
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, Url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return nil
	}

	// Заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.ApiKey)
	req.Header.Set("X-Bin-Name", nameBin)

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: получен статус %s\n", resp.Status)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка парсинга:", err)
		return nil
	}

	var idB MetaData
	json.Unmarshal(body, &idB)

	fmt.Println(idB)

	return &idB.MetaData

}

func (a *ApiStruct) Delete(id string) bool {
	req, err := http.NewRequest(http.MethodDelete, Url+id, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return false
	}

	// Заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.ApiKey)

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: получен статус %s\n", resp.Status)
		return false
	}

	return true
}

func (a *ApiStruct) Update(id string, fileName string) error {
	isJson := file.IsJSON(fileName)
	if !isJson {
		fmt.Println("Выбран не JSON файл")
		return errors.New("ERROR")
	}

	jsonData, err := file.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPut, Url+id, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return err
	}

	// Заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.ApiKey)

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: получен статус %s\n", resp.Status)
		return errors.New("ERROR")
	}

	return nil
}

func (a *ApiStruct) Get(id string) error {
	req, err := http.NewRequest(http.MethodGet, Url+id, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return err
	}

	// Заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", a.ApiKey)

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: получен статус %s\n", resp.Status)
		return errors.New("ERROR")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка парсинга:", err)
		return err
	}

	fmt.Println(string(body))
	var idB MetaData
	json.Unmarshal(body, &idB)

	fmt.Println(idB)
	return nil
}

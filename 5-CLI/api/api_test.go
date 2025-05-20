package api_test

import (
	"test/api/api"
	"test/api/config"
	"testing"
)

func TestCreate(t *testing.T) {
	newReq := api.NewApi(*config.NewConfig())
	result := newReq.Create("test", "test.json")
	if result.Id == "" {
		t.Error("Ошибка запроса Create")
	} else {
		newReq.Delete(result.Id)
	}
}

func TestUpdate(t *testing.T) {
	newReq := api.NewApi(*config.NewConfig())
	err := newReq.Update("123", "test.json")
	if err != nil {
		t.Error("Ошибка запроса Update")
	}
}
func TestDelete(t *testing.T) {
	newReq := api.NewApi(*config.NewConfig())
	isDeleted := newReq.Delete("123")
	if !isDeleted {
		t.Error("Ошибка запроса Delete")
	}
}

func TestGet(t *testing.T) {
	newReq := api.NewApi(*config.NewConfig())
	err := newReq.Get("123")
	if err != nil {
		t.Error("Ошибка запроса GET")
	}
}

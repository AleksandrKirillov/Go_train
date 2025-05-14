package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{filename: filename}
}

func (db *JsonDb) Read() ([]byte, error) {
	return os.ReadFile(db.filename)
}

func (db *JsonDb) Write(content []byte) {
	err := os.WriteFile(db.filename, content, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// ReadFile читает содержимое любого файла
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// IsJSON проверяет, что файл имеет расширение .json
func IsJSON(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".json"
}

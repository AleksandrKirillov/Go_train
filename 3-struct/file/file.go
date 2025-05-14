package file

import (
	"os"
	"path/filepath"
	"strings"
)

// ReadFile читает содержимое любого файла
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// IsJSON проверяет, что файл имеет расширение .json
func IsJSON(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".json"
}

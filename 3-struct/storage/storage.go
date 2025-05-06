package storage

import (
	"encoding/json"
	"os"
	"test/api/bins"
)

const storageFile = "bins.json"

func SaveBin(bin bins.Bin) error {
	bins, err := LoadBins()
	if err != nil {
		return err
	}

	bins = append(bins, bin)

	data, err := json.MarshalIndent(bins, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0644)
}

// LoadBins загружает список Bin из JSON-файла
func LoadBins() ([]bins.Bin, error) {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []bins.Bin{}, nil // Файла нет — возвращаем пустой список
		}
		return nil, err
	}

	var bins []bins.Bin
	if err := json.Unmarshal(data, &bins); err != nil {
		return nil, err
	}

	return bins, nil
}

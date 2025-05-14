package storage

import (
	"encoding/json"
	"test/api/bins"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write(content []byte)
}

type BinList struct {
	Bins []bins.Bin `json:"bins"`
}

type BinListWithDb struct {
	BinList
	Db Db
}

func NewBin(db Db) *BinListWithDb {
	data, err := db.Read()
	if err != nil {
		return &BinListWithDb{
			BinList: BinList{},
			Db:      db,
		}
	}

	return &BinListWithDb{
		BinList: *LoadBins(data),
		Db:      db,
	}
}

func (b *BinList) AddBin(bin bins.Bin) {
	b.Bins = append(b.Bins, bin)
}

func (b *BinListWithDb) SaveBin() error {

	data, err := json.MarshalIndent(b.BinList, "", "  ")
	if err != nil {
		return err
	}

	b.Db.Write(data)

	return nil
}

// LoadBins загружает список Bin из JSON-файла
func LoadBins(data []byte) *BinList {
	var bins BinList
	if err := json.Unmarshal(data, &bins); err != nil {
		color.Red("Error: %v", err.Error())
	}

	return &bins
}

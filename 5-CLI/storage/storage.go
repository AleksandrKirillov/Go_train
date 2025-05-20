package storage

import (
	"encoding/json"
	"fmt"
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

func (b *BinList) DeleteBinById(id string) {
	for i, v := range b.Bins {
		if v.ID == id {
			b.Bins = append(b.Bins[:i], b.Bins[i+1:]...)
			break
		}
	}
}

func (b *BinListWithDb) SaveBin() error {

	data, err := json.MarshalIndent(b.BinList, "", "  ")
	if err != nil {
		return err
	}

	b.Db.Write(data)

	return nil
}

func (b *BinList) ShowBin() {
	for _, value := range b.Bins {
		fmt.Printf("\nID: %s, Name: %s", value.ID, value.Name)
	}
}

// LoadBins загружает список Bin из JSON-файла
func LoadBins(data []byte) *BinList {
	var bins BinList
	if err := json.Unmarshal(data, &bins); err != nil {
		color.Red("Error: %v", err.Error())
	}

	return &bins
}

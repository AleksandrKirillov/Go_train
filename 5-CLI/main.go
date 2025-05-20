package main

import (
	"test/api/api"
	"test/api/bins"
	"test/api/cmd"
	"test/api/config"
	"test/api/file"
	"test/api/storage"
)

func main() {
	param := cmd.Execute()
	db := file.NewJsonDb("bins.json")
	// Initialize the API with the configuration
	api := api.NewApi(*config.NewConfig())
	// Initialize the bin storage with the database
	// and the API key
	binStorage := storage.NewBin(db)

	switch param.Mode {
	case "create":
		newReqBin := api.Create(*param.Name, *param.Filename)
		if newReqBin.Id == "" {
			return
		}
		bin := bins.CreateBin(newReqBin.Id, newReqBin.Private, *param.Name)
		binStorage.AddBin(bin)
		binStorage.SaveBin()
	case "update":
		api.Update(*param.Id, *param.Filename)
	case "get":
		api.Get(*param.Id)
	case "delete":
		isDeleted := api.Delete(*param.Id)
		if !isDeleted {
			return
		}

		binStorage.DeleteBinById(*param.Id)
		binStorage.SaveBin()

	case "list":
		binStorage.ShowBin()
	default:
		println("Unknown mode:", param.Mode)
	}

}

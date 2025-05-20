package cmd

import (
	"flag"

	"github.com/joho/godotenv"
)

type CmdParam struct {
	Mode     string
	Filename *string
	Name     *string
	Id       *string
}

func Execute() *CmdParam {
	// Execute the root command
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	param := &CmdParam{}

	isCreate := flag.Bool("create", false, "Creatm.Mode")
	isUpdate := flag.Bool("update", false, "Updatm.Mode")
	isDelete := flag.Bool("delete", false, "Deletm.Mode")
	isGet := flag.Bool("get", false, "Gem.Mode")
	isList := flag.Bool("list", false, "Lism.Mode")

	//Parameters
	param.Filename = flag.String("file", "", "File Json")
	param.Name = flag.String("name", "", "NameBin")
	param.Id = flag.String("id", "", "IdBin")

	flag.Parse()

	switch true {
	case *isCreate:
		param.Mode = "create"
	case *isUpdate:
		param.Mode = "update"
	case *isDelete:
		param.Mode = "delete"
	case *isGet:
		param.Mode = "get"
	case *isList:
		param.Mode = "list"
	}

	return param

}

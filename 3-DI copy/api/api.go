package api

import "test/api/config"

type ApiStruct struct {
	ApiKey string
}

func NewApi(conf config.Config) *ApiStruct {
	return &ApiStruct{
		ApiKey: conf.Key,
	}

}

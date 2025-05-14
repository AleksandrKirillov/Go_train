package api

import "test/api/config"

type Api struct {
	// Add fields as needed
	ApiKey string
}

func NewApi(conf config.Config) *Api {
	return &Api{
		ApiKey: conf.Key,
	}
}

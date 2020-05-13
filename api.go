package main

import "github.com/gin-gonic/gin"

type API struct {
	e      *gin.Engine
	engine *Engine
	Port   string
	prefix string
	done   chan error
}

func newDBAPI(prefix, port string) (*API, error) {
	eng, err := NewEngine()
	if err != nil {
		return nil, err
	}
	return &API{
		e:      gin.Default(),
		Port:   port,
		prefix: prefix,
		engine: eng,
	}, nil
}

func (api *API) registerEndpoints() {
	r := api.e.Group(api.prefix)

	api.createTable(r)
	api.GetAll(r)
	api.createNewRecord(r)
}

func (api *API) Launch() error {
	api.registerEndpoints()
	return api.e.Run(api.Port)
}

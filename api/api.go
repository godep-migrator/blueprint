package api

import (
	"github.com/twitchscience/blueprint/core"
	"github.com/twitchscience/blueprint/scoopclient"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

type server struct {
	docRoot    string
	datasource scoopclient.ScoopClient
}

func New(docRoot string, client scoopclient.ScoopClient) core.Subprocess {
	return &server{
		docRoot:    docRoot,
		datasource: client,
	}
}

func (s *server) Setup() error {
	files := web.New()
	files.Get("/*", s.fileHandler)
	files.NotFound(fourOhFour)

	api := web.New()
	api.Use(jsonResponse)
	api.Put("/schema", s.createSchema)
	api.Get("/schemas", s.allSchemas)
	api.Get("/schema/:id", s.schema)
	api.Post("/schema/:id", s.updateSchema)
	api.Get("/types", s.types)
	api.Post("/expire", s.expire)

	// Order is important here
	goji.Handle("/schema*", api)
	goji.Handle("/types", api)
	goji.Handle("/expire", api)
	goji.Handle("/*", files)

	// Stop() provides our shutdown semantics
	graceful.ResetSignals()

	return nil
}

func (s *server) Start() {

	goji.Serve()
}

func (s *server) Stop() {
	graceful.Shutdown()
}

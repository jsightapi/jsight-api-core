package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type OpenAPI struct {
	catalog *catalog.Catalog

	OpenAPI string    `json:"openapi"`
	Info    *Info     `json:"info"`
	Servers *[]Server `json:"servers"`
	Paths Paths `json:"paths"`
}

func NewOpenAPI(c *catalog.Catalog) *OpenAPI {
	return &OpenAPI{
		catalog: c,
		OpenAPI: "3.0.3",
		Info:    NewInfo(c.Info),
		Servers: NewServers(c.Servers),
    Paths: NewPaths(c),
	}
}

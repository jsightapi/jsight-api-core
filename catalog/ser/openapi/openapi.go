package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type OpenAPI struct {
	catalog *catalog.Catalog

	OpenAPI string `json:"openapi"`
	Info    string `json:"info"`
}

func NewOpenAPI(c *catalog.Catalog) *OpenAPI {
	return &OpenAPI{
		catalog: c,
		OpenAPI: "3.1.0",
	}
}

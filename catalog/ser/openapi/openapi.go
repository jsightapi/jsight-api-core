package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog/iface"
)

type OpenAPI struct {
	catalog iface.Catalog

	OpenAPI string `json:"openapi"`
	Info    *Info  `json:"info"`
}

func NewOpenAPI(c iface.Catalog) *OpenAPI {
	return &OpenAPI{
		catalog: c,
		OpenAPI: "3.1.0",
		Info:    NewInfo(c.GetInfo()),
	}
}

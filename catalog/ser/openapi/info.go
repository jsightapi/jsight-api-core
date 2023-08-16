package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog/iface"
)

type Info struct {
	Title string `json:"title"`
}

func NewInfo(i iface.Info) *Info {
	return &Info{
		Title: i.GetTitle(),
	}
}

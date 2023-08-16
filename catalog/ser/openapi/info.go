package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Info struct {
	Title string `json:"title"`
}

func NewInfo(i *catalog.Info) *Info {
	return &Info{
		Title: i.GetTitle(),
	}
}

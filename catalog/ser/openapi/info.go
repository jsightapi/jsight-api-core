package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Info struct {
	Title string `json:"title"`
}

func NewInfo(i *catalog.Info) *Info {
	if i == nil {
		return &Info{
			Title: "",
		}
	}

	return &Info{
		Title: i.Title,
	}
}

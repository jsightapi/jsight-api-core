package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	// Parameters  *PathItemParameters `json:"parameters,omitempty"`// TODO: deal with params
}

// type PathItemParameters struct{}

func NewPathItem(i *catalog.HTTPInteraction) *PathItem {
	pi := PathItem{}

	pi.assignOperation(i.HttpMethod, NewOperation(i))

	return &pi
}

// TODO: deal with possible ovewriting of method (improbable)
func (pi *PathItem) assignOperation(method catalog.HTTPMethod, o *Operation) {
	switch method {
	case catalog.GET:
		pi.Get = o
	case catalog.PUT:
		pi.Put = o
	case catalog.POST:
		pi.Post = o
	case catalog.PATCH:
		pi.Patch = o
	case catalog.DELETE:
		pi.Delete = o
	default:
		panic("Unsupported method")
	}
}

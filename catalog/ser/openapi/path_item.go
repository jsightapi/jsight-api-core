package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type PathItem struct {
	Parameters []*ParameterObject `json:"parameters,omitempty"`
	Get        *Operation         `json:"get,omitempty"`
	Put        *Operation         `json:"put,omitempty"`
	Post       *Operation         `json:"post,omitempty"`
	Patch      *Operation         `json:"patch,omitempty"`
	Delete     *Operation         `json:"delete,omitempty"`
	// Parameters  *PathItemParameters `json:"parameters,omitempty"`// TODO: deal with params
}

func newPathItem(i *catalog.HTTPInteraction) *PathItem {
	pi := PathItem{
		Parameters: fillPathParams(i),
	}
	pi.assignOperation(i.HttpMethod, newOperation(i))
	return &pi
}

func fillPathParams(i *catalog.HTTPInteraction) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	appendPathParams(r, i)
	return r
}

func pathSchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.PathVariables != nil &&
		i.PathVariables.Schema != nil
}

func appendPathParams(p []*ParameterObject, i *catalog.HTTPInteraction) {
	if pathSchemaDefined(i) {
		_ = append(p, getPathParams(i)...)
	}
}

// TODO: test no path directive
func getPathParams(i *catalog.HTTPInteraction) []*ParameterObject {
	return paramsFromJSchema(i.PathVariables.Schema, ParameterLocationPath)
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

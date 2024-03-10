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

func NewPathItem(i *catalog.HTTPInteraction) *PathItem {
	pi := PathItem{
		Parameters: FillPathParams(i),
	}
	pi.assignOperation(i.HttpMethod, NewOperation(i))
	return &pi
}

func FillPathParams(i *catalog.HTTPInteraction) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	AppendPathParams(r, i)
	return r
}

func PathSchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.PathVariables != nil &&
		i.PathVariables.Schema != nil
}

func AppendPathParams(p []*ParameterObject, i *catalog.HTTPInteraction) {
	if PathSchemaDefined(i) {
		_ = append(p, GetPathParams(i)...)
	}
}

// TODO: test no path directive
func GetPathParams(i *catalog.HTTPInteraction) []*ParameterObject {
	return ParamsFromSchema(i.PathVariables.Schema, ParameterLocationPath)
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

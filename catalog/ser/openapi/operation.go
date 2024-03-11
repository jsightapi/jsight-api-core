package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Operation struct {
	Summary     *string            `json:"summary,omitempty"`
	Description *string            `json:"description,omitempty"`
	Parameters  []*ParameterObject `json:"parameters,omitempty"`
	RequestBody *RequestBody       `json:"requestBody,omitempty"`
	Responses   *Responses         `json:"responses"`
}

func newOperation(i *catalog.HTTPInteraction) *Operation {
	return &Operation{
		Summary:     i.Annotation,
		Description: i.Description,
		Parameters:  fillOperationParams(i),
		RequestBody: newRequestBody(i.Request),
		Responses:   newResponses(i),
	}
}

func fillOperationParams(i *catalog.HTTPInteraction) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	r = appendQueryParams(r, i)
	r = appendHeaderParams(r, i)
	return r
}

func appendQueryParams(p []*ParameterObject, i *catalog.HTTPInteraction) []*ParameterObject {
	if querySchemaDefined(i) {
		return append(p, paramsFromSchema(i.Query.Schema, ParameterLocationQuery)...)
	}
	return p
}

func appendHeaderParams(p []*ParameterObject, i *catalog.HTTPInteraction) []*ParameterObject {
	if headerSchemaDefined(i) {
		return append(p, paramsFromSchema(i.Request.HTTPRequestHeaders.Schema, ParameterLocationHeader)...)
	}
	return p
}

func querySchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.Query != nil && i.Query.Schema != nil
}

func headerSchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.Request != nil &&
		i.Request.HTTPRequestHeaders != nil &&
		i.Request.HTTPRequestHeaders.Schema != nil
}

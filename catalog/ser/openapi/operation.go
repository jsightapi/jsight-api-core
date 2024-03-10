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
		RequestBody: newRequestBody(i.Request), // TODO: is request enough?
		Responses:   newResponses(i),
	}
}

func fillOperationParams(i *catalog.HTTPInteraction) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	appendQueryParams(r, i)
	appendHeaderParams(r, i)
	return r
}

func appendQueryParams(p []*ParameterObject, i *catalog.HTTPInteraction) {
	if querySchemaDefined(i) {
		p = append(p, getQueryParams(i)...)
	}
}

func appendHeaderParams(p []*ParameterObject, i *catalog.HTTPInteraction) {
	if headerSchemaDefined(i) {
		p = append(p, getHeaderParams(i)...)
	}
}

func querySchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.Query != nil &&
		i.Query.Schema != nil
}

func headerSchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.Request != nil &&
		i.Request.HTTPRequestHeaders != nil &&
		i.Request.HTTPRequestHeaders.Schema != nil
}

func getQueryParams(i *catalog.HTTPInteraction) []*ParameterObject {
	// TODO: nil check
	return paramsFromSchema(i.Query.Schema, ParameterLocationQuery)
}

func getHeaderParams(i *catalog.HTTPInteraction) []*ParameterObject {
	return paramsFromSchema(i.Request.HTTPRequestHeaders.Schema, ParameterLocationHeader)
}

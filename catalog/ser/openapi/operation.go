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

func NewOperation(i *catalog.HTTPInteraction) *Operation {
	return &Operation{
		Summary:     i.Annotation,
		Description: i.Description,
		Parameters:  FillOperationParams(i),
		RequestBody: NewRequestBody(i.Request), // TODO: is request enough?
		Responses:   NewResponses(i),
	}
}

func FillOperationParams(i *catalog.HTTPInteraction) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	AppendQueryParams(r, i)
	AppendHeaderParams(r, i)
	return r
}

func AppendQueryParams(p []*ParameterObject, i *catalog.HTTPInteraction) {
	if QuerySchemaDefined(i) {
		p = append(p, GetQueryParams(i)...)
	}
}

func AppendHeaderParams(p []*ParameterObject, i *catalog.HTTPInteraction) {
	if HeaderSchemaDefined(i) {
		p = append(p, GetHeaderParams(i)...)
	}
}

func QuerySchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.Query != nil &&
		i.Query.Schema != nil
}

func HeaderSchemaDefined(i *catalog.HTTPInteraction) bool {
	return i.Request != nil &&
		i.Request.HTTPRequestHeaders != nil &&
		i.Request.HTTPRequestHeaders.Schema != nil
}

func GetQueryParams(i *catalog.HTTPInteraction) []*ParameterObject {
	// TODO: nil check
	return ParamsFromSchema(i.Query.Schema, ParameterLocationQuery)
}

func GetHeaderParams(i *catalog.HTTPInteraction) []*ParameterObject {
	return ParamsFromSchema(i.Request.HTTPRequestHeaders.Schema, ParameterLocationHeader)
}

package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Operation struct {
	Summary     string             `json:"summary,omitempty"`
	Description string             `json:"description,omitempty"`
	Parameters  []*ParameterObject `json:"parameters,omitempty"`
	RequestBody *RequestBody       `json:"requestBody,omitempty"`
	Responses   *Responses         `json:"responses"`
	Tags        []string           `json:"tags,omitempty"`
}

func newOperation(i *catalog.HTTPInteraction) *Operation {
	return &Operation{
		Summary:     processMethodSummary(i.Annotation),
		Description: processMethodDescription(i.Description),
		Parameters:  fillOperationParams(i),
		RequestBody: newRequestBody(i.Request),
		Responses:   newResponses(i),
		Tags:        getTags(i),
	}
}

func getTags(i *catalog.HTTPInteraction) []string {
	r := make([]string, 0, len(i.Tags))
	for _, t := range i.Tags {
		r = append(r, string(t))
	}
	return r
}

func processMethodDescription(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func processMethodSummary(s *string) string {
	if s == nil {
		return ""
	}

	r := *s
	return r
}

func fillOperationParams(i *catalog.HTTPInteraction) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	r = appendQueryParams(r, i)
	r = appendHeaderParams(r, i)
	return r
}

func appendQueryParams(p []*ParameterObject, i *catalog.HTTPInteraction) []*ParameterObject {
	if querySchemaDefined(i) {
		params := paramsFromJSchema(i.Query.Schema, ParameterLocationQuery)
		for _, par := range params {
			par.Style = ParameterStyleDeepObject
			par.Explode = true
		}
		return append(p, params...)
	}
	return p
}

func appendHeaderParams(p []*ParameterObject, i *catalog.HTTPInteraction) []*ParameterObject {
	if headerSchemaDefined(i) {
		return append(p, paramsFromJSchema(i.Request.HTTPRequestHeaders.Schema, ParameterLocationHeader)...)
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

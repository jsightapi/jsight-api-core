package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Operation struct {
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
  // Parameters 
  RequestBody *RequestBody `json:"requestBody,omitempty"`
  // Responses
}

func NewOperation(i *catalog.HTTPInteraction) *Operation {
	return &Operation{
    Summary: *i.Annotation,
    Description: *i.Description,
    RequestBody: NewRequestBody(i.Request),
  }
}

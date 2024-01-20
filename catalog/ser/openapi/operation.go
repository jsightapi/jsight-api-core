package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Operation struct {
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
  // Parameters 
  // RequestBody
  // Responses
}

func NewOperation(o *catalog.HTTPInteraction) *Operation {
	return &Operation{}
}

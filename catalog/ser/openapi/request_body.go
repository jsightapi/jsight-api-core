package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

// import "github.com/jsightapi/jsight-api-core/catalog"

type RequestBody struct {
	// Description string `json:"description,omitempty"` // not supported in JSight
  Content Content `json:"content"`
  Required bool `json:"required,omitempty"`
}

func NewRequestBody(r *catalog.HTTPRequest) *RequestBody {
  return &RequestBody{
    Required: r.,
  }
}

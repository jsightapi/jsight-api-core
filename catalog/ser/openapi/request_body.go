package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"
)

// import "github.com/jsightapi/jsight-api-core/catalog"

type RequestBody struct {
	Content  *Content `json:"content"`
	Required bool     `json:"required,omitempty"`
}

func NewRequestBody(r *catalog.HTTPRequest) *RequestBody {
	if r == nil {
		return nil
	}

	var c *Content
	s := r.HTTPRequestBody.Schema
	switch s.Notation() {
	case notation.SchemaNotationJSight, notation.SchemaNotationRegex:
		c = ContentForSchema(r.Format, s)
	case notation.SchemaNotationAny:
		c = ContentForAny()
	case notation.SchemaNotationEmpty:
		c = ContentForEmpty()
	}

	return &RequestBody{
		Required: isRequestBodyRequired(r),
		Content:  c,
	}
}

func isRequestBodyRequired(r *catalog.HTTPRequest) bool {
	return false // TODO: discuss
}

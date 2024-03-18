package openapi

import (
	"fmt"

	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"
	schema "github.com/jsightapi/jsight-schema-core"
)

type RequestBody struct {
	Content  *Content `json:"content"`
	Required bool     `json:"required,omitempty"`
}

func newRequestBody(r *catalog.HTTPRequest) *RequestBody {
	if r == nil {
		return nil
	}

	var c *Content
	s := r.HTTPRequestBody.Schema
	switch s.Notation() {
	case notation.SchemaNotationJSight, notation.SchemaNotationRegex:
		fmt.Printf("FORMAT IS: %s", r.Format)
		c = contentForSchema(r.Format, s)
	case notation.SchemaNotationAny:
		c = contentForAny()
	case notation.SchemaNotationEmpty:
		c = contentForEmpty()
	}

	return &RequestBody{
		Required: requestBodyRequired(r),
		Content:  c,
	}
}

func requestBodyRequired(r *catalog.HTTPRequest) bool {
	if r.HTTPRequestBody == nil ||
		r.HTTPRequestBody.Schema == nil ||
		r.HTTPRequestBody.Schema.Notation() == notation.SchemaNotationAny ||
		schemaReferencesTypeAny(r.HTTPRequestBody.Schema) {
		return false
	}

	return true
}

// TODO: cannot be finished until parser bug if fixed and may be smth extra in schemaInfo
func schemaReferencesTypeAny(es catalog.ExchangeSchema) bool {
	if es.Notation() == notation.SchemaNotationJSight {
		// js := es.(*catalog.ExchangeJSightSchema)
		ast, _ := es.GetAST()
		if ast.TokenType == schema.TokenTypeShortcut {
			// types, _ := es.UsedUserTypes()
			// fmt.Printf("\nTypes %v", types)
			// typeName := types[0]
		}
	}
	return false
}

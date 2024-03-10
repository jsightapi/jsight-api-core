package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	schema "github.com/jsightapi/jsight-schema-core"
)

/*
Other properties of OA MediaTypeObject are not used in JSight
*/
type MediaTypeObject struct {
	Schema SchemaObject `json:"schema,omitempty"` // TODO: empty?
}

func NewMediaTypeObjectFromExchangeSchema(es catalog.ExchangeSchema) *MediaTypeObject {
	so := SchemaObjectFromExchangeSchema(es)

	return &MediaTypeObject{
		Schema: so,
	}
}

func NewMediaTypeObjectFromSchema(s schema.Schema) *MediaTypeObject {
	so := SchemaObjectFromSchema(s)

	return &MediaTypeObject{
		Schema: so,
	}
}

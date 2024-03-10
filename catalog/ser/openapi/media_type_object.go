package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	// schema "github.com/jsightapi/jsight-schema-core"
)

// Other properties of OA MediaTypeObject are not used in JSigh
type MediaTypeObject struct {
	Schema SchemaObject `json:"schema,omitempty"` // TODO: empty?
}

func mediaTypeObjectForSchema(es catalog.ExchangeSchema) *MediaTypeObject {
	return &MediaTypeObject{
		Schema: schemaObjectFromExchangeSchema(es),
	}
}

// func NewMediaTypeObjectFromSchema(s schema.Schema) *MediaTypeObject {
// 	so := SchemaObjectFromSchema(s)
//
// 	return &MediaTypeObject{
// 		Schema: so,
// 	}
// }

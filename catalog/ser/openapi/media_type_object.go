package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	schema "github.com/jsightapi/jsight-schema-core"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

/*
Other properties of OA MediaTypeObject are not used in JSight
*/
type MediaTypeObject struct {
	Schema sc.SchemaObject `json:"schema,omitempty"` // TODO: empty?
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

package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

type MediaTypeObject struct {
	Schema sc.SchemaObject `json:"schema,omitempty"` // TODO: empty?
}

func NewMediaTypeObject(es catalog.ExchangeSchema) *MediaTypeObject {
	so := SchemaObjectFromExchangeSchema(es)

	return &MediaTypeObject{
		Schema: so,
	}
}

func NewMediaTypeObjectFromSchemaKeeper[T sc.SchemaKeeper](sk T) *MediaTypeObject {
	so := SchemaObjectFromSchemaKeeper(sk)

	return &MediaTypeObject{
		Schema: so,
	}
}

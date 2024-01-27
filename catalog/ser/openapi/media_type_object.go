package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type MediaTypeObject struct {
	Schema   SchemaObject     `json:"schema"` // TODO: empty?
	Examples []*ExampleObject `json:"examples,omitempty"`
	// Example // mutully excl. with Examples and deprecated in OAS 3.1, hence not using in favor of Examples
	// Encoding // irrelevant to currently supported media types
}

func NewMediaTypeObject(es catalog.ExchangeSchema) *MediaTypeObject {

	so := ExhangeSchemaToSchemaObject(es)

	eb, err := es.Example()
	if err != nil {
		panic(err) // TODO: discuss
	}

	ex := make([]*ExampleObject, 0, 1)
	ex = append(ex, NewExampleObject(eb))

	return &MediaTypeObject{
		Schema:   so,
		Examples: ex,
	}
}

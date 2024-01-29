package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type MediaTypeObject struct {
	Schema SchemaObject `json:"schema,omitempty"` // TODO: empty?
	// Examples map[string]*ExampleObject `json:"examples,omitempty"` // Not using until decided otherwise
	// Example // mutully excl. with Examples and deprecated in OAS 3.1, hence not using in favor of Examples
	// Encoding // irrelevant to currently supported media types
}

// const defaultExampleName string = "example"

func NewMediaTypeObject(es catalog.ExchangeSchema) *MediaTypeObject {
	so := ExhangeSchemaToSchemaObject(es)

	return &MediaTypeObject{
		Schema: so,
		// Examples: makeExamples(es),
	}
}

// func makeExample(es catalog.ExchangeSchema) map[string]*ExampleObject {
// 	eb, err := es.Example()
// 	if err != nil {
// 		panic(err) // TODO: discuss
// 	}
//
// 	ex := make(map[string]*ExampleObject, 1)
// 	ex[defaultExampleName] = NewExampleObject(eb)
//
// 	return ex
// }

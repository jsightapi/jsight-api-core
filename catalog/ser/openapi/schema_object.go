package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"

	openapiSchema "github.com/jsightapi/jsight-schema-core/openapi"
)

type SchemaObject interface{}

/*
TODO:

	SchemaObject is the only structure allowed as a schema inside comonents
	so we must convert all of our notations into it.
	If some notation affects higher level structure of OA,
	this approach should be changed.
*/
func SchemaObjectFromUserType(t *catalog.UserType) SchemaObject {
	return ExhangeSchemaToSchemaObject(t.Schema)
}

func dummySchemaObject() SchemaObject {
	return &DummySchemaObject{}
}

type DummySchemaObject struct{}

func ExhangeSchemaToSchemaObject(e catalog.ExchangeSchema) SchemaObject {
	// t.Schema is an ExchangeSchema iface.
	// must cast to ExchangeJSightSchema
	switch v := e.(type) {
	case *catalog.ExchangeJSightSchema:
		// pass to schema-core
		// TODO: discuss interface
		return (openapiSchema.NewSchemaObject(v.JSchema))
	case catalog.ExchangeRegexSchema:
		// make SchemaObject appr. for regex
		return dummySchemaObject() // TODO:
	case catalog.ExchangePseudoSchema:
		// any - anything allowed. Srialize to empty object
		// empty - must be empty
		return dummySchemaObject() // TODO:
	default:
		panic("UserType schema has unsupported type")
	}
}

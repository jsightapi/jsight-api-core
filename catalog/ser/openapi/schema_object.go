package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	schema "github.com/jsightapi/jsight-schema-core"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

/*
TODO:

	SchemaObject is the only structure allowed as a schema inside comonents
	so we must convert all of our notations into it.
	If some notation affects higher level structure of OA,
	this approach should be changed.
*/
func SchemaObjectFromUserType(t *catalog.UserType) sc.SchemaObject {
	// TODO: can we have UserType for any / empty ?

	return SchemaObjectFromExchangeSchema(t.Schema)
}

// TODO: might become unnecessary
func SchemaObjectFromExchangeSchema(es catalog.ExchangeSchema) sc.SchemaObject {
	switch s := es.(type) {
	case *catalog.ExchangeJSightSchema:
		return sc.NewSchemaObject(s.JSchema)
	case catalog.ExchangeRegexSchema:
		return sc.NewSchemaObject(s.RSchema)
	case catalog.ExchangePseudoSchema:
		// should have been dealt with at the content level
		panic("cannot convert pseudo schema to SchemaObject at this level")
	default:
		panic("unsupported ExchangeSchema type provided")
	}
}

func SchemaObjectFromSchema(s schema.Schema) sc.SchemaObject {
	return sc.NewSchemaObject(s)
}

func SchemaObjectAny() sc.SchemaObject {
	return &schemaObjectAny{}
}

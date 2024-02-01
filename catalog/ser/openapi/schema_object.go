package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"

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
	// t.Schema is an ExchangeSchema iface nad must cast to a type
	switch s := es.(type) {
	case *catalog.ExchangeJSightSchema:
		return sc.NewSchemaObject(s.JSchema)
	case catalog.ExchangeRegexSchema:
		// make SchemaObject appr. for regex
		return sc.NewSchemaObject(s.RSchema)
	case catalog.ExchangePseudoSchema:
		// should have been dealt with at the content level
		panic("cannot convert pseudo schema to SchemaObject at this level")
	default:
		panic("unsupported ExchangeSchema type provided")
	}
}

func SchemaObjectFromSchemaKeeper[T sc.SchemaKeeper](sk T) sc.SchemaObject {
	return sc.NewSchemaObject(sk)
}

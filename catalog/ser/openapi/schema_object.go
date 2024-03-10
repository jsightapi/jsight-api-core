package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"
	schema "github.com/jsightapi/jsight-schema-core"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

type SchemaObject interface {
	sc.SchemaObject
}


/*
TODO:

	SchemaObject is the only structure allowed as a schema inside comonents
	so we must convert all of our notations into it.
	If some notation affects higher level structure of OA,
	this approach should be changed.
*/
// func SchemaObjectFromUserType(t *catalog.UserType) SchemaObject {
// 	// TODO: can we have UserType for any / empty ?
//
// 	return SchemaObjectFromExchangeSchema(t.Schema)
// }

// TODO: might become unnecessary
func SchemaObjectFromExchangeSchema(es catalog.ExchangeSchema) SchemaObject {

	debugExchangeSchema(es)

	switch es.Notation() {
	case notation.SchemaNotationJSight, notation.SchemaNotationRegex:
		return sc.NewSchemaObject(es)
	case notation.SchemaNotationAny:
		return SchemaObjectAny()
	case notation.SchemaNotationEmpty:
		panic("cannot convert notation any to SchemaObject")
	default:
		panic("unsupported schema notation")
	}
}

func SchemaObjectFromSchema(s schema.Schema) SchemaObject {
	return sc.NewSchemaObject(s)
}

func SchemaObjectAny() SchemaObject {
	return &schemaObjectAny{}
}

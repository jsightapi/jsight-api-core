package openapi

import (
	_ "fmt"

	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"
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
	// fmt.Printf("es type is %T\n", es)
	switch s := es.(type) {
	case *catalog.ExchangeJSightSchema:
		return sc.NewSchemaObject(s.JSchema)
	case *catalog.ExchangeRegexSchema: // TODO: may be a pointer here
		return sc.NewSchemaObject(s.RSchema)
	case *catalog.ExchangePseudoSchema:
		switch s.Notation {
		case notation.SchemaNotationAny:
			return SchemaObjectAny()
		case notation.SchemaNotationEmpty:
			// should have been dealt with at the content level
			panic("cannot convert pseudo schema to SchemaObject at this level")
		default:
			panic("unsupported notation provided")
		}
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

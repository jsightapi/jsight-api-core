package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

type SchemaObject interface {
	sc.SchemaObject
}

func schemaObjectFromExchangeSchema(es catalog.ExchangeSchema) SchemaObject {
	// debugExchangeSchema(es)

	switch es.Notation() {
	case notation.SchemaNotationJSight, notation.SchemaNotationRegex:
		return schemaObjectFromSchema(es)
	case notation.SchemaNotationAny:
		return schemaObjectForAny()
	case notation.SchemaNotationEmpty:
		panic("notation 'empty' cannot be represented by SchemaObject")
	default:
		panic("unsupported schema notation")
	}
}

func schemaObjectFromSchema(es catalog.ExchangeSchema) SchemaObject {
	switch s := es.(type) {
	case *catalog.ExchangeJSightSchema:
		return sc.NewSchemaObject(s.JSchema)
	case *catalog.ExchangeRegexSchema:
		return sc.NewSchemaObject(s.RSchema)
	default:
		panic("unsupported ExchangeSchema type for this method")
	}
}

func schemaObjectForAny() SchemaObject {
	return &schemaObjectAny{}
}

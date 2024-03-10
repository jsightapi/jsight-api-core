package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

// as part of the «components» object
type Schemas map[string]SchemaObject

func NewSchemas(tt *catalog.UserTypes) *Schemas {
	if tt.Len() == 0 {
		return nil
	}

	ss := make(Schemas, tt.Len())
	_ = tt.Each(func(name string, ut *catalog.UserType) error {
		ss[typeNameToSchemaName(name)] = SchemaObjectFromExchangeSchema(ut.Schema)
		return nil
	})

	return &ss
}

// all types in JSight start with `@`
func typeNameToSchemaName(n string) string {
	return n[1:]
}

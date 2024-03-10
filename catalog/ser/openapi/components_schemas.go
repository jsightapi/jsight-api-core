package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type ComponentsSchemas map[string]SchemaObject

func newSchemas(tt *catalog.UserTypes) ComponentsSchemas {
	if tt.Len() == 0 {
		return nil
	}

	ss := make(ComponentsSchemas, tt.Len())
	_ = tt.Each(func(name string, ut *catalog.UserType) error {
		ss[typeNameToSchemaName(name)] = schemaObjectFromExchangeSchema(ut.Schema)

		return nil
	})

	return ss
}

// all names in JSight start with `@`
func typeNameToSchemaName(n string) string {
	return n[1:]
}

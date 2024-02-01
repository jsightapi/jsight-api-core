package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

// as part of the «components» object
type Schemas map[string]sc.SchemaObject

func NewSchemas(tt *catalog.UserTypes) *Schemas {
	if tt.Len() == 0 {
		return nil
	}

	ss := make(Schemas, tt.Len())
	_ = tt.Each(func(name string, v *catalog.UserType) error {
		ss[convertTypeName(name)] = SchemaObjectFromUserType(v)
		return nil
	})

	return &ss
}

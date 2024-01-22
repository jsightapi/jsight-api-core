package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)
// as part of the «components» object
type Schemas map[string]*Schema

func NewSchemas(tt *catalog.UserTypes) *Schemas {
  if tt.Len() == 0 {
    return nil
  }

  ss := make(Schemas, tt.Len())
  tt.Each(func(k string, v *catalog.UserType) error {
    ss[k] = NewSchemaFromUserType(v)
    return nil
  })

  return &ss
}

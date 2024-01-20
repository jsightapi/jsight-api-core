package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Components struct {
  Schemas Schemas `json:"schemas,omitempty"`
}

func NewComponents(c *catalog.Catalog) *Components {
  if !hasComponents(c) {
    return nil
  }

  return &Components{
    Schemas: *NewSchemas(c.UserTypes),
  }
}

// TODO: think about other components. In JSS we have MACRO which is not convertable
func hasComponents(c *catalog.Catalog) bool {
	return c.UserTypes.Len() > 0 
}

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

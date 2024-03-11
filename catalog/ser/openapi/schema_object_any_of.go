package openapi

import (
	"encoding/json"
)

func schemaObjectForAnyOf(anyOf []SchemaObject, rootDescription string) SchemaObject {
	return &schemaObjectAnyOf{anyOf, rootDescription}
}

type schemaObjectAnyOf struct {
	AnyOf       []SchemaObject `json:"anyOf"`
	Description string         `json:"description,omitempty"`
}

func (s *schemaObjectAnyOf) SetDescription(d string) {
	s.Description = d
}

func (s schemaObjectAnyOf) MarshalJSON() (b []byte, err error) {
	type Alias schemaObjectAnyOf
	return json.Marshal(&struct {
		Alias
	}{
		Alias(s),
	})
}

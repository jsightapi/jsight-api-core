package openapi

import (
	"encoding/json"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

func SchemaObjectAnyOf(anyOf []sc.SchemaObject, rootDescription string) sc.SchemaObject {
	return &schemaObjectAnyOf{anyOf, rootDescription}
}

type schemaObjectAnyOf struct {
	AnyOf       []sc.SchemaObject `json:"anyOf"`
	Description string `json:"description,omitempty"`
}

func (s *schemaObjectAnyOf) SetDescription(d string) {
	s.Description = d
}

func (s schemaObjectAnyOf) MarshalJSON() (b []byte, err error) {
	return json.Marshal(s)
}

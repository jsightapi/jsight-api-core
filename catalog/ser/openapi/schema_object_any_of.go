package openapi

import (
	_ "encoding/json"

	_ "github.com/jsightapi/jsight-schema-core/openapi"
)

func SchemaObjectAnyOf(anyOf interface{}, rootDescription string) interface{} {
	return &schemaObjectAnyOf{anyOf, rootDescription}
}

type schemaObjectAnyOf struct {
	AnyOf       interface{} `json:"anyOf"`
	Description string `json:"description,omitempty"`
}

func (s *schemaObjectAnyOf) SetDescription(d string) {
	s.Description = d
}

// func (s schemaObjectAnyOf) MarshalJSON() (b []byte, err error) {
// 	return json.Marshal(s)
// }

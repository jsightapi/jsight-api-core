package openapi

import sc "github.com/jsightapi/jsight-schema-core/openapi"

/*
	Common fields for HeaderObject and ParameterObject

In OAS Header Object is officially a subset of Parameter Object
*/
type ParameterBase struct {
	Description     string          `json:"description,omitempty"`
	Required        bool            `json:"required"`
	Deprecated      bool            `json:"deprecated,omitempty"`
	AllowEmptyValue bool            `json:"allow_empty_value,omitempty"`
	Schema          sc.SchemaObject `json:"schema"`
}

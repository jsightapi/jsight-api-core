package openapi

import sc "github.com/jsightapi/jsight-schema-core/openapi"

type ParameterObject struct {
	ParameterBase
	Name string            `json:"name"`
	In   parameterLocation `json:"in"`
}

func NewParameterObject(
	in parameterLocation,
	name string,
	required bool,
	allowEmptyValue bool,
	schema sc.SchemaObject,
) *ParameterObject {
	return &ParameterObject{
		In:   in,
		Name: name,
		ParameterBase: ParameterBase{
			Required:        required,
			AllowEmptyValue: allowEmptyValue,
			Schema:          schema,
		},
	}
}

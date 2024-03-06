package openapi

import sc "github.com/jsightapi/jsight-schema-core/openapi"

type HeaderObject struct {
	ParameterBase
}

func NewHeaderObject(
	required bool, allowEmptyValue bool, description string, schema sc.SchemaObject,
) *HeaderObject {
	return &HeaderObject{
		ParameterBase{
			Required:        required,
			AllowEmptyValue: allowEmptyValue,
			Description:     description,
			Schema:          schema,
		}}
}

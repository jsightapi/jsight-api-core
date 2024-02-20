package openapi

import _ "github.com/jsightapi/jsight-schema-core/openapi"

type HeaderObject struct {
	ParameterBase
}

func NewHeaderObject(
	required bool, allowEmptyValue bool, description string, schema interface{},
) *HeaderObject {
	return &HeaderObject{
		ParameterBase{
			Required:        required,
			AllowEmptyValue: allowEmptyValue,
			Description:     description,
			Schema:          schema,
		}}
}

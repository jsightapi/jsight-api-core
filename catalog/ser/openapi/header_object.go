package openapi

type HeaderObject struct {
	ParameterBase
}

func NewHeaderObject(
	required bool, allowEmptyValue bool, description string, schema SchemaObject,
) *HeaderObject {
	return &HeaderObject{
		ParameterBase{
			Required:        required,
			AllowEmptyValue: allowEmptyValue,
			Description:     description,
			Schema:          schema,
		}}
}

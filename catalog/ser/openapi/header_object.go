package openapi

type HeaderObject struct {
	ParameterBase
}

func newHeaderObject(
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

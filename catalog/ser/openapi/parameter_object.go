package openapi

type ParameterObject struct {
	ParameterBase
	Name string            `json:"name"`
	In   parameterLocation `json:"in"`
}

// TODO: do we need a constructor?
func newParameterObject(
	in parameterLocation,
	name string,
	required bool,
	schema SchemaObject,
) *ParameterObject {
	return &ParameterObject{
		In:   in,
		Name: name,
		ParameterBase: ParameterBase{
			Required:        required,
			Schema:          schema,
		},
	}
}

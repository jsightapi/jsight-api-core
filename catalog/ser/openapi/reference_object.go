package openapi

type Reference string

const (
	refBaseSchema string = "#/components/schemas/"
)

type ReferenceObject struct {
	Ref Reference `json:"$ref"`
}

func NewReferenceObject(typeName string) *ReferenceObject {
	return &ReferenceObject{
		schemaReference(convertTypeName(typeName)),
	}
}

func schemaReference(name string) Reference {
	return Reference(refBaseSchema + name)
}

package openapi

import schema "github.com/jsightapi/jsight-schema-core"

import sc "github.com/jsightapi/jsight-schema-core/openapi"

// TODO: just a draft
type parametersInfo interface {
	Required() bool
	Description() string
	AllowEmptyValue() bool
	Schema() sc.SchemaObject
}

type paramInfo struct{}

func (p paramInfo) Required() bool {
	return true
}

func (p paramInfo) AllowEmptyValue() bool {
	return true
}

func (p paramInfo) Description() string {
	return "dummy Description"
}

func (p paramInfo) Schema() sc.SchemaObject {
	return nil
}

func GetParamInfo(schema.ASTNode) *paramInfo {
	return &paramInfo{}
}

// TODO: just a demo
type schemaInfo interface {
	Optional() bool
	Annptation() string
	Nullable() bool
	ASTNode() sc.SchemaObject
}

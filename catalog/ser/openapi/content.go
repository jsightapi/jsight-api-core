package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	schema "github.com/jsightapi/jsight-schema-core"
)

type Content map[mediaType]*MediaTypeObject

func defaultContent() *Content {
	return ContentForAny()
}

func ContentForAny() *Content {
	c := make(Content, 1)
	c[MediaTypeRangeAny] = &MediaTypeObject{
		Schema: SchemaObjectAny(),
	}
	return &c
}

func ContentForEmpty() *Content {
	c := make(Content, 0)
	return &c
}

func ContentWithMediaTypeObject(mt mediaType, o *MediaTypeObject) *Content {
	c := make(Content)
	c[mt] = o
	return &c
}

func NewContent(f catalog.SerializeFormat, s catalog.ExchangeSchema) *Content {
	c := make(Content)
	mt := FormatToMediaType(f)
	c[mt] = NewMediaTypeObjectFromExchangeSchema(s)
	return &c
}

func ContentForAnyOf(schemaObjects []SchemaObject) *Content {
	c := make(Content, 1)
	c[MediaTypeRangeAny] = &MediaTypeObject{SchemaObjectAnyOf(schemaObjects, "")}
	return &c
}

func NewContentFromSchema(f catalog.SerializeFormat, s schema.Schema) *Content {
	c := make(Content)
	mt := FormatToMediaType(f)
	c[mt] = NewMediaTypeObjectFromSchema(s)
	return &c
}

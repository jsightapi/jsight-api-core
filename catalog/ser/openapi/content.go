package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

type Content map[mediaType]*MediaTypeObject

func defaultContent() *Content {
	return ContentForAny()
}

func NewContent(f catalog.SerializeFormat, s catalog.ExchangeSchema) *Content {
	c := make(Content)
	mt := FormatToMediaType(f)
	c[mt] = NewMediaTypeObject(s)
	return &c
}

func ContentForAny() *Content {
	c := make(Content, 1)
	c[MediaTypeRangeAny] = &MediaTypeObject{}
	return &c
}

func ContentForEmpty() *Content {
	c := make(Content, 0)
	return &c
}

func NewContentFromSchemaKeeper[T sc.SchemaKeeper](f catalog.SerializeFormat, sk T) *Content {
	c := make(Content)
	mt := FormatToMediaType(f)
	c[mt] = NewMediaTypeObjectFromSchemaKeeper(sk)
	return &c
}

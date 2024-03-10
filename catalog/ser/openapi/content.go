package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	// schema "github.com/jsightapi/jsight-schema-core"
)

/* Content is used in Responses and Requests
 */
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

/* JSight's pseudo-notation empty is expressed via empty OA's content object. 
*/
func ContentForEmpty() *Content {
	c := make(Content, 0)
	return &c
}

func ContentForAnyOf(schemaObjects []SchemaObject) *Content {
	c := make(Content, 1)
	c[MediaTypeRangeAny] = &MediaTypeObject{SchemaObjectAnyOf(schemaObjects, "")}
	return &c
}

func ContentForSchema(f catalog.SerializeFormat, es catalog.ExchangeSchema) *Content {
	c := make(Content)
	mt := FormatToMediaType(f)
	c[mt] = MediaTypeObjectForSchema(es)
	return &c
}

// func ContentWithMediaTypeObject(mt mediaType, o *MediaTypeObject) *Content {
// 	c := make(Content)
// 	c[mt] = o
// 	return &c
// }

// func NewContent(f catalog.SerializeFormat, s catalog.ExchangeSchema) *Content {
// 	c := make(Content)
// 	mt := FormatToMediaType(f)
// 	c[mt] = NewMediaTypeObjectFromExchangeSchema(s)
// 	return &c
// }

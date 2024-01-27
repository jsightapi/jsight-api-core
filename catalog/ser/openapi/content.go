package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Content map[mediaType]*MediaTypeObject

/*
	TODO:

Can content contain more than one media type? Which cases?
*/
func NewContent(r *catalog.HTTPRequest) *Content {
	c := make(Content)

	mt := FormatToMediaType(r.Format)

	c[mt] = NewMediaTypeObject(r.HTTPRequestBody.Schema)

	return &c
}

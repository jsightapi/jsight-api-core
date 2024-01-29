package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

// TODO: make one constructor when var cases solved

type Content map[mediaType]*MediaTypeObject

/*
	TODO:

Can content contain more than one media type? Which cases?
*/
func NewContentFromRequest(r *catalog.HTTPRequest) *Content {
	c := make(Content)

	mt := FormatToMediaType(r.Format)

	c[mt] = NewMediaTypeObject(r.HTTPRequestBody.Schema)

	return &c
}

/*
	TODO:

Can content contain more than one media type? Which cases?
*/
func NewContentFromResponse(r *catalog.HTTPResponse) *Content {
	c := make(Content)

	mt := FormatToMediaType(r.Body.Format)

	c[mt] = NewMediaTypeObject(r.Body.Schema)

	return &c
}

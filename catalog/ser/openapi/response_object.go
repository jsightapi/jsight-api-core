package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type ResponseObject struct {
	Description string   `json:"description"`
	Headers     *Headers `json:"headers"`
	Content     *Content `json:"content,omitempty"`
}

func NewResponse(r *catalog.HTTPResponse) *ResponseObject {
	return &ResponseObject{
		Description: r.Annotation,
		Content:     NewContent(r.Body.Format, r.Body.Schema),
		Headers:     NewHeaders(r.Headers),
	}
}

func defaultResponse() *ResponseObject {
	return &ResponseObject{
		Description: "",
		Content:     defaultContent(),
	}
}

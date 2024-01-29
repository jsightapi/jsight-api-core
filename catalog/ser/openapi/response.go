package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Response struct {
	Description string `json:"description"`
	// Headers *Headers
	Content *Content `json:"content,omitempty"`
	// Links // irrelevant to JSight
}

func NewResponse(resp *catalog.HTTPResponse) *Response {
	return &Response{
		Description: "", // TODO
		Content:     NewContentFromResponse(resp),
	}
}

func defaultResponses() *Responses {
	r := make(Responses, 1)
	r["default"] = defaultResponse()
	return &r
}

package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Responses map[responseCode]*Response

type responseCode string

func NewResponses(i *catalog.HTTPInteraction) *Responses {
	if len(i.Responses) == 0 {
		return defaultResponses()
	}

	r := make(Responses, 1)
	return &r
}

func defaultResponses() *Responses {
	r := make(Responses, 1)
	r["default"] = defaultResponse()
	return &r
}

func defaultResponse() *Response {
	return &Response{
		Description: "",
		Content:     defaultContent(),
	}
}

func defaultContent() *Content {
	c := make(Content, 1)

	c["*/*"] = &MediaTypeObject{}

	return &c
}

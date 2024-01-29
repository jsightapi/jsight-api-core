package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Responses map[responseCode]*Response

type responseCode string

func NewResponses(i *catalog.HTTPInteraction) *Responses {
	if len(i.Responses) == 0 {
		return defaultResponses()
	}

	r := make(Responses, 1)

	for idx, resp := range i.Responses { // TODO: Order ?
		rc := responseCode(resp.Code)
		if _, exists := r[rc]; exists {
			// if not in a map – add
		} else {
			// if in a map – rebuild in a
			r[rc] = NewResponse(&i.Responses[idx])
		}
	}

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

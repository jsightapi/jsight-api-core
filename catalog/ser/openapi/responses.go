package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Responses map[responseCode]*Response

type responseCode string

func NewResponses(i *catalog.HTTPInteraction) *Responses {
	if len(i.Responses) == 0 {
		return defaultResponses()
	}

	r := make(Responses, 1)

	for idx := range i.Responses {
		resp := &i.Responses[idx] // safer way to get a pointer in this case
		rc := responseCode(resp.Code)
		if _, exists := r[rc]; exists {
			r[rc] = MergeResponse(r[rc], resp)
		} else {
			r[rc] = NewResponse(resp)
		}
	}

	return &r
}

func MergeResponse(r *Response, add *catalog.HTTPResponse) *Response {
	// TODO: merge

	return nil
}

func NewResponse(resp *catalog.HTTPResponse) *Response {
	return &Response{
		Description: resp.Annotation,
		Content:     NewContentFromResponse(resp),
	}
}

func defaultResponses() *Responses {
	r := make(Responses, 1)
	r["default"] = defaultResponse()
	return &r
}

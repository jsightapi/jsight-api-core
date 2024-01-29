package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Responses map[responseCode]*Response

type responseCode string

func NewResponses(i *catalog.HTTPInteraction) *Responses {
	if len(i.Responses) == 0 {
		return defaultResponses()
	}

	r := make(Responses, 1)

	for idx, _ := range i.Responses { // TODO: Order ?
		resp := &i.Responses[idx]
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
		Description: "", // TODO
		Content:     NewContentFromResponse(resp),
	}
}

func defaultResponses() *Responses {
	r := make(Responses, 1)
	r["default"] = defaultResponse()
	return &r
}

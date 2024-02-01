package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

type Responses map[responseCode]*ResponseObject

type responseCode string

func NewResponses(i *catalog.HTTPInteraction) *Responses {
	if len(i.Responses) == 0 {
		return defaultResponses()
	}

	r := make(Responses, 1)

	for idx := range i.Responses {

		// TODO: rewrite

		// sort responses to map[code][]Response.
		// range this map
		// if map[code] len > 1 , then merger responses into one
		// else – make response

		resp := &i.Responses[idx] // safer way to get a pointer in this case
		rc := responseCode(resp.Code)
		if _, exists := r[rc]; exists {
			r[rc] = MergeResponseObjects(r[rc], resp)
		} else {
			r[rc] = NewResponse(resp)
		}
	}

	return &r
}

func MergeResponseObjects(r *ResponseObject, add *catalog.HTTPResponse) *ResponseObject {
	// TODO: merge

	// take r Schema f@

	//

	return nil
}

func defaultResponses() *Responses {
	r := make(Responses, 1)
	r["default"] = defaultResponse()
	return &r
}

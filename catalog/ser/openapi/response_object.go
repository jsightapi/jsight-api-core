package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"
)

type ResponseObject struct {
	Description string          `json:"description"`
	Headers     ResponseHeaders `json:"headers,omitempty"`
	Content     *Content        `json:"content,omitempty"`
}

func defaultResponse() *ResponseObject {
	return &ResponseObject{
		Description: "",
		Content:     defaultContent(),
	}
}

func newResponse(r *catalog.HTTPResponse) *ResponseObject {
	return &ResponseObject{
		Description: r.Annotation,
		Content:     contentForSchema(r.Body.Format, r.Body.Schema),
		Headers:     makeResponseHeaders(r.Headers),
	}
}

func newResponseAnyOf(responses []*catalog.HTTPResponse) *ResponseObject {
	hh := make([]*catalog.HTTPResponseHeaders, 0)
	sos := make(map[mediaType][]schemaObject, 0)

	for _, response := range responses {
		hh = append(hh, response.Headers)
		respAnnotation := response.Annotation

		var so schemaObject
		var desc string
		var mt mediaType
		if response.Body == nil {
			so = schemaObjectForAny()
			desc = respAnnotation
			mt = MediaTypeRangeAny
		} else {
			s := response.Body.Schema
			switch s.Notation() {
			case notation.SchemaNotationJSight:
				si := getJSchemaInfo(s.(*catalog.ExchangeJSightSchema).JSchema)
				so = si.SchemaObject()
				desc = concatenateDescription(respAnnotation, si.Annotation())
			case notation.SchemaNotationRegex: 
				si := getRSchemaInfo(s.(*catalog.ExchangeRegexSchema).RSchema)
				so = si.SchemaObject()
				desc = concatenateDescription(respAnnotation, si.Annotation())
			case notation.SchemaNotationAny:
				so = schemaObjectForAny()
				desc = respAnnotation
			case notation.SchemaNotationEmpty:
				// TODO: ???
				panic("TODO: empty response body in same-code responses: not decided")
			}
			mt = formatToMediaType(response.Body.Format)
		}
		so.SetDescription(desc)
		sos[mt] = append(sos[mt], so)
	}

	return &ResponseObject{
		Headers: makeResponseHeaders(hh...),
		Content: contentForVariousMediaTypes(sos),
	}
}

package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/notation"

	// "github.com/jsightapi/jsight-schema-core/notations/jschema"
	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

type ResponseObject struct {
	Description string   `json:"description"`
	Headers     Headers  `json:"headers,omitempty"`
	Content     *Content `json:"content,omitempty"`
}

func defaultResponse() *ResponseObject {
	return &ResponseObject{
		Description: "",
		Content:     defaultContent(),
	}
}

func NewResponse(r *catalog.HTTPResponse) *ResponseObject {
	return &ResponseObject{
		Description: r.Annotation,
		Content:     NewContent(r.Body.Format, r.Body.Schema),
		Headers:     NewHeaders(r.Headers),
	}
}

func NewResponseAnyOf(responses []*catalog.HTTPResponse) *ResponseObject {
	hh := make([]*catalog.HTTPResponseHeaders, 0)
	sos := make([]sc.SchemaObject, 0)

	for _, response := range responses {
		hh = append(hh, response.Headers)

		respAnnotation := response.Body.Directive.Annotation

		var so sc.SchemaObject
		var desc string
		if response.Body == nil {
			so = SchemaObjectAny()
			so.SetDescription(respAnnotation)
		} else {
			s := response.Body.Schema
			switch s.Notation() {
			case notation.SchemaNotationJSight:
				si := sc.NewSchemaInfo(s.(*catalog.ExchangeJSightSchema).JSchema)
				so = si.SchemaObject()
				desc = concatenateDescription(respAnnotation, si.Annotation())
			case notation.SchemaNotationRegex:
				so = sc.NewSchemaObject(s)
				desc = respAnnotation
			case notation.SchemaNotationAny:
				so = SchemaObjectAny()
				desc = respAnnotation
			case notation.SchemaNotationEmpty:
				// TODO: probably needs to be ignored in any of.
				//  if there are any valid schemas then adding an empty means nothing
			}

			so.SetDescription(desc)
			sos = append(sos, so)
		}
	}

	mto := MediaTypeObject{
		SchemaObjectAnyOf(sos, ""),
	}

	return &ResponseObject{
		Headers: MakeResponseHeaders(hh),
		Content: ContentWithMediaTypeObject(MediaTypeRangeAny, &mto),
	}
}

func responsesToSchemas(rr []*catalog.HTTPResponse) []catalog.ExchangeSchema {
	if rr == nil {
		return nil
	}

	ss := make([]catalog.ExchangeSchema, len(rr))
	for _, r := range rr {
		ss = append(ss, r.Body.Schema)
	}
	return ss
}


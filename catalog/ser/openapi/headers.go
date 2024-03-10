package openapi

import (
	 _ "fmt"

	"github.com/jsightapi/jsight-api-core/catalog"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

// Only for Response objects. For request headers refer to "parameters"
type ResponseHeaders map[string]*HeaderObject

// TODO: may not be a schema, may be a literal | array
// func NewHeaders(h *catalog.HTTPResponseHeaders) ResponseHeaders {
// 	if h == nil {
// 		return nil
// 	}
//
// 	r := make(ResponseHeaders, 0)
//
// 	// TODO: children may be nil if schema is literal
// 	for _, ch := range h.Schema.ASTNode.Children {
// 		pi := GetParamInfo(ch)
// 		r[ch.Key] = NewHeaderObject(pi.Required(), pi.AllowEmptyValue(), pi.Description(), pi.Schema())
// 	}
//
// 	return r
// }
//
type headerInfo struct {
	schemaInfo        sc.SchemaInfo
	contextAnnotation string
}

func MakeResponseHeaders(headersArr ...*catalog.HTTPResponseHeaders) ResponseHeaders {
	r := make(ResponseHeaders, 0)

	sortedHeaders := make(map[string][]headerInfo)
	for _, headers := range headersArr {
		if headers == nil {
			continue
		}

		headersSchemaInfo := GetSchemaInfo(headers.Schema.JSchema)
		propIterator := headersSchemaInfo.PropertiesInfos()
		for propIterator.Next() {
			hName := propIterator.GetKey()
			sortedHeaders[hName] = append(sortedHeaders[hName],
				headerInfo{
					propIterator.GetInfo(),
					headers.Directive.Annotation,
				},
			)
		}
	}

	for name, headerInfos := range sortedHeaders {
		if len(headerInfos) == 1 {
			i := headerInfos[0]
			r[name] = NewHeaderObject(
				!i.schemaInfo.Optional(), false,
				concatenateDescription(i.contextAnnotation, i.schemaInfo.Annotation()),
				i.schemaInfo.SchemaObject(),
			)
		} else {
			r[name] = HeaderObjectForAnyOf(headerInfos)
		}
	}

	// fmt.Printf("\n headers len: %d", len(r))

	return r
}

// // TODO: use strings.go
// func combineHeaderDescription(context string, propertyAnnotation string) string {
// 	return context + ": " + propertyAnnotation
// }

func HeaderObjectForAnyOf(headersInfos []headerInfo) *HeaderObject {
	schemaObjects := make([]SchemaObject, 0)
	var required bool

	for _, i := range headersInfos {
		if i.schemaInfo.Optional() {
			required = false
		}

		so := i.schemaInfo.SchemaObject()
		so.SetDescription(concatenateDescription(i.contextAnnotation, i.schemaInfo.Annotation()))
		schemaObjects = append(schemaObjects, so)
	}

	return NewHeaderObject(
		required,
		false,
		"",
		&schemaObjectAnyOf{schemaObjects, ""},
	)
}

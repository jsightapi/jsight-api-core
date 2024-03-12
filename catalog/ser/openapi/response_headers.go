package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"

	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

// Only for Response objects. For request headers refer to "parameters"
type ResponseHeaders map[string]*HeaderObject

type headerInfo struct {
	schemaInfo        sc.SchemaInfo
	contextAnnotation string
}

func makeResponseHeaders(headersArr ...*catalog.HTTPResponseHeaders) ResponseHeaders {
	r := make(ResponseHeaders, 0)

	sortedHeaders := make(map[string][]headerInfo)
	for _, headers := range headersArr {
		if headers == nil {
			continue
		}

		headersSchemaInfo := getSchemaInfo(headers.Schema.JSchema)
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
			r[name] = newHeaderObject(
				!i.schemaInfo.Optional(),
				concatenateDescription(i.contextAnnotation, i.schemaInfo.Annotation()),
				i.schemaInfo.SchemaObject(),
			)
		} else {
			r[name] = headerObjectForAnyOf(headerInfos)
		}
	}
	return r
}

func headerObjectForAnyOf(headersInfos []headerInfo) *HeaderObject {
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

	return newHeaderObject(
		required, "",
		&schemaObjectAnyOf{schemaObjects, ""},
	)
}

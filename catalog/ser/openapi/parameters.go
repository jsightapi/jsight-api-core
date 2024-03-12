package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

// TODO: cleanup constructor based on use cases
// TODO: if schema is a reference?
// if a reference to Regex
func paramsFromJSchema(es *catalog.ExchangeJSightSchema, in parameterLocation) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	if es == nil {
		return r
	}

	schemaInfo := getSchemaInfo(es.JSchema)
	propIterator := schemaInfo.PropertiesInfos()
	for propIterator.Next() {
		pi := propIterator.GetInfo()
		po := newParameterObject(
			in,
			propIterator.GetKey(),
			pi.Annotation(),
			!pi.Optional(),
			pi.SchemaObject(),
		)
		r = append(r, po)
	}
	return r
}

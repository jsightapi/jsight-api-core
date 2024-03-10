package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

func paramsFromSchema(es *catalog.ExchangeJSightSchema, loc parameterLocation) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	if es == nil {
		return r
	}

	schemaInfo := getSchemaInfo(es.JSchema)
	propIterator := schemaInfo.PropertiesInfos()
	for propIterator.Next() {
		pi := propIterator.GetInfo()
		po := newParameterObject(
			loc,
			propIterator.GetKey(),
			pi.Optional(),
			false, // TODO: ?
			pi.SchemaObject(),
		)
		r = append(r, po)
	}
	return r
}

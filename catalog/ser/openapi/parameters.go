package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

func ParamsFromSchema(es *catalog.ExchangeJSightSchema, loc parameterLocation) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	if es == nil {
		return r
	}

	schemaInfo := GetSchemaInfo(es.JSchema)
	propIterator := schemaInfo.PropertiesInfos()
	for propIterator.Next() {
		pi := propIterator.GetInfo()
		po := NewParameterObject(
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

package openapi

import (
	// "fmt"

	"github.com/jsightapi/jsight-api-core/catalog"
)

func paramsFromSchema(es *catalog.ExchangeJSightSchema, loc parameterLocation) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	if es == nil {
		return r
	}

	schemaInfo := getSchemaInfo(es.JSchema)
	propIterator := schemaInfo.PropertiesInfos()
	for propIterator.Next() {
		// pk := propIterator.GetKey()
		pi := propIterator.GetInfo()

		// fmt.Printf("key: %s\n annot: %s\n opt: %t\n so: %s\n\n",
			// pk, pi.Annotation(), pi.Optional(), schemaObjectToString(pi.SchemaObject()))

		po := newParameterObject(
			loc,
			propIterator.GetKey(),
			!pi.Optional(),
			pi.SchemaObject(),
		)
		r = append(r, po)
	}
	return r
}

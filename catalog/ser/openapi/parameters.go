package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

// TODO: cleanup constructor based on use cases
// TODO: if schema is a reference?
// if a reference to Regex
func paramsFromJSchema(es *catalog.ExchangeJSightSchema, in parameterLocation) []*ParameterObject {
	r := make([]*ParameterObject, 0)
	if es == nil {
		return r
	}

	realSchemas := dereferenceJSchema(es.JSchema)
	if len(realSchemas) == 1 {
		rs := realSchemas[0]
		switch rs.Type() {
		case sc.SchemaInfoTypeObject: // TODO: get rid of sc. import?
			properties := rs.(sc.ObjectInformer).PropertiesInfos()
			for _, pi := range properties {
				po := newParameterObject(
					in,
					pi.Key(),
					pi.Annotation(),
					!pi.Optional(),
					pi.SchemaObject(),
				)
				r = append(r, po)
			}
		default:
			panic("parameters directive's schema is not an object")
		}
	} else {
			panic("or-references conversion not supported for parameter directives")
	}
	return r
}

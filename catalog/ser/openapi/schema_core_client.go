package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	"github.com/jsightapi/jsight-schema-core/notations/regex"
	sc "github.com/jsightapi/jsight-schema-core/openapi"
)


type SchemaObject interface {
	sc.SchemaObject
}

type schemaPropertyInfo interface {
	sc.PropertyInformer
}

type schemaObjectInfo interface {
	sc.ObjectInformer
}

type schemaInfo interface {
	sc.SchemaInformer
}

func getSchemaObjectInfo(s *jschema.JSchema) schemaObjectInfo {
	sd := dereferenceJSchema(s)
	if len(sd) > 1 {
		panic("or-references not supported")
	} else {
		ei := sd[0]
		if ei.Type() == sc.SchemaInfoTypeObject {
			return ei.(sc.ObjectInformer)
		} else {
			panic("schema is not an object")
		}
	}
}

func getJSchemaInfo(s *jschema.JSchema) schemaInfo {
	return sc.NewJSchemaInfo(s)
}

func getRSchemaInfo(s *regex.RSchema) schemaInfo {
	return sc.NewRSchemaInfo(s)
}

func dereferenceJSchema(s *jschema.JSchema) []sc.SchemaInformer {
	return sc.Dereference(s)
}

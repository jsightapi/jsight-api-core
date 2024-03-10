package openapi

import (
	"github.com/jsightapi/jsight-schema-core/notations/jschema"
	sc "github.com/jsightapi/jsight-schema-core/openapi"
)

func getSchemaInfo(s *jschema.JSchema) sc.SchemaInfo {
	return sc.NewSchemaInfo(s)
}

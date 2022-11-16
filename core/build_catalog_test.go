package core

import (
	"testing"

	schema "github.com/jsightapi/jsight-schema-core"
	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-api-core/catalog"
	"github.com/jsightapi/jsight-api-core/directive"
)

func TestJApiCore_compileUserTypes(t *testing.T) {
	c := catalog.NewCatalog()

	ut := map[string]string{
		"@foo": "42",
		"@bar": `{
	"foo": @foo
}`,
		"@fizz": `{
		"bar": @bar
}`,
		"@buzz": `{
		"fizz": @fizz
}`,
	}

	core := JApiCore{
		rawUserTypes:       &directive.Directives{},
		userTypes:          &catalog.UserSchemas{},
		processedUserTypes: map[string]struct{}{},
		catalog:            c,
	}

	for n, p := range ut {
		d := directive.New(directive.Jsight, directive.Coords{})
		d.BodyCoords = directive.NewCoords(
			fs.NewFile(n, p),
			0,
			bytes.Index(len(p)-1),
		)
		require.NoError(t, d.SetNamedParameter("Name", n))
		core.AddRawUserType(d)
	}

	err := core.compileUserTypes()
	require.Nil(t, err)

	_ = core.userTypes.Each(func(k string, v schema.Schema) error {
		require.NoErrorf(t, v.Check(), "Check %q user type", k)
		return nil
	})
}

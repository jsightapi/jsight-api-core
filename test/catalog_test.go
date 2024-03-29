package test

import (
	"path/filepath"
	"testing"

	"github.com/jsightapi/jsight-schema-core/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-api-core/core"
)

func TestGetAllTypesSchemas(t *testing.T) {
	filename := filepath.Join(GetTestDataDir(), "jsight_0.3", "others", "full.jst")
	f := reader.Read(filename)
	japi := core.NewJApiCore(f)
	err := japi.BuildCatalog()
	require.Nil(t, err)

	c := japi.Catalog()
	types := c.UserTypes

	assert.Equal(t, 5, types.Len())

	assert.True(t, types.Has("@error"))
	assert.True(t, types.Has("@user"))
	assert.True(t, types.Has("@profile"))
	assert.True(t, types.Has("@task"))
	assert.True(t, types.Has("@attachment"))
	assert.True(t, c.UserEnums.Has("@userType"))
}

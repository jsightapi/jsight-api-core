package core

import (
	"testing"

	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-api-core/directive"
)

func TestWithBannedDirectives(t *testing.T) {
	err := NewJApiCore(fs.NewFile("", "JSIGHT 0.3"), WithBannedDirectives(directive.Jsight)).BuildCatalog()
	assert.EqualError(t, err, "directive not allowed (JSIGHT)")
}

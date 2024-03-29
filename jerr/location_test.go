package jerr

import (
	"testing"

	"github.com/jsightapi/jsight-schema-core/bytes"
	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/stretchr/testify/assert"
)

func TestNewLocation(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		const index = bytes.Index(3)
		file := fs.NewFile("foo", "12\n34\n56")

		l := NewLocation(file, index)

		assert.Same(t, file, l.File)

		assert.Equal(t, index, l.Index)

		assert.Equal(t, "34", l.Quote)

		assert.Equal(t, bytes.Index(2), l.Line)
	})

	t.Run("negative", func(t *testing.T) {
		assert.Panics(t, func() {
			NewLocation(nil, 0)
		})
	})
}

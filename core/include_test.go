package core

import (
	"testing"

	"github.com/jsightapi/jsight-schema-core/fs"
	"github.com/stretchr/testify/assert"

	"github.com/jsightapi/jsight-api-core/scanner"
)

func Test_isIncludeKeyword(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cc := map[string]struct {
			given    *scanner.Lexeme
			expected bool
		}{
			"nil": {},
			"not a keyword": {
				given: scanner.NewLexeme(scanner.Schema, 0, 0, nil),
			},
			"not an include keyword": {
				given: scanner.NewLexeme(scanner.Keyword, 0, 1, fs.NewFile("", "12")),
			},
			"an include keyword": {
				given:    scanner.NewLexeme(scanner.Keyword, 0, 6, fs.NewFile("", "INCLUDE")),
				expected: true,
			},
		}

		for n, c := range cc {
			t.Run(n, func(t *testing.T) {
				actual := isIncludeKeyword(c.given)
				assert.Equal(t, c.expected, actual)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		assert.Panics(t, func() {
			isIncludeKeyword(&scanner.Lexeme{})
		})
	})
}

func Test_validateIncludeFileName(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		ss := []string{
			"foo",
			".foo",
			"foo/bar",
			".foo/bar",
			"..foo/bar",
			"foo/bar.jst",
		}

		for _, s := range ss {
			t.Run(s, func(t *testing.T) {
				err := validateIncludeFileName(s)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		cc := map[string]string{
			"/foo/bar":   "cannot not start with `/`",
			"/./bar":     "cannot not start with `/`",
			"/../bar":    "cannot not start with `/`",
			"./foo/bar":  "cannot contain `..` or `.`",
			"../foo/bar": "cannot contain `..` or `.`",
			"foo/.":      "cannot contain `..` or `.`",
			"foo/./":     "cannot contain `..` or `.`",
			"foo/./bar":  "cannot contain `..` or `.`",
			"foo/..":     "cannot contain `..` or `.`",
			"foo/../":    "cannot contain `..` or `.`",
			"foo/../bar": "cannot contain `..` or `.`",
			"foo/./bar/../fizz/../buzz/./fizzbuzz/..": "cannot contain `..` or `.`",
			"foo\\bar":           "directories must be separated by slashes `/`",
			"foo/bar\\fizz/buzz": "directories must be separated by slashes `/`",
		}

		for given, expected := range cc {
			t.Run(given, func(t *testing.T) {
				err := validateIncludeFileName(given)
				assert.EqualError(t, err, expected)
			})
		}
	})
}

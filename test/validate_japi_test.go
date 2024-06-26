package test

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-api-core/jerr"
	"github.com/jsightapi/jsight-api-core/kit"
)

const errorIndexIsNotSpecified = "the error index is not specified in the file name"

func TestValidateJapi(t *testing.T) {
	filenames := jstFilenames(GetTestDataDir())

	t.Run("positive", func(t *testing.T) {
		positive := positiveJstFilenames(filenames)
		for _, f := range positive {
			t.Run(cutRepositoryPath(f), func(t *testing.T) {
				_, je := kit.NewJapi(f)
				if je != nil {
					logJAPIError(t, je)
					t.FailNow()
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		negative := negativeJstFilenames(filenames)
		for _, f := range negative {
			t.Run(cutRepositoryPath(f), func(t *testing.T) {
				_, je := kit.NewJapi(f)
				require.NotNil(t, je)

				logJAPIError(t, je)

				want, err := wantIndex(f)
				require.NoError(t, err)

				assert.Equal(t, want, je.Index.Int())

				expectedError, err := getExpectedError(f)
				require.NoError(t, err)

				// We print absolute path to file in the stacktrace. This leads
				// to errors when somebody wants to run test on different machine.
				// So, we should use relative paths in the expected error and make
				// all path relative in the actual error message.
				if expectedError != "" {
					sep := string(os.PathSeparator)
					basePath := strings.TrimSuffix(filepath.Dir(f), sep) + sep
					actualError := getActualErrorMessage(basePath, je)
					assert.Equal(t, expectedError, actualError)
				}
			})
		}
	})
}

func cutRepositoryPath(p string) string {
	p, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(p, string(filepath.Separator))
	var idx int
	for _, p := range parts {
		idx++
		if p == "testdata" {
			break
		}
	}

	return filepath.Join(parts[idx:]...)
}

func wantIndex(filename string) (int, error) {
	base := filepath.Base(filename)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	s := strings.Split(name, "_")
	if len(s) < 2 {
		return 0, errors.New(errorIndexIsNotSpecified)
	}
	i, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, errors.New(errorIndexIsNotSpecified)
	}
	return i, nil
}

func getExpectedError(filename string) (string, error) {
	filename = strings.TrimSuffix(filename, ".jst") + ".error"
	c, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return "", err
	}
	return strings.TrimSpace(string(c)), nil
}

func getActualErrorMessage(basePath string, je *jerr.JApiError) string {
	if je == nil {
		return ""
	}
	// get rid of base path file paths for windows paths (considering `\` slashes)
	escapedPath := strings.ReplaceAll(basePath, `\`, `\\`)
	escapedPath = strings.ReplaceAll(escapedPath, `.`, `\.`)
	m := regexp.MustCompile(`(` + escapedPath + `)(([^\s\\]+)\\)+`)
	r := m.ReplaceAllStringFunc(je.Error(), func(s string) string {
		r := strings.ReplaceAll(s, basePath, "")
		return strings.ReplaceAll(r, `\`, `/`)
	})
	// get rid of basePath in linux file paths
	r = strings.ReplaceAll(r, basePath, "")
	return r
}

func logJAPIError(t *testing.T, e *jerr.JApiError) {
	t.Log("Got:")
	t.Log("- Line: " + e.Line.String())
	t.Log("- Index: " + e.Index.String())
	t.Log("- Message: " + e.Error())
	t.Log("- Quote: " + e.Quote)
}

func jstFilenames(dir string) []string {
	filenames := make([]string, 0, 30)

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			base := filepath.Base(path)

			if info.IsDir() && (base == ".unused" || base == "scanner" || base == "mixins") {
				return filepath.SkipDir
			}

			if !info.IsDir() && filepath.Ext(path) == ".jst" {
				filenames = append(filenames, path)
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	return filenames
}

func positiveJstFilenames(filenames []string) []string {
	list := make([]string, 0, 10)
	for _, f := range filenames {
		if !strings.HasPrefix(filepath.Base(f), "err_") {
			list = append(list, f)
		}
	}
	return list
}

func negativeJstFilenames(filenames []string) []string {
	list := make([]string, 0, 10)
	for _, f := range filenames {
		if strings.HasPrefix(filepath.Base(f), "err_") {
			list = append(list, f)
		}
	}
	return list
}

package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-api-core/kit"
)

func TestOpenAPI(t *testing.T) {
	openapiFilesPaths := openapiFilePaths(GetTestDataDir())

	for _, openapiPath := range openapiFilesPaths {
		t.Run(cutRepositoryPath(openapiPath), func(t *testing.T) {
			json, err := os.ReadFile(openapiPath)
			require.NoError(t, err)

			japiPath, err := japiFilePath(openapiPath)
			require.NoError(t, err)

			// TODO: throw error if file does not exist

			j, je := kit.NewJapi(japiPath)
			if je != nil {
				logJAPIError(t, je)
				t.FailNow()
			}

			actual, err := j.ToOpenAPIJsonIndent()
			require.NoError(t, err)

			expected := string(json)

			ok := assert.JSONEq(t, expected, string(actual))

			if !ok {
				t.Log("Actual OpenAPI JSON:")
				t.Log(string(actual))
			}
		})
	}
}

func openapiFilePaths(dir string) []string {
	filenames := make([]string, 0, 30)

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			base := filepath.Base(path)

			if info.IsDir() && (base == ".unused" || base == "scanner") {
				return filepath.SkipDir
			}

			if !info.IsDir() && filepath.Ext(path) == ".openapi" {
				filenames = append(filenames, path)
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	return filenames
}

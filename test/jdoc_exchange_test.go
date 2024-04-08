package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jsightapi/jsight-api-core/kit"
)

func TestJDocExchange(t *testing.T) {
	jsonFilesPaths := jsonFilePaths(GetTestDataDir())
	for _, jsonPath := range jsonFilesPaths {
		runJDocExchangeTests(t, jsonPath)
	}
}

func TestJDocExchangeSpecific(t *testing.T) {
	// path := filepath.Join(GetTestDataDir(), "/jsight_0.3/req.japi.http_method.path/path_04_undefined.json")
	// runJDocExchangeTests(t, path)
}

func runJDocExchangeTests(t *testing.T, jsonFilePath string) {
	t.Run(cutRepositoryPath(jsonFilePath), func(t *testing.T) {
		json, err := os.ReadFile(jsonFilePath)
		require.NoError(t, err)

		japiPath, err := japiFilePath(jsonFilePath)
		require.NoError(t, err)

		j, je := kit.NewJapi(japiPath)
		if je != nil {
			logJAPIError(t, je)
			t.FailNow()
		}

		actual, err := j.ToJsonIndent()
		require.NoError(t, err)

		expected := string(json)

		ok := assert.JSONEq(t, expected, string(actual))

		if !ok {
			t.Log("Actual JSON:")
			t.Log(string(actual))
		}
	})
}

func jsonFilePaths(dir string) []string {
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

			if !info.IsDir() && filepath.Ext(path) == ".json" {
				filenames = append(filenames, path)
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	return filenames
}

func japiFilePath(resFilePath string) (string, error) {
	ext := filepath.Ext(resFilePath)
	japiFilePath := resFilePath[:len(resFilePath)-len(ext)] + ".jst"
	return japiFilePath, nil
}

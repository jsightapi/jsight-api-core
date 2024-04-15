package test

import (
	"encoding/json"
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
		runOpenapiTest(t, openapiPath)
	}
}

func TestOpenAPISpecific(t *testing.T) {
	// path := filepath.Join(GetTestDataDir(), "/jsight_0.3/ser-to-openapi/dev/params/annotations_serv-378.openapi")
	// runOpenapiTest(t, path)
}

func runOpenapiTest(t *testing.T, openapiPath string) {
	t.Run(cutRepositoryPath(openapiPath), func(t *testing.T) {
		openapiJsonBytes, err := os.ReadFile(openapiPath)
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

		ok := assert.JSONEq(t, string(openapiJsonBytes), string(actual))

		dir := filepath.Dir(openapiPath)
		openapiFileName := filepath.Base(openapiPath)
		actualFilePath := filepath.Join(dir, openapiFileName+".ACTUAL.json")
		expectedFilePath := filepath.Join(dir, openapiFileName+".EXPECTED.json")

		if ok {
			if os.Remove(actualFilePath) == nil {
				t.Log("Bug fixed, ACTUAL OpenApi JSON removed: " + actualFilePath)
			}
			if os.Remove(expectedFilePath) == nil {
				t.Log("Bug fixed, ACTUAL OpenApi JSON removed: " + expectedFilePath)
			}
		} else {
			/*oa, err := openapi.NewOpenAPI(nil)
			t.Log("OpenApi: ")
			t.Log(oa)
			t.Log(err)
			json.Unmarshal([]byte(openapiJsonBytes), oa)
			expected, _ := json.MarshalIndent(oa, "", "  ")
			*/

			expectedAligned, err := alignJSON(openapiJsonBytes)
			if err != nil {
				t.Logf("Expected openapi file alignment failed: %s", err)
			}

			actualAligned, err := alignJSON(actual)
			if err != nil {
				t.Logf("Actual openapi file alignment failed: %s", err)
			}

			os.WriteFile(actualFilePath, actualAligned, 0644)
			os.WriteFile(expectedFilePath, expectedAligned, 0644)
			t.Log("ACTUAL OpenApi JSON written to: " + actualFilePath)
			t.Log("EXPECTED OpenApi JSON written to: " + expectedFilePath)
		}
	})
}

func alignJSON(j []byte) ([]byte, error) {
	var JSONAsInterface interface{}

	if err := json.Unmarshal(j, &JSONAsInterface); err != nil {
		return []byte{}, err
	}

	alignedJson, err := json.MarshalIndent(&JSONAsInterface, "", "  ")
	if err != nil {
		return []byte{}, err
	}
	return alignedJson, nil
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

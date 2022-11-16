package test

import (
	"path/filepath"

	"github.com/jsightapi/jsight-schema-core/test"
)

func GetTestDataDir() string {
	return filepath.Join(test.GetProjectRoot(), "testdata")
}

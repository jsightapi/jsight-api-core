package test

import (
	"path/filepath"
	"testing"

	"github.com/jsightapi/jsight-api-core/kit"
)

func BenchmarkJAPI(b *testing.B) {
	filename := filepath.Join(GetTestDataDir(), "jsight_0.3", "others", "full.jst")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		j, je := kit.NewJapi(filename)
		if je != nil {
			b.Error(je)
		}

		_, err := j.ToJson()
		if err != nil {
			b.Error(je)
		}
	}
}

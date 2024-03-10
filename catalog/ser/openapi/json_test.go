package openapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type Custom struct {
	SomeString string
}

func testMarshal(t *testing.T) {
	c := Custom{
		"some",
	}
	b, err := json.Marshal(c)
	
	require.NoError(t, err)

	t.Log(string(b))
}

func (o Custom) MarshalJSON() (b []byte, err error) {
	type Alias Custom
	return json.Marshal(&struct {
		Alias
	}{
		Alias:    (Alias)(o),
	})
}

package openapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type wrapper struct {
	Mapper Mapper `json:"mapper"`
}

type ztruct struct {
	SomeInt int `json:"int,omitempty"`
	SomeMap map[string]ztruct `json:"someMap,omitempty"`
	SomeMapper *Mapper `json:"someMapper,omitempty"`

}

type Mapper map[string]*ztruct

func TestOmitEmpty(t *testing.T) {
	map_ := make(map[string]ztruct, 0)
	mapper := make(Mapper, 0)
	
	r := ztruct{
		1,
		map_,
		&mapper,
	}

	b, err := json.Marshal(r)
	
	require.NoError(t, err)

	t.Log(string(b))
}

package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatenateDescription(t *testing.T) {
	type test struct {
		l   string
		r   string
		exp string
	}

	tests := []test{
		{"l", "r", "l: r"},
		{"", "r", "r"},
		{"l", "", "l"},
		{"", "", ""},
	}

	for _, test := range tests {
		assert.Equal(t, test.exp, concatenateDescription(test.l, test.r))
	}
}

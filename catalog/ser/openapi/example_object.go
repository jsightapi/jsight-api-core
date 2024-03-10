package openapi

import "encoding/json"

// Not used until decided to include full examples
type ExampleObject struct {
	Summary     string          `json:"summary,omitempty"`
	Description string          `json:"description,omitempty"`
	Value       json.RawMessage `json:"value"`
	// ExternalValue string // irrelevant for JS 0.3
}

func newExampleObject(b []byte) *ExampleObject {
	return &ExampleObject{
		Value: toExampleValue(b),
	}
}

func toExampleValue(b []byte) json.RawMessage { // TODO: discuss
	return json.RawMessage(b)
}

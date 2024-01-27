package openapi

type ExampleObject struct {
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Value       interface{} `json:"value"`
	// ExternalValue string // irrelevant for JS 0.3
}

func NewExampleObject(example []byte) *ExampleObject {
	return &ExampleObject{
		Value: example,
	}
}

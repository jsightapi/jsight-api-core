package openapi

type Response struct {
	Description string `json:"description"`
	// Headers *Headers
	Content *Content `json:"content,omitempty"`
	// Links // irrelevant to JSight
}

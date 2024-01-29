package openapi

type Response struct {
	Description string `json:"description"`
	// Headers *Headers
	Content *Content `json:"content,omitempty"`
	// Links // irrelevant to JSight
}

func defaultResponse() *Response {
	return &Response{
		Description: "",
		Content:     defaultContent(),
	}
}

func defaultContent() *Content {
	c := make(Content, 1)

	c["*/*"] = &MediaTypeObject{}

	return &c
}

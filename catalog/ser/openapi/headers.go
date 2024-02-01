package openapi

import "github.com/jsightapi/jsight-api-core/catalog"

// Only for Response objects. For requests refer to "parameters"
type Headers map[string]*HeaderObject

// TODO: may not be a schema, may be a literal | array
func NewHeaders(h *catalog.HTTPResponseHeaders) *Headers {
	if h == nil {
		return nil
	}

	r := make(Headers, 0)

	// TODO: children may be nil of schema is literal
	for _, ch := range h.Schema.ASTNode.Children {
		pi := GetParamInfo(ch)
		r[ch.Key] = NewHeaderObject(pi.Required(), pi.AllowEmptyValue(), pi.Description(), pi.Schema())
	}

	return &r
}

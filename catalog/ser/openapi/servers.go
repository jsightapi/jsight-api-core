package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Server struct {
	Url         string                    `json:"url"`
	Description string                    `json:"description"` // def "", md support
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

func defaultServers() *[]Server {
	return &[]Server{
		{
			Url: "/",
		},
	}
}

func NewServers(ss *catalog.Servers) *[]Server {
	if ss.Len() == 0 {
		return defaultServers()
	}

	r := make([]Server, 0, ss.Len())
  ss.Each(func(k string, v *catalog.Server) error {
    r = append(r, NewServer(v))
    return nil // TODO: what kind of error?
  })

	return &r
}

func NewServer(s *catalog.Server) Server {
	return Server{
		Url:         s.BaseUrl,
		Description: s.Annotation,
	}
}

// TODO:
type ServerVariable struct {
	Enum        []string `json:"enum"`        // An enumeration of string values to be used if the substitution options are from a limited set. The array SHOULD NOT be empty.
	Default     string   `json:"default`      // REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object's treatment of default values, because in those cases parameter values are optional. If the enum is defined, the value SHOULD exist in the enum's values.
	Description string   `json:"description"` // An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.
}

func NewServerVariable(default_ string) {

}

/* NOTES:
OAS has no concept of names for servers.
*/

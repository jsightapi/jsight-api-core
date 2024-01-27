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
	return nil
}

func NewServers(ss *catalog.Servers) *[]Server {
	if ss.Len() == 0 {
		return defaultServers()
	}

	r := make([]Server, 0, ss.Len())
	_ = ss.Each(func(k string, v *catalog.Server) error {
		r = append(r, NewServer(v))
		return nil
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
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

func NewServerVariable(default_ string) {

}

/* NOTES:
OAS has no concept of names for servers.
*/

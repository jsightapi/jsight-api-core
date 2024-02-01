package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

// TODO: 1. Server name to desciption? 2. empty baseUrl
type Server struct {
	Url         string `json:"url"`
	Description string `json:"description"` // def "", md support
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
		r = append(r, NewServer(k, v))
		return nil
	})

	return &r
}

func NewServer(name string, s *catalog.Server) Server {
	return Server{
		Url:         s.BaseUrl,
		Description: ServerDescription(name, s.Annotation),
	}
}

func ServerDescription(name string, annotation string) string {
	return name[1:] + ": " + annotation
}

/* NOTES:
OAS has no concept of names for servers.
*/

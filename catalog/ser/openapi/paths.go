package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Path struct {

}

type Paths map[string]*Path    // TODO: temp for stub

func defaultPaths() *Paths {
	return &Paths{
    "/": &Path{},
  }
}

func NewPaths(ss *catalog.Catalog) *Paths {
  return defaultPaths() // TODO: just a stub
}

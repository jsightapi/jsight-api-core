package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Paths map[string]*PathItem // TODO: temp for stub

func defaultPaths() *Paths {
	return &Paths{
		"/": &PathItem{},
	}
}

func NewPaths(c *catalog.Catalog) *Paths {
	if c.Interactions.Len() == 0 {
		return defaultPaths()
	}

	p := make(Paths, c.Interactions.Len())
	p.readInteractions(c)

	return &p
}

func (p Paths) readInteractions(c *catalog.Catalog) {
	_ = c.Interactions.Each(func(k catalog.InteractionID, v catalog.Interaction) error {
		if k.Protocol() == catalog.HTTP {
			i := v.(*catalog.HTTPInteraction)
			p.addInteraction(i)
		}
		return nil
	})
}

func (p Paths) addInteraction(i *catalog.HTTPInteraction) {
	path := i.Path().String()
	if _, exists := p[path]; exists {
		p[path].assignOperation(i.HttpMethod, NewOperation(i))
	} else {
		p[path] = NewPathItem(i)
	}
}

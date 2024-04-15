package openapi

import (
	"github.com/jsightapi/jsight-api-core/catalog"
)

type Paths map[string]*PathItem

func defaultPaths() Paths {
	return Paths{}
}

func newPaths(c *catalog.Catalog) Paths {
	if c.Interactions.Len() == 0 {
		return defaultPaths()
	}

	p := make(Paths, c.Interactions.Len())
	fillPaths(p, c)
	return p
}

func fillPaths(p Paths, c *catalog.Catalog) {
	_ = c.Interactions.Each(func(k catalog.InteractionID, v catalog.Interaction) error {
		if k.Protocol() == catalog.HTTP {
			i := v.(*catalog.HTTPInteraction)
			addOperation(p, i, c)
		}
		return nil
	})
}

func addOperation(p Paths, i *catalog.HTTPInteraction, c *catalog.Catalog) {
	path := i.Path().String()
	tags := gatherTags(c, i.Tags)
	tagTitles := getTagTitles(tags)
	if _, exists := p[path]; exists {
		p[path].assignOperation(i.HttpMethod, newOperation(i, tagTitles))
	} else {
		p[path] = newPathItem(i, tagTitles)
	}
}

func gatherTags(c *catalog.Catalog, names []catalog.TagName) []*catalog.Tag {
	tags := make([]*catalog.Tag, len(names))
	for _, n := range names {
		if t, ok := c.Tags.Get(n); ok {
			tags = append(tags, t)
		}
	}
	return tags
}

func getTagTitles(tags []*catalog.Tag) []string {
	titles := make([]string, len(tags))
	for _, t := range tags {
		titles = append(titles, t.Title)
	}
	return titles
}

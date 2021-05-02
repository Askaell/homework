package models

type Category struct {
	Id       int
	Slug     string
	Name     string
	ParentId int
	Parent   *Category
	Childs   []*Category
}

func NewCategory(id int, slug string, name string, parent *Category) *Category {
	var c *Category
	c.Id = id
	c.Slug = slug
	c.Name = name
	c.Childs = make(Categories, 0)

	if parent == nil {
		c.ParentId = 0
		c.Parent = nil
		return c
	}

	c.ParentId = parent.Id
	c.Parent = parent
	c.appendChildToParent()

	return c
}

func (c *Category) appendChildToParent() {
	if c.Parent != nil {
		c.Parent.Childs = append(c.Parent.Childs, c)
	}
}

func (c *Category) GetRootParent() *Category {
	if c.Id == c.ParentId {
		return c
	}
	return c.Parent.GetRootParent()
}

func (c *Category) GetBreadCrumbs() Categories {
	if c.Id == c.ParentId {
		return []*Category{c}
	}
	return append(c.Parent.GetBreadCrumbs(), c)
}

type Categories []*Category

func (cs *Categories) Names() (names []string) {
	for _, c := range *cs {
		if c != nil {
			names = append(names, c.Name)
		}
	}
	return names
}

func (cs *Categories) Slugs() (slugs []string) {
	for _, c := range *cs {
		if c != nil {
			slugs = append(slugs, c.Slug)
		}
	}
	return slugs
}

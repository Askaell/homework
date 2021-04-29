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
	c.Childs = make([]*Category, 0)

	if parent == nil {
		c.ParentId = id
		c.Parent = c
		return c
	}

	c.ParentId = parent.Id
	c.Parent = parent

	return c
}

func (c *Category) GetRootParent() *Category {
	if c.Id == c.ParentId {
		return c
	}
	return c.Parent.GetRootParent()
}

func (c *Category) GetBreadCrumbs() []*Category {
	if c.Id == c.ParentId {
		return []*Category{c}
	}
	return append(c.Parent.GetBreadCrumbs(), c)
}

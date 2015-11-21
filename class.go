package ospokemon

type Class interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Description() string
	SetDescription(description string)
}

type BasicClass struct {
	ID          int
	NAME        string
	DESCRIPTION string
}

func (c *BasicClass) Id() int {
	return c.ID
}

func (c *BasicClass) SetId(id int) {
	c.ID = id
}

func (c *BasicClass) Name() string {
	return c.NAME
}

func (c *BasicClass) SetName(name string) {
	c.NAME = name
}

func (c *BasicClass) Description() string {
	return c.DESCRIPTION
}

func (c *BasicClass) SetDescription(description string) {
	c.DESCRIPTION = description
}

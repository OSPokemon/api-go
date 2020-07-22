package pokedex

type Cache struct {
	cache map[int]*Region
}

func (c *Cache) Get(id int) *Region {
	return c.cache[id]
}

func (c *Cache) Add(r *Region) error {
	if c.cache[r.ID()] != nil {
		return ErrDuplicateRegionID
	}
	c.cache[r.ID()] = r
	return nil
}

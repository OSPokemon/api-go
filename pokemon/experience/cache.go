package experience

// Cache stores experience curves by id
type Cache struct {
	cache map[int]Curve
}

// New creates a new Cache
func New() *Cache {
	return &Cache{
		cache: make(map[int]Curve),
	}
}

func (c *Cache) Get(id int) Curve { return c.cache[id] }

func (c *Cache) Set(id int, f Curve) { c.cache[id] = f }

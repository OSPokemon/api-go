package pokedex

// Region is a regional pokedex
type Region struct {
	id    int
	cache map[int][]Track
}

// Track is tracking data
type Track struct {
	Zone      int
	Weather   int
	Frequency float64
	Hidden    bool
}

// NewRegion makes a new Region
func NewRegion(id int) *Region {
	return &Region{
		id:    id,
		cache: make(map[int][]Track),
	}
}

// ID returns the ID
func (r *Region) ID() int { return r.id }

// Get returns tracking for the regional id
func (r *Region) Get(id int) []Track { return r.cache[id] }

// Add appends tracking data for the regional id
func (r *Region) Add(id int, track Track) { r.cache[id] = append(r.cache[id], track) }

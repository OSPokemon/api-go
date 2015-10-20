package ospokemon

type StatusCondition interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Volatile() bool
	SetVolatile(volatile bool)
}

type BasicStatusCondition struct {
	id       int
	name     string
	volatile bool
}

func (sc *BasicStatusCondition) Id() int {
	return sc.id
}

func (sc *BasicStatusCondition) SetId(id int) {
	sc.id = id
}

func (sc *BasicStatusCondition) Name() string {
	return sc.name
}

func (sc *BasicStatusCondition) SetName(name string) {
	sc.name = name
}

func (sc *BasicStatusCondition) Volatile() bool {
	return sc.volatile
}

func (sc *BasicStatusCondition) SetVolatile(volatile bool) {
	sc.volatile = volatile
}

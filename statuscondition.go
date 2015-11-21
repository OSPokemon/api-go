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
	ID       int
	NAME     string
	VOLATILE bool
}

func (sc *BasicStatusCondition) Id() int {
	return sc.ID
}

func (sc *BasicStatusCondition) SetId(id int) {
	sc.ID = id
}

func (sc *BasicStatusCondition) Name() string {
	return sc.NAME
}

func (sc *BasicStatusCondition) SetName(name string) {
	sc.NAME = name
}

func (sc *BasicStatusCondition) Volatile() bool {
	return sc.VOLATILE
}

func (sc *BasicStatusCondition) SetVolatile(volatile bool) {
	sc.VOLATILE = volatile
}

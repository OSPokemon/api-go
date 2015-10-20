package ospokemon

type Ability interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Description() string
	SetDescription(description string)
}

type BasicAbility struct {
	id          int
	name        string
	description string
}

func (a *BasicAbility) Id() int {
	return a.id
}

func (a *BasicAbility) SetId(id int) {
	a.id = id
}

func (a *BasicAbility) Name() string {
	return a.name
}

func (a *BasicAbility) SetName(name string) {
	a.name = name
}

func (a *BasicAbility) Description() string {
	return a.description
}

func (a *BasicAbility) SetDescription(description string) {
	a.description = description
}

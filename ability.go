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
	ID          int
	NAME        string
	DESCRIPTION string
}

func (a *BasicAbility) Id() int {
	return a.ID
}

func (a *BasicAbility) SetId(id int) {
	a.ID = id
}

func (a *BasicAbility) Name() string {
	return a.NAME
}

func (a *BasicAbility) SetName(name string) {
	a.NAME = name
}

func (a *BasicAbility) Description() string {
	return a.DESCRIPTION
}

func (a *BasicAbility) SetDescription(description string) {
	a.DESCRIPTION = description
}

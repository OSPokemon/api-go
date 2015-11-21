package ospokemon

type Trainer interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Class() int
	SetClass(class int)

	Money() int
	SetMoney(money int)

	Items() map[int]int
	SetItems(items map[int]int)

	Pokemon() []int
	SetPokemon(pokemon []int)
}

type BasicTrainer struct {
	ID      int
	NAME    string
	CLASS   int
	MONEY   int
	ITEMS   map[int]int
	POKEMON []int
}

func (t *BasicTrainer) Id() int {
	return t.ID
}

func (t *BasicTrainer) SetId(ID int) {
	t.ID = ID
}

func (t *BasicTrainer) Name() string {
	return t.NAME
}

func (t *BasicTrainer) SetName(name string) {
	t.NAME = name
}

func (t *BasicTrainer) Class() int {
	return t.CLASS
}

func (t *BasicTrainer) SetClass(class int) {
	t.CLASS = class
}

func (t *BasicTrainer) Money() int {
	return t.MONEY
}

func (t *BasicTrainer) SetMoney(money int) {
	t.MONEY = money
}

func (t *BasicTrainer) Items() map[int]int {
	return t.ITEMS
}

func (t *BasicTrainer) SetItems(items map[int]int) {
	t.ITEMS = items
}

func (t *BasicTrainer) Pokemon() []int {
	return t.POKEMON
}

func (t *BasicTrainer) SetPokemon(pokemon []int) {
	t.POKEMON = pokemon
}

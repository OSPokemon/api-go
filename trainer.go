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
	id      int
	name    string
	class   int
	money   int
	items   map[int]int
	pokemon []int
}

func (t *BasicTrainer) Id() int {
	return t.id
}

func (t *BasicTrainer) SetId(id int) {
	t.id = id
}

func (t *BasicTrainer) Name() string {
	return t.name
}

func (t *BasicTrainer) SetName(name string) {
	t.name = name
}

func (t *BasicTrainer) Class() int {
	return t.class
}

func (t *BasicTrainer) SetClass(class int) {
	t.class = class
}

func (t *BasicTrainer) Money() int {
	return t.money
}

func (t *BasicTrainer) SetMoney(money int) {
	t.money = money
}

func (t *BasicTrainer) Items() map[int]int {
	return t.items
}

func (t *BasicTrainer) SetItems(items map[int]int) {
	t.items = items
}

func (t *BasicTrainer) Pokemon() []int {
	return t.pokemon
}

func (t *BasicTrainer) SetPokemon(pokemon []int) {
	t.pokemon = pokemon
}

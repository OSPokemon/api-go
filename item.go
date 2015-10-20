package ospokemon

type Item interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Description() string
	SetDescription(description string)

	Sellable() bool
	SetSellable(sellable bool)

	SalePrice() int
	SetSalePrice(salePrice int)

	UsableOutsideBattle() bool
	SetUsableOutsideBattle(usable bool)

	UsableDuringBattle() bool
	SetUsableDuringBattle(usable bool)

	Consumable() bool
	SetConsumable(consumable bool)

	Holdable() bool
	SetHoldable(holdable bool)
}

type BasicItem struct {
	id                  int
	name                string
	description         string
	sellable            bool
	salePrice           int
	usableOutsideBattle bool
	usableDuringBattle  bool
	consumable          bool
	holdable            bool
}

func (i *BasicItem) Id() int {
	return i.id
}

func (i *BasicItem) SetId(id int) {
	i.id = id
}

func (i *BasicItem) Name() string {
	return i.name
}

func (i *BasicItem) SetName(name string) {
	i.name = name
}

func (i *BasicItem) Description() string {
	return i.description
}

func (i *BasicItem) SetDescription(description string) {
	i.description = description
}

func (i *BasicItem) Sellable() bool {
	return i.sellable
}

func (i *BasicItem) SetSellable(sellable bool) {
	i.sellable = sellable
}

func (i *BasicItem) SalePrice() int {
	return i.salePrice
}

func (i *BasicItem) SetSalePrice(salePrice int) {
	i.salePrice = salePrice
}

func (i *BasicItem) UsableOutsideBattle() bool {
	return i.usableOutsideBattle
}

func (i *BasicItem) SetUsableOutsideBattle(usable bool) {
	i.usableOutsideBattle = usable
}

func (i *BasicItem) UsableDuringBattle() bool {
	return i.usableDuringBattle
}

func (i *BasicItem) SetUsableDuringBattle(usable bool) {
	i.usableDuringBattle = usable
}

func (i *BasicItem) Consumable() bool {
	return i.consumable
}

func (i *BasicItem) SetConsumable(consumable bool) {
	i.consumable = consumable
}

func (i *BasicItem) Holdable() bool {
	return i.holdable
}

func (i *BasicItem) SetHoldable(holdable bool) {
	i.holdable = holdable
}

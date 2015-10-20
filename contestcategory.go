package ospokemon

type ContestCategory interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Flavor() string
	SetFlavor(flavor string)

	Color() string
	SetColor(color string)

	Stat() string
	SetStat(stat string)

	Reactions() map[int]string
	SetReactions(reactions map[int]string)
}

type BasicContestCategory struct {
	id        int
	name      string
	flavor    string
	color     string
	stat      string
	reactions map[int]string
}

func (cc *BasicContestCategory) Id() int {
	return cc.id
}

func (cc *BasicContestCategory) SetId(id int) {
	cc.id = id
}

func (cc *BasicContestCategory) Name() string {
	return cc.name
}

func (cc *BasicContestCategory) SetName(name string) {
	cc.name = name
}

func (cc *BasicContestCategory) Flavor() string {
	return cc.flavor
}

func (cc *BasicContestCategory) SetFlavor(flavor string) {
	cc.flavor = flavor
}

func (cc *BasicContestCategory) Color() string {
	return cc.color
}

func (cc *BasicContestCategory) SetColor(color string) {
	cc.color = color
}

func (cc *BasicContestCategory) Stat() string {
	return cc.stat
}

func (cc *BasicContestCategory) SetStat(stat string) {
	cc.stat = stat
}

func (cc *BasicContestCategory) Reactions() map[int]string {
	return cc.reactions
}

func (cc *BasicContestCategory) SetReactions(reactions map[int]string) {
	cc.reactions = reactions
}

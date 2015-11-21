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
	ID        int
	NAME      string
	FLAVOR    string
	COLOR     string
	STAT      string
	REACTIONS map[int]string
}

func (cc *BasicContestCategory) Id() int {
	return cc.ID
}

func (cc *BasicContestCategory) SetId(id int) {
	cc.ID = id
}

func (cc *BasicContestCategory) Name() string {
	return cc.NAME
}

func (cc *BasicContestCategory) SetName(name string) {
	cc.NAME = name
}

func (cc *BasicContestCategory) Flavor() string {
	return cc.FLAVOR
}

func (cc *BasicContestCategory) SetFlavor(flavor string) {
	cc.FLAVOR = flavor
}

func (cc *BasicContestCategory) Color() string {
	return cc.COLOR
}

func (cc *BasicContestCategory) SetColor(color string) {
	cc.COLOR = color
}

func (cc *BasicContestCategory) Stat() string {
	return cc.STAT
}

func (cc *BasicContestCategory) SetStat(stat string) {
	cc.STAT = stat
}

func (cc *BasicContestCategory) Reactions() map[int]string {
	return cc.REACTIONS
}

func (cc *BasicContestCategory) SetReactions(reactions map[int]string) {
	cc.REACTIONS = reactions
}

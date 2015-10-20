package ospokemon

type Move interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Type() int
	SetType(t int)

	Category() string
	SetCategory(category string)

	Description() string
	SetDescription(description string)

	PP() int
	SetPP(pp int)

	Power() int
	SetPower(power int)

	Accuracy() float64
	SetAccuracy(accuracy float64)

	Target() int
	SetTarget(target int)

	ContestCategory() int
	SetContestCategory(contestCategory int)

	Appeal() int
	SetAppeal(appeal int)

	Jam() int
	SetJam(jam int)

	ContestDescription() string
	SetContestDescription(contestDescription string)

	Priority() int
	SetPriority(priority int)
}

type BasicMove struct {
	id                 int
	name               string
	t                  int
	category           string
	description        string
	pp                 int
	power              int
	accuracy           float64
	target             int
	contestCategory    int
	appeal             int
	jam                int
	contestDescription string
	priority           int
}

func (m *BasicMove) Id() int {
	return m.id
}

func (m *BasicMove) SetId(id int) {
	m.id = id
}

func (m *BasicMove) Name() string {
	return m.name
}

func (m *BasicMove) SetName(name string) {
	m.name = name
}

func (m *BasicMove) Type() int {
	return m.t
}

func (m *BasicMove) SetType(t int) {
	m.t = t
}

func (m *BasicMove) Category() string {
	return m.category
}

func (m *BasicMove) SetCategory(category string) {
	m.category = category
}

func (m *BasicMove) Description() string {
	return m.description
}

func (m *BasicMove) SetDescription(description string) {
	m.description = description
}

func (m *BasicMove) PP() int {
	return m.pp
}

func (m *BasicMove) SetPP(pp int) {
	m.pp = pp
}

func (m *BasicMove) Power() int {
	return m.power
}

func (m *BasicMove) SetPower(power int) {
	m.power = power
}

func (m *BasicMove) Accuracy() float64 {
	return m.accuracy
}

func (m *BasicMove) SetAccuracy(accuracy float64) {
	m.accuracy = accuracy
}

func (m *BasicMove) Target() int {
	return m.target
}

func (m *BasicMove) SetTarget(target int) {
	m.target = target
}

func (m *BasicMove) ContestCategory() int {
	return m.contestCategory
}

func (m *BasicMove) SetContestCategory(contestCategory int) {
	m.contestCategory = contestCategory
}

func (m *BasicMove) Appeal() int {
	return m.appeal
}

func (m *BasicMove) SetAppeal(appeal int) {
	m.appeal = appeal
}

func (m *BasicMove) Jam() int {
	return m.jam
}

func (m *BasicMove) SetJam(jam int) {
	m.jam = jam
}

func (m *BasicMove) ContestDescription() string {
	return m.contestDescription
}

func (m *BasicMove) SetContestDescription(contestDescription string) {
	m.contestDescription = contestDescription
}

func (m *BasicMove) Priority() int {
	return m.priority
}

func (m *BasicMove) SetPriority(priority int) {
	m.priority = priority
}

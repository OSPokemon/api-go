package ospokemon

type Pokemon interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Species() int
	SetSpecies(species int)

	Level() int
	SetLevel(level int)

	Experience() int
	SetExperience(experience int)

	Ability() int
	SetAbility(ability int)

	Friendship() int
	SetFriendship(friendship int)

	Gender() int
	SetGender(gender int)

	Nature() int
	SetNature(nature int)

	Height() float64
	SetHeight(height float64)

	Weight() float64
	SetWeight(weight float64)

	OriginalTrainer() int
	SetOriginalTrainer(originalTrainer int)

	Shiny() bool
	SetShiny(shiny bool)

	Item() int
	SetItem(item int)

	Stats() map[string]Stat
	SetStats(stats map[string]Stat)

	Skills() []Skill
	SetSkills(skills []Skill)

	StatusConditions() []int
	SetStatusConditions(statusConditions []int)

	Markings() []string
	SetMarkings(markings []string)

	Ribbons() []string
	SetRibbons(ribbons []string)
}

type BasicPokemon struct {
	ID               int
	NAME             string
	SPECIES          int
	LEVEL            int
	EXPERIENCE       int
	ABILITY          int
	FRIENDSHIP       int
	GENDER           int
	NATURE           int
	HEIGHT           float64
	WEIGHT           float64
	ORIGINALTRAINER  int
	SHINY            bool
	ITEM             int
	STATS            map[string]Stat
	SKILLS           []Skill
	STATUSCONDITIONS []int
	MARKINGS         []string
	RIBBONS          []string
}

func (p *BasicPokemon) Id() int {
	return p.ID
}

func (p *BasicPokemon) SetId(id int) {
	p.ID = id
}

func (p *BasicPokemon) Name() string {
	return p.NAME
}

func (p *BasicPokemon) SetName(name string) {
	p.NAME = name
}

func (p *BasicPokemon) Species() int {
	return p.SPECIES
}

func (p *BasicPokemon) SetSpecies(species int) {
	p.SPECIES = species
}

func (p *BasicPokemon) Level() int {
	return p.LEVEL
}

func (p *BasicPokemon) SetLevel(level int) {
	p.LEVEL = level
}

func (p *BasicPokemon) Experience() int {
	return p.EXPERIENCE
}

func (p *BasicPokemon) SetExperience(experience int) {
	p.EXPERIENCE = experience
}

func (p *BasicPokemon) Ability() int {
	return p.ABILITY
}

func (p *BasicPokemon) SetAbility(ability int) {
	p.ABILITY = ability
}

func (p *BasicPokemon) Friendship() int {
	return p.FRIENDSHIP
}

func (p *BasicPokemon) SetFriendship(friendship int) {
	p.FRIENDSHIP = friendship
}

func (p *BasicPokemon) Gender() int {
	return p.GENDER
}

func (p *BasicPokemon) SetGender(gender int) {
	p.GENDER = gender
}

func (p *BasicPokemon) Nature() int {
	return p.NATURE
}

func (p *BasicPokemon) SetNature(nature int) {
	p.NATURE = nature
}

func (p *BasicPokemon) Height() float64 {
	return p.HEIGHT
}

func (p *BasicPokemon) SetHeight(height float64) {
	p.HEIGHT = height
}

func (p *BasicPokemon) Weight() float64 {
	return p.WEIGHT
}

func (p *BasicPokemon) SetWeight(weight float64) {
	p.WEIGHT = weight
}

func (p *BasicPokemon) OriginalTrainer() int {
	return p.ORIGINALTRAINER
}

func (p *BasicPokemon) SetOriginalTrainer(originalTrainer int) {
	p.ORIGINALTRAINER = originalTrainer
}

func (p *BasicPokemon) Shiny() bool {
	return p.SHINY
}

func (p *BasicPokemon) SetShiny(shiny bool) {
	p.SHINY = shiny
}

func (p *BasicPokemon) Item() int {
	return p.ITEM
}

func (p *BasicPokemon) SetItem(item int) {
	p.ITEM = item
}

func (p *BasicPokemon) Stats() map[string]Stat {
	return p.STATS
}

func (p *BasicPokemon) SetStats(stats map[string]Stat) {
	p.STATS = stats
}

func (p *BasicPokemon) Skills() []Skill {
	return p.SKILLS
}

func (p *BasicPokemon) SetSkills(skills []Skill) {
	p.SKILLS = skills
}

func (p *BasicPokemon) StatusConditions() []int {
	return p.STATUSCONDITIONS
}

func (p *BasicPokemon) SetStatusConditions(statusConditions []int) {
	p.STATUSCONDITIONS = statusConditions
}

func (p *BasicPokemon) Markings() []string {
	return p.MARKINGS
}

func (p *BasicPokemon) SetMarkings(markings []string) {
	p.MARKINGS = markings
}

func (p *BasicPokemon) Ribbons() []string {
	return p.RIBBONS
}

func (p *BasicPokemon) SetRibbons(ribbons []string) {
	p.RIBBONS = ribbons
}

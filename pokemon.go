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
	id               int
	name             string
	species          int
	level            int
	experience       int
	ability          int
	friendship       int
	gender           int
	nature           int
	height           float64
	weight           float64
	originalTrainer  int
	shiny            bool
	item             int
	stats            map[string]Stat
	skills           []Skill
	statusConditions []int
	markings         []string
	ribbons          []string
}

func (p *BasicPokemon) Id() int {
	return p.id
}

func (p *BasicPokemon) SetId(id int) {
	p.id = id
}

func (p *BasicPokemon) Name() string {
	return p.name
}

func (p *BasicPokemon) SetName(name string) {
	p.name = name
}

func (p *BasicPokemon) Species() int {
	return p.species
}

func (p *BasicPokemon) SetSpecies(species int) {
	p.species = species
}

func (p *BasicPokemon) Level() int {
	return p.level
}

func (p *BasicPokemon) SetLevel(level int) {
	p.level = level
}

func (p *BasicPokemon) Experience() int {
	return p.experience
}

func (p *BasicPokemon) SetExperience(experience int) {
	p.experience = experience
}

func (p *BasicPokemon) Ability() int {
	return p.ability
}

func (p *BasicPokemon) SetAbility(ability int) {
	p.ability = ability
}

func (p *BasicPokemon) Friendship() int {
	return p.friendship
}

func (p *BasicPokemon) SetFriendship(friendship int) {
	p.friendship = friendship
}

func (p *BasicPokemon) Gender() int {
	return p.gender
}

func (p *BasicPokemon) SetGender(gender int) {
	p.gender = gender
}

func (p *BasicPokemon) Nature() int {
	return p.nature
}

func (p *BasicPokemon) SetNature(nature int) {
	p.nature = nature
}

func (p *BasicPokemon) Height() float64 {
	return p.height
}

func (p *BasicPokemon) SetHeight(height float64) {
	p.height = height
}

func (p *BasicPokemon) Weight() float64 {
	return p.weight
}

func (p *BasicPokemon) SetWeight(weight float64) {
	p.weight = weight
}

func (p *BasicPokemon) OriginalTrainer() int {
	return p.originalTrainer
}

func (p *BasicPokemon) SetOriginalTrainer(originalTrainer int) {
	p.originalTrainer = originalTrainer
}

func (p *BasicPokemon) Shiny() bool {
	return p.shiny
}

func (p *BasicPokemon) SetShiny(shiny bool) {
	p.shiny = shiny
}

func (p *BasicPokemon) Item() int {
	return p.item
}

func (p *BasicPokemon) SetItem(item int) {
	p.item = item
}

func (p *BasicPokemon) Stats() map[string]Stat {
	return p.stats
}

func (p *BasicPokemon) SetStats(stats map[string]Stat) {
	p.stats = stats
}

func (p *BasicPokemon) Skills() []Skill {
	return p.skills
}

func (p *BasicPokemon) SetSkills(skills []Skill) {
	p.skills = skills
}

func (p *BasicPokemon) StatusConditions() []int {
	return p.statusConditions
}

func (p *BasicPokemon) SetStatusConditions(statusConditions []int) {
	p.statusConditions = statusConditions
}

func (p *BasicPokemon) Markings() []string {
	return p.markings
}

func (p *BasicPokemon) SetMarkings(markings []string) {
	p.markings = markings
}

func (p *BasicPokemon) Ribbons() []string {
	return p.ribbons
}

func (p *BasicPokemon) SetRibbons(ribbons []string) {
	p.ribbons = ribbons
}

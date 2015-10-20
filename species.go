package ospokemon

type Species interface {
	Id() int
	SetId(id int)

	Name() string
	SetName(name string)

	Tag() string
	SetTag(tag string)

	Description() string
	SetDescription(desctiption string)

	Types() []int
	SetTypes(types []int)

	Abilities() []int
	SetAbilities(abilities []int)

	HiddenAbility() int
	SetHiddenAbility(ability int)

	Genderless() bool
	SetGenderless(genderless bool)

	GenderRatio() float64
	SetGenderRatio(genderRatio float64)

	CatchRate() int
	SetCatchRate(catchRate int)

	Breedable() bool
	SetBreedable(breedable bool)

	EggGroups() []string
	SetEggGroups(eggGroups []string)

	EggCycles() int
	SetEggCycles(eggCycles int)

	Height() float64
	SetHeight(height float64)

	Weight() float64
	SetWeight(weight float64)

	ExperienceYield() int
	SetExperienceYield(experienceYield int)

	ExperienceCurve() int
	SetExperienceCurve(experienceCurve int)

	EffortValues() map[string]int
	SetEffortValues(effortValues map[string]int)

	BodyStyle() string
	SetBodyStyle(bodyStyle string)

	Color() string
	SetColor(color string)

	Tameness() int
	SetTameness(tameness int)

	Stats() map[string]int
	SetStats(stats map[string]int)

	Evolutions() []int
	SetEvolutions(evolutions []int)

	MoveLists() map[int][]int
	SetMoveLists(moveLists map[int][]int)

	EggMoves() []int
	SetEggMoves(eggMoves []int)

	MachineMoves() []int
	SetMachineMoves(machineMoves []int)

	TutorMoves() []int
	SetTutorMoves(tutorMoves []int)
}

type BasicSpecies struct {
	id              int
	name            string
	tag             string
	description     string
	types           []int
	abilities       []int
	hiddenAbility   int
	genderless      bool
	genderRatio     float64
	catchRate       int
	breedable       bool
	eggGroups       []string
	eggCycles       int
	height          float64
	weight          float64
	experienceYield int
	experienceCurve int
	effortValues    map[string]int
	bodyStyle       string
	color           string
	tameness        int
	stats           map[string]int
	evolutions      []int
	moveLists       map[int][]int
	eggMoves        []int
	machineMoves    []int
	tutorMoves      []int
}

func (s *BasicSpecies) Id() int {
	return s.id
}

func (s *BasicSpecies) SetId(id int) {
	s.id = id
}

func (s *BasicSpecies) Name() string {
	return s.name
}

func (s *BasicSpecies) SetName(name string) {
	s.name = name
}

func (s *BasicSpecies) Tag() string {
	return s.tag
}

func (s *BasicSpecies) SetTag(tag string) {
	s.tag = tag
}

func (s *BasicSpecies) Description() string {
	return s.description
}

func (s *BasicSpecies) SetDescription(desctiption string) {
	s.description = desctiption
}

func (s *BasicSpecies) Types() []int {
	return s.types
}

func (s *BasicSpecies) SetTypes(types []int) {
	s.types = types
}

func (s *BasicSpecies) Abilities() []int {
	return s.abilities
}

func (s *BasicSpecies) SetAbilities(abilities []int) {
	s.abilities = abilities
}

func (s *BasicSpecies) HiddenAbility() int {
	return s.hiddenAbility
}

func (s *BasicSpecies) SetHiddenAbility(hiddenAbility int) {
	s.hiddenAbility = hiddenAbility
}

func (s *BasicSpecies) Genderless() bool {
	return s.genderless
}

func (s *BasicSpecies) SetGenderless(genderless bool) {
	s.genderless = genderless
}

func (s *BasicSpecies) GenderRatio() float64 {
	return s.genderRatio
}

func (s *BasicSpecies) SetGenderRatio(genderRatio float64) {
	s.genderRatio = genderRatio
}

func (s *BasicSpecies) CatchRate() int {
	return s.catchRate
}

func (s *BasicSpecies) SetCatchRate(catchRate int) {
	s.catchRate = catchRate
}

func (s *BasicSpecies) Breedable() bool {
	return s.breedable
}

func (s *BasicSpecies) SetBreedable(breedable bool) {
	s.breedable = breedable
}

func (s *BasicSpecies) EggGroups() []string {
	return s.eggGroups
}

func (s *BasicSpecies) SetEggGroups(eggGroups []string) {
	s.eggGroups = eggGroups
}

func (s *BasicSpecies) EggCycles() int {
	return s.eggCycles
}

func (s *BasicSpecies) SetEggCycles(eggCycles int) {
	s.eggCycles = eggCycles
}

func (s *BasicSpecies) Height() float64 {
	return s.height
}

func (s *BasicSpecies) SetHeight(height float64) {
	s.height = height
}

func (s *BasicSpecies) Weight() float64 {
	return s.weight
}

func (s *BasicSpecies) SetWeight(weight float64) {
	s.weight = weight
}

func (s *BasicSpecies) ExperienceYield() int {
	return s.experienceYield
}

func (s *BasicSpecies) SetExperienceYield(experienceYield int) {
	s.experienceYield = experienceYield
}

func (s *BasicSpecies) ExperienceCurve() int {
	return s.experienceCurve
}

func (s *BasicSpecies) SetExperienceCurve(experienceCurve int) {
	s.experienceCurve = experienceCurve
}

func (s *BasicSpecies) EffortValues() map[string]int {
	return s.effortValues
}

func (s *BasicSpecies) SetEffortValues(effortValues map[string]int) {
	s.effortValues = effortValues
}

func (s *BasicSpecies) BodyStyle() string {
	return s.bodyStyle
}

func (s *BasicSpecies) SetBodyStyle(bodyStyle string) {
	s.bodyStyle = bodyStyle
}

func (s *BasicSpecies) Color() string {
	return s.color
}

func (s *BasicSpecies) SetColor(color string) {
	s.color = color
}

func (s *BasicSpecies) Tameness() int {
	return s.tameness
}

func (s *BasicSpecies) SetTameness(tameness int) {
	s.tameness = tameness
}

func (s *BasicSpecies) Stats() map[string]int {
	return s.stats
}

func (s *BasicSpecies) SetStats(stats map[string]int) {
	s.stats = stats
}

func (s *BasicSpecies) Evolutions() []int {
	return s.evolutions
}

func (s *BasicSpecies) SetEvolutions(evolutions []int) {
	s.evolutions = evolutions
}

func (s *BasicSpecies) MoveLists() map[int][]int {
	return s.moveLists
}

func (s *BasicSpecies) SetMoveLists(moveLists map[int][]int) {
	s.moveLists = moveLists
}

func (s *BasicSpecies) EggMoves() []int {
	return s.eggMoves
}

func (s *BasicSpecies) SetEggMoves(eggMoves []int) {
	s.eggMoves = eggMoves
}

func (s *BasicSpecies) MachineMoves() []int {
	return s.machineMoves
}

func (s *BasicSpecies) SetMachineMoves(machineMoves []int) {
	s.machineMoves = machineMoves
}

func (s *BasicSpecies) TutorMoves() []int {
	return s.tutorMoves
}

func (s *BasicSpecies) SetTutorMoves(tutorMoves []int) {
	s.tutorMoves = tutorMoves
}

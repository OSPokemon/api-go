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
	ID              int
	NAME            string
	TAG             string
	DESCRIPTION     string
	TYPES           []int
	ABILITIES       []int
	HIDDENABILITY   int
	GENDERLESS      bool
	GENDERRATIO     float64
	CATCHRATE       int
	BREEDABLE       bool
	EGGGROUPS       []string
	EGGCYCLES       int
	HEIGHT          float64
	WEIGHT          float64
	EXPERIENCEYIELD int
	EXPERIENCECURVE int
	EFFORTVALUES    map[string]int
	BODYSTYLE       string
	COLOR           string
	TAMENESS        int
	STATS           map[string]int
	EVOLUTIONS      []int
	MOVELISTS       map[int][]int
	EGGMOVES        []int
	MACHINEMOVES    []int
	TUTORMOVES      []int
}

var AllSpecies = make(map[int]Species)

var LoadSpecies func(speciesId int) (Species, error)
var CreateSpecies func(name string) (Species, error)
var SaveSpecies func(species Species) error

func (s *BasicSpecies) Id() int {
	return s.ID
}

func (s *BasicSpecies) SetId(id int) {
	s.ID = id
}

func (s *BasicSpecies) Name() string {
	return s.NAME
}

func (s *BasicSpecies) SetName(name string) {
	s.NAME = name
}

func (s *BasicSpecies) Tag() string {
	return s.TAG
}

func (s *BasicSpecies) SetTag(tag string) {
	s.TAG = tag
}

func (s *BasicSpecies) Description() string {
	return s.DESCRIPTION
}

func (s *BasicSpecies) SetDescription(desctiption string) {
	s.DESCRIPTION = desctiption
}

func (s *BasicSpecies) Types() []int {
	return s.TYPES
}

func (s *BasicSpecies) SetTypes(types []int) {
	s.TYPES = types
}

func (s *BasicSpecies) Abilities() []int {
	return s.ABILITIES
}

func (s *BasicSpecies) SetAbilities(abilities []int) {
	s.ABILITIES = abilities
}

func (s *BasicSpecies) HiddenAbility() int {
	return s.HIDDENABILITY
}

func (s *BasicSpecies) SetHiddenAbility(hiddenAbility int) {
	s.HIDDENABILITY = hiddenAbility
}

func (s *BasicSpecies) Genderless() bool {
	return s.GENDERLESS
}

func (s *BasicSpecies) SetGenderless(genderless bool) {
	s.GENDERLESS = genderless
}

func (s *BasicSpecies) GenderRatio() float64 {
	return s.GENDERRATIO
}

func (s *BasicSpecies) SetGenderRatio(genderRatio float64) {
	s.GENDERRATIO = genderRatio
}

func (s *BasicSpecies) CatchRate() int {
	return s.CATCHRATE
}

func (s *BasicSpecies) SetCatchRate(catchRate int) {
	s.CATCHRATE = catchRate
}

func (s *BasicSpecies) Breedable() bool {
	return s.BREEDABLE
}

func (s *BasicSpecies) SetBreedable(breedable bool) {
	s.BREEDABLE = breedable
}

func (s *BasicSpecies) EggGroups() []string {
	return s.EGGGROUPS
}

func (s *BasicSpecies) SetEggGroups(eggGroups []string) {
	s.EGGGROUPS = eggGroups
}

func (s *BasicSpecies) EggCycles() int {
	return s.EGGCYCLES
}

func (s *BasicSpecies) SetEggCycles(eggCycles int) {
	s.EGGCYCLES = eggCycles
}

func (s *BasicSpecies) Height() float64 {
	return s.HEIGHT
}

func (s *BasicSpecies) SetHeight(height float64) {
	s.HEIGHT = height
}

func (s *BasicSpecies) Weight() float64 {
	return s.WEIGHT
}

func (s *BasicSpecies) SetWeight(weight float64) {
	s.WEIGHT = weight
}

func (s *BasicSpecies) ExperienceYield() int {
	return s.EXPERIENCEYIELD
}

func (s *BasicSpecies) SetExperienceYield(experienceYield int) {
	s.EXPERIENCEYIELD = experienceYield
}

func (s *BasicSpecies) ExperienceCurve() int {
	return s.EXPERIENCECURVE
}

func (s *BasicSpecies) SetExperienceCurve(experienceCurve int) {
	s.EXPERIENCECURVE = experienceCurve
}

func (s *BasicSpecies) EffortValues() map[string]int {
	return s.EFFORTVALUES
}

func (s *BasicSpecies) SetEffortValues(effortValues map[string]int) {
	s.EFFORTVALUES = effortValues
}

func (s *BasicSpecies) BodyStyle() string {
	return s.BODYSTYLE
}

func (s *BasicSpecies) SetBodyStyle(bodyStyle string) {
	s.BODYSTYLE = bodyStyle
}

func (s *BasicSpecies) Color() string {
	return s.COLOR
}

func (s *BasicSpecies) SetColor(color string) {
	s.COLOR = color
}

func (s *BasicSpecies) Tameness() int {
	return s.TAMENESS
}

func (s *BasicSpecies) SetTameness(tameness int) {
	s.TAMENESS = tameness
}

func (s *BasicSpecies) Stats() map[string]int {
	return s.STATS
}

func (s *BasicSpecies) SetStats(stats map[string]int) {
	s.STATS = stats
}

func (s *BasicSpecies) Evolutions() []int {
	return s.EVOLUTIONS
}

func (s *BasicSpecies) SetEvolutions(evolutions []int) {
	s.EVOLUTIONS = evolutions
}

func (s *BasicSpecies) MoveLists() map[int][]int {
	return s.MOVELISTS
}

func (s *BasicSpecies) SetMoveLists(moveLists map[int][]int) {
	s.MOVELISTS = moveLists
}

func (s *BasicSpecies) EggMoves() []int {
	return s.EGGMOVES
}

func (s *BasicSpecies) SetEggMoves(eggMoves []int) {
	s.EGGMOVES = eggMoves
}

func (s *BasicSpecies) MachineMoves() []int {
	return s.MACHINEMOVES
}

func (s *BasicSpecies) SetMachineMoves(machineMoves []int) {
	s.MACHINEMOVES = machineMoves
}

func (s *BasicSpecies) TutorMoves() []int {
	return s.TUTORMOVES
}

func (s *BasicSpecies) SetTutorMoves(tutorMoves []int) {
	s.TUTORMOVES = tutorMoves
}

func GetSpecies(speciesId int) Species {
	if AllSpecies[speciesId] == nil {
		species, _ := LoadSpecies(speciesId)

		AllSpecies[speciesId] = species
	}

	return AllSpecies[speciesId]
}

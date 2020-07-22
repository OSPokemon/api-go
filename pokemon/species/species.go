package species

import "github.com/ospokemon/api-go/pokedex/id"

type T struct {
	ID              *id.T
	Name            string
	Tag             string
	Description     string
	Types           []int
	Abilities       []int
	HiddenAbility   int
	GenderRatio     float64
	CatchRate       int
	EggGroups       []string
	EggCycles       int
	Height          float64
	Weight          float64
	ExperienceYield int
	ExperienceCurve int
	EffortValues    map[string]int
	BodyStyle       string
	Color           string
	Tameness        int
	Stats           map[string]float64
	Evolutions      []int
	MoveLists       map[int][]int
	EggMoves        []int
	MachineMoves    []int
	TutorMoves      []int
}

func New(id int, name string) *T {
	return &T{
		Name:         name,
		Types:        make([]int, 0),
		Abilities:    make([]int, 0),
		EggGroups:    make([]string, 0),
		EffortValues: make(map[string]int),
		Stats:        make(map[string]float64),
		Evolutions:   make([]int, 0),
		MoveLists:    make(map[int][]int),
		EggMoves:     make([]int, 0),
		MachineMoves: make([]int, 0),
		TutorMoves:   make([]int, 0),
	}
}

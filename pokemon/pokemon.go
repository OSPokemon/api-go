package pokemon

import (
	"github.com/ospokemon/api-go/pokedex/id"
	"github.com/ospokemon/api-go/pokemon/species"
	"github.com/ospokemon/api-go/skill"
	"github.com/ospokemon/api-go/stat"
)

type T struct {
	ID               *id.T
	Name             string
	Level            int
	Experience       int
	Ability          int
	Friendship       int
	Gender           int
	Nature           int
	Height           float64
	Weight           float64
	Trainer          int
	Shiny            bool
	Item             int
	Stats            map[string]stat.Stat
	Skills           []skill.Skill
	StatusConditions []int
	Markings         []string
	Ribbons          []string
}

func New(species *species.T) *T {
	return &T{
		ID:               species.ID,
		Stats:            make(map[string]stat.Stat),
		Skills:           make([]skill.Skill, 0),
		StatusConditions: make([]int, 0),
		Markings:         make([]string, 0),
		Ribbons:          make([]string, 0),
	}
}

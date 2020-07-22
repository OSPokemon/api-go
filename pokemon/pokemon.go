package pokemon

import (
	"github.com/ospokemon/api-go/pokedex/id"
	"github.com/ospokemon/api-go/pokemon/species"
	"github.com/ospokemon/api-go/skill"
	"github.com/ospokemon/api-go/stat"
	"github.com/ospokemon/api-go/status"
)

type T struct {
	ID         *id.T
	Name       string
	Level      int
	Experience int
	Ability    int
	Friendship int
	Gender     int
	Nature     int
	Height     float64
	Weight     float64
	Trainer    int
	Shiny      bool
	Item       int
	Stats      stat.Group
	Status     status.T
	Skills     [4]skill.T
	Markings   []string
	Ribbons    []string
}

func New(species *species.T) *T {
	return &T{
		ID:       species.ID,
		Skills:   [4]skill.T{},
		Status:   status.New(),
		Markings: make([]string, 0),
		Ribbons:  make([]string, 0),
	}
}

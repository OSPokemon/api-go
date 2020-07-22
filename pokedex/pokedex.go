package pokedex

import "github.com/ospokemon/api-go/pokedex/id"

type T struct {
	cache map[*id.T]bool
}

package pokedex

import (
	"pokedexcli/internal/pokeapi"
)

type Pokemon struct {
	Id      int
	Name    string
	BaseExp int
	Height  int
	Weight  int
	Stats   Stats
	Types   []string
}

type Stats struct {
	Hp             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}

type Pokedex struct {
	Caught map[string]Pokemon
}

func NewPokemon(name string, id, baseExp, height, weight int, stats Stats, types []string) Pokemon {
	return Pokemon{
		Id:      id,
		Name:    name,
		BaseExp: baseExp,
		Height:  height,
		Weight:  weight,
		Stats:   stats,
		Types:   types,
	}
}

func NewPokemonFromResp(resp pokeapi.RespPokemon) Pokemon {
	statsList := make(map[string]int, 6)
	for _, s := range resp.Stats {
		statsList[s.Stat.Name] = s.BaseStat
	}

	typesList := []string{}
	for _, t := range resp.Types {
		typesList = append(typesList, t.Type.Name)
	}

	return NewPokemon(
		resp.Name,
		resp.ID,
		resp.BaseExperience,
		resp.Height,
		resp.Weight,
		Stats{
			Hp:             statsList["hp"],
			Attack:         statsList["attack"],
			Defense:        statsList["defense"],
			SpecialAttack:  statsList["special-attack"],
			SpecialDefense: statsList["special-defense"],
			Speed:          statsList["speed"],
		},
		typesList,
	)
}

func NewPokedex() Pokedex {
	caught := make(map[string]Pokemon)
	return Pokedex{
		Caught: caught,
	}
}
func (pokedex Pokedex) Add(pokemon Pokemon) {
	pokedex.Caught[pokemon.Name] = pokemon
}

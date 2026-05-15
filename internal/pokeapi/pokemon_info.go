package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemonInfo(pokemonName string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	data, err := c.GetData(url)

	if err != nil {
		return RespPokemon{}, err
	}

	pokemonInfoResp := RespPokemon{}
	err = json.Unmarshal(data, &pokemonInfoResp)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemonInfoResp, nil

}

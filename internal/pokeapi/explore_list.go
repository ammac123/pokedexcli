package pokeapi

func (resp *RespLocationInfo) ListPokemonInLocation() []string {
	pokemonList := make([]string, len(resp.PokemonEncounters))
	for _, encounter := range resp.PokemonEncounters {
		pokemonList = append(pokemonList, encounter.Pokemon.Name)
	}
	return pokemonList
}

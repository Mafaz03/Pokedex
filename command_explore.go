package main

import (
	"errors"
	"fmt"
	"log"
)

func Command_Explore(cfg *config, location ...string) error {
	if len(location) != 1 {
		return errors.New("location not entered or not valid")
	}
	loc := location[0]
	client := cfg.pokeapi
	resp, err := client.ListPokeInLocation(loc)
	if err != nil {
		log.Fatal(err)
	}
	results := resp.PokemonEncounters
	for _, pokemon := range results {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}
	return err
}


package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func Command_Catch(cfg *config, pokemonName ...string) error {
	if len(pokemonName) != 1 {
		return errors.New("pokemon name not entered or not valid")
	}
	pokemon := pokemonName[0]
	client := cfg.pokeapi
	resp, err := client.PokemonsCatch(pokemon)
	if err != nil {
		log.Fatal(err)
	}
	const num = 50
	random_num := rand.Intn(resp.BaseExperience)
	if random_num > num {
		return fmt.Errorf("could not catch %s", pokemon)
	}
	fmt.Printf("Caught %v\n", pokemon)
	return nil
}


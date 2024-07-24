package main

import (
	"fmt"
	"log"

	"github.com/Mafaz03/Pokedex/internal/pokeapi"
)

func Command_Map() error {
	client := pokeapi.NewClient()
	resp, err := client.ListLocationArea()
	if err != nil {
		log.Fatal(err)
	}
	results := resp.Results
	for _, area:= range results{
		fmt.Printf(" - %v\n", area.Name)
	}
	return err
}

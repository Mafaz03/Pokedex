package main

import (
	"errors"
	"fmt"
	"log"
)

func Command_Map(cfg *config) error {
	client := cfg.pokeapi
	resp, err := client.ListLocationArea(cfg.nextURL)
	if err != nil {
		log.Fatal(err)
	}
	results := resp.Results
	for _, area := range results {
		fmt.Printf(" - %v\n", area.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous
	return err
}

func Command_Mapb(cfg *config) error {
	if cfg.prevURL == nil {
		return errors.New("you are on the first page")
	}
	client := cfg.pokeapi
	resp, err := client.ListLocationArea(cfg.prevURL)
	if err != nil {
		log.Fatal(err)
	}
	results := resp.Results
	for _, area := range results {
		fmt.Printf(" - %v\n", area.Name)
	}
	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous
	return err
}

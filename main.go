package main

import (
	"fmt"
	"log"

	"github.com/Mafaz03/Pokedex/internal/pokeapi"
)

func main() {
	// repl()
	client := pokeapi.NewClient()
	resp, err := client.ListLocationArea()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
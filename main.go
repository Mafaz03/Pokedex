package main

import (
	"time"

	"github.com/Mafaz03/Pokedex/internal/pokeapi"
)

// "fmt"
// "log"



type config struct {
	pokeapi pokeapi.Client
	prevURL *string
	nextURL *string
}

func main() {
	cfg := config{
		pokeapi: pokeapi.NewClient(time.Hour),
	}
	repl(&cfg)

}
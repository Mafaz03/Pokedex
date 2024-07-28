package main

import "fmt"

func Command_Help(cfg *config) error {
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

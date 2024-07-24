package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		commands := getCommands()
		cmd, ok := commands[command]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		cmd.callback()
	}
}

type commandsStruct struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]commandsStruct {
	return map[string]commandsStruct{
		"exit": {
			name:        "exit",
			description: "Exits from the CLI",
			callback:    Command_Exit,
		},
		"help": {
			name:        "help",
			description: "Gets quick walkthrough the commands",
			callback:    Command_Help,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

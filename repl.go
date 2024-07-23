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
		fmt.Println("Echoing: ", scanner.Text())
	}
}

func clearInput(str string) []string{
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

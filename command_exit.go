package main

import "os"

func Command_Exit(cfg *config) error {
	os.Exit(0)
	return nil
}

package main

import "os"

func Command_Exit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

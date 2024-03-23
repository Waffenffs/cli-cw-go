package main

import (
	"log"
	"os"
)

func main() {
	// Check for flags
	if len(os.Args) <= 1 {
		log.Fatal("\nUsage: go <program_name> <file_name>")
	}
	// Check for help command
	for _, flag := range os.Args {
		if flag == "--help" {
			log.Fatal("\nHelp command currently does not exist")
		}
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("\nFile: ", os.Args[1], " does not exist")
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal("\nCould not obtain file information")
	}
	log.Println("\nFile:", os.Args[1], "\nSize:", info.Size())

	defer os.Exit(0)
}

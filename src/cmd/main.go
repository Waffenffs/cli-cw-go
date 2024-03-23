package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	defer os.Exit(0)

	flags := os.Args

	if len(flags) <= 2 {
		log.Fatal("\nUsage: go <program_name> <command: --c, --l> <file_name>")
	}

	file, err := os.Open(flags[2])
	if err != nil {
		log.Fatal("\nFile does not exist")
	}

	for _, flag := range flags {
		switch flag {
		case "--help":
			log.Fatal("\nHelp command currently not implemented")
		case "--c":
			getFileSize(file)
		case "--l":
			getFileTotalLines(file)
		}
	}

	file.Close()
}

func getFileSize(f *os.File) {
	info, err := f.Stat()
	if err != nil {
		log.Fatal("\nCould not obtain file information")
	}

	log.Println("\nFile:", os.Args[2], "\nSize:", info.Size())
}

func getFileTotalLines(f *os.File) {
	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		line++
	}

	log.Println("Total lines:", line)
}

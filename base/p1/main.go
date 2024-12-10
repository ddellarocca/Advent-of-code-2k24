package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("partial.txt")
	// file, err := os.Open("full.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// TODO
	}
}

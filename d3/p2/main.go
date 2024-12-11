package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// file, err := os.ReadFile("partial.txt")
	file, err := os.ReadFile("full.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	content := string(file)

	r, _ := regexp.Compile(`(mul\(([0-9]{0,3}),([0-9]{0,3})\))|(do\(\))|(don\'t\(\))`)
	matches := r.FindAllStringSubmatch(content, -1)

	fmt.Println(matches)

	sum := 0
	sumEnabled := true
	for _, match := range matches {
		fmt.Println(match[0])

		switch match[0] {
		case "do()":
			fmt.Println("enabling")
			sumEnabled = true
		case "don't()":
			fmt.Println("disabling")
			sumEnabled = false
		default:
			if sumEnabled {
				fmt.Println("summing")
				first, _ := strconv.Atoi(match[2])
				second, _ := strconv.Atoi(match[3])
				sum += first * second
			} else {
				fmt.Println("skipping")
			}
		}
	}

	fmt.Println("total:", sum)
}

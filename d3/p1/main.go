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

	r, _ := regexp.Compile(`mul\(([0-9]{0,3}),([0-9]{0,3})\)`)
	muls := r.FindAllStringSubmatch(content, -1)

	sum := 0
	for _, mul := range muls {
		first, _ := strconv.Atoi(mul[1])
		second, _ := strconv.Atoi(mul[2])
		sum += first * second
	}

	fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("partial.txt")
	file, err := os.Open("full.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	var left []int
	var right map[int]int
	right = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		val1, err1 := strconv.Atoi(parts[0])
		val2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing line", line)
		}

		left = append(left, val1)
		right[val2] += 1
	}

	score := 0

	for i := 0; i < len(left); i++ {
		score += left[i] * right[left[i]]
	}

	fmt.Println(score)
}

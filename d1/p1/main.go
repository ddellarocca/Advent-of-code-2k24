package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var first, second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		val1, err1 := strconv.Atoi(parts[0])
		val2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing line", line)
		}

		first = append(first, val1)
		second = append(second, val2)
	}

	sort.Slice(first, func(i, j int) bool {
		return first[i] < first[j]
	})

	sort.Slice(second, func(i, j int) bool {
		return second[i] < second[j]
	})

	diff := 0

	for i := 0; i < len(first); i++ {
		diff += absint(first[i], second[i])
		fmt.Println(first[i], second[i], absint(first[i], second[i]))
	}

	fmt.Println(diff)
}

func absint(x int, y int) (z int) {
	if x < y {
		z = y - x
		return
	}
	z = x - y

	return
}

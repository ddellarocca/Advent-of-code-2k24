package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	// file, err := os.Open("partial.txt")
	file, err := os.Open("full.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	antennas := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		antennas = append(antennas, strings.Split(line, ""))
	}

	h := len(antennas)
	l := len(antennas[0])

	out := make([][]string, h)
	for i := range h {
		out[i] = slices.Repeat([]string{"."}, l)
	}

	total := 0
	for y := range h {
		for x := range l {
			char := antennas[y][x]

			if char != "." && char != "#" {
				if char == "0" {
					fmt.Println(y, x)
				}
				// search for other equal chars
				for ty := 0; ty < h; ty++ {
					for tx := 0; tx < l; tx++ {
						next := antennas[ty][tx]
						if next == char && (ty != y && tx != x) {
							dy, dx := ty-y, tx-x

							uy, ux := y-dy, x-dx
							by, bx := y+dy, x+dx

							if isInRange(uy, h) && isInRange(ux, l) && antennas[uy][ux] != char {
								out[uy][ux] = "#"
							}

							if isInRange(by, h) && isInRange(bx, l) && antennas[by][bx] != char {
								out[by][bx] = "#"
							}
						}
					}
				}
			}
		}
	}

	for _, ln := range out {
		fmt.Println(ln)
		for _, c := range ln {
			if c == "#" {
				total++
			}
		}
	}

	fmt.Println(total)
}

func isInRange(val int, max int) bool {
	if val < 0 || val >= max {
		return false
	}
	return true
}

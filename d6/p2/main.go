package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("partial.txt")
	file, err := os.Open("full.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	original := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		original = append(original, strings.Split(line, ""))
	}
	height := len(original)
	length := len(original[0])
	loops := 0

	checks := 0

	for h := range height {
		for l := range length {

			view := make([][]string, height)
			for i := range original {
				view[i] = make([]string, length)
				copy(view[i], original[i])
			}

			r, c := findGuard(original)

			view[h][l] = "#"

			checks++
			loops += hasLoop(view, r, c)
		}
	}

	fmt.Println("loops:", loops)
}

func hasLoop(view [][]string, r int, c int) int {
	height := len(view)
	length := len(view[0])
	directionChanges := make(map[string]bool)
	direction := "up"
	for {
		view[r][c] = "X"
		if checkBoundaries(r, c, height-1, length-1) {
			return 0
		}

		// If i end up in a cell two times with the same direction then it is a loop
		if directionChanges[string(r)+direction+string(c)] == true {
			return 1
		}
		directionChanges[string(r)+direction+string(c)] = true

		switch direction {
		case "up":
			if view[r-1][c] == "#" {
				direction = "right"
			} else {
				r--
			}
		case "right":
			if view[r][c+1] == "#" {
				direction = "bottom"
			} else {
				c++
			}
			break
		case "bottom":
			if view[r+1][c] == "#" {
				direction = "left"
			} else {
				r++
			}
			break
		case "left":
			if view[r][c-1] == "#" {
				direction = "up"
			} else {
				c--
			}
			break
		default:
			break
		}
	}
}

func checkBoundaries(r int, c int, h int, l int) bool {
	return (r == 0 || c == 0 || r == h || c == l)
}

func findGuard(view [][]string) (r int, c int) {
	var row []string
	for r, row = range view {
		for c, _ = range row {
			if view[r][c] == "^" {
				return
			}
		}
	}
	return
}

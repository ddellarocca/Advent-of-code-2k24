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

	view := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		view = append(view, strings.Split(line, ""))
	}

	for _, row := range view {
		fmt.Println(row)
	}

	r, c := findGuard(view)
	fmt.Println("guard position:", r, c)

	direction := "up"
	height := len(view) - 1
	length := len(view[0]) - 1
	for {
		view[r][c] = "X"
		if checkBoundaries(r, c, height, length) {
			break
		}
		fmt.Println("current:", direction, r, c)

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

	for _, row := range view {
		fmt.Println(row)
	}

	fmt.Println("total unique steps:", countUnique(view))
}

func checkBoundaries(r int, c int, h int, l int) bool {
	fmt.Println(r, c, h, l)
	return (r == 0 || c == 0 || r == h || c == l)
}

func countUnique(view [][]string) (u int) {
	for _, row := range view {
		for _, col := range row {
			if col == "X" {
				u++
			}
		}
	}
	return
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

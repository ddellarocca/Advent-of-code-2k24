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

	scanner := bufio.NewScanner(file)
	topoMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		values := []int{}
		for _, v := range strings.Split(line, "") {
			vp, _ := strconv.Atoi(v)
			values = append(values, vp)
		}

		topoMap = append(topoMap, values)
	}

	total := 0
	for y, l := range topoMap {
		for x, v := range l {
			if v == 0 {
				topoCopy := make([][]int, len(topoMap))
				for i := range topoMap {
					topoCopy[i] = make([]int, len(topoMap[0]))
					copy(topoCopy[i], topoMap[i])
				}
				total += getTrailScore(0, y, x, topoCopy)
			}
		}
	}

	fmt.Println(total)
}

func getTrailScore(actual int, y int, x int, topoMap [][]int) (total int) {
	if actual == 9 {
		return 1
	}

	top := y - 1
	if top >= 0 && topoMap[top][x] == actual+1 {
		total += getTrailScore(actual+1, top, x, topoMap)
	}

	right := x + 1
	if right < len(topoMap[0]) && topoMap[y][right] == actual+1 {
		total += getTrailScore(actual+1, y, right, topoMap)
	}

	bottom := y + 1
	if bottom < len(topoMap) && topoMap[bottom][x] == actual+1 {
		total += getTrailScore(actual+1, bottom, x, topoMap)
	}

	left := x - 1
	if left >= 0 && topoMap[y][left] == actual+1 {
		total += getTrailScore(actual+1, y, left, topoMap)
	}

	return
}

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

	var safeReports int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		levelsString := strings.Fields(report)

		var levels []int
		for _, valueString := range levelsString {
			value, _ := strconv.Atoi(valueString)
			levels = append(levels, value)
		}

		fmt.Println(levels)

		if validReport(levels) {
			safeReports += 1
			fmt.Println("Report safe")
		} else {
			fmt.Println("Report not safe, trying other configurations")
			for index, _ := range levels {
				copied := make([]int, len(levels))
				copy(copied, levels)
				partial := remove(copied, index)
				if validReport(partial) {
					safeReports += 1
					fmt.Println("Report safe after one config change")
					break
				}
			}
		}

	}
	fmt.Println(safeReports)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func validReport(levels []int) bool {
	var diffs []int
	for i := 0; i < len(levels)-1; i++ {
		diffs = append(diffs, levels[i]-levels[i+1])
	}

	safe := true

	startingTrend := diffs[0] > 0
	for _, value := range diffs {
		if absint(value) < 1 || absint(value) > 3 {
			safe = false
			fmt.Println("Difference outside boundaries")
			break
		}

		currentTrend := value > 0

		if startingTrend != currentTrend {
			safe = false
			fmt.Println("Trend changed")
			break
		}
	}

	return safe
}

func outsideBoundaries(value int) bool {
	return absint(value) < 1 || absint(value) > 3
}

func absint(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

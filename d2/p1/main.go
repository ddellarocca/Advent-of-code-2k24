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
		levels := strings.Fields(report)

		fmt.Println(levels)

		var diffs []int
		for i := 0; i < len(levels)-1; i++ {
			current, _ := strconv.Atoi(levels[i])
			next, _ := strconv.Atoi(levels[i+1])
			diffs = append(diffs, current-next)
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

		if safe {
			fmt.Println("Report safe")
			safeReports += 1
		}
	}
	fmt.Println(safeReports)
}

func absint(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

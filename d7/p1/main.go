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
	lines := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ":", "")

		linePieces := strings.Fields(line)

		values := []int{}
		for _, piece := range linePieces {
			value, _ := strconv.Atoi(piece)
			values = append(values, value)
		}

		lines = append(lines, values)
	}

	total := 0
	for _, line := range lines {
		if solveLine(line[0], line[1:]) {
			total += line[0]
		}
	}

	fmt.Println(total)

	// solveLine(lines[0][0], lines[0][1:])
}

func solveLine(result int, operands []int) bool {
	numOperands := len(operands)
	numOperators := numOperands - 1
	numRun := intPow(2, numOperators)

	for i := range numRun {
		calc := operands[0]
		for j := range numOperators {
			if i>>j&1 == 0 {
				calc += operands[j+1]
			} else {
				calc *= operands[j+1]
			}
		}
		if calc == result {
			fmt.Println(calc, operands)
			return true
		}
	}

	return false
}

func intPow(num int, pow int) (res int) {
	res = 1
	for range pow {
		res *= num
	}
	return
}

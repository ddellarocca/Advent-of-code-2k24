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
	fmt.Println(lines)

	total := 0
	for _, line := range lines {
		if solveLine(line[0], line[1:]) {
			total += line[0]
		}
	}

	fmt.Println(total)

	// solveLine(lines[1][0], lines[1][1:])
}

func solveLine(result int, operands []int) bool {
	numOperands := len(operands)
	numOperators := numOperands - 1
	numRun := intPow(3, numOperators)

	for i := range numRun {
		calc := operands[0]
		base3, _ := strconv.Atoi(strconv.FormatInt(int64(i), 3))
		for j := range numOperators {
			op := base3 % 10
			if op == 0 {
				calc += operands[j+1]
			} else if op == 1 {
				calc *= operands[j+1]
			} else if op == 2 {
				calcs := strconv.FormatInt(int64(calc), 10)
				next := strconv.FormatInt(int64(operands[j+1]), 10)
				calc, _ = strconv.Atoi(calcs + next)
			}
			base3 /= 10
		}
		if calc == result {
			fmt.Println(result, operands, calc)
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

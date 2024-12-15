package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	lines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ":", "")

		lines = append(lines, strings.Fields(line))
	}

	var total int64 = 0
	for _, line := range lines {
		result, _ := strconv.ParseInt(line[0], 10, 64)
		if solveLine(result, line[1:]) {
			total += result
		}
	}

	fmt.Println(total)

	// result, _ := strconv.ParseInt(lines[8][0], 10, 64)
	// solveLine(result, lines[8][1:])
}

func solveLine(result int64, operands []string) bool {
	numOperands := len(operands)
	numOperators := numOperands - 1
	numRun := intPow(2, numOperators)

	// fmt.Println(numOperands, numOperators, numRun)

	decode := map[int]string{
		0: "+",
		1: "*",
	}

	for i := range numRun {
		expr := []string{}
		opStack := []int{}

		for j := range numOperands {
			expr = append(expr, operands[j])

			currOp := i >> j & 1 // 1 means mul, 0 means add
			// fmt.Println(j, operands, currOp, i, j)

			if len(opStack) == 0 {
				opStack = append(opStack, i>>j&1)
			} else if opStack[len(opStack)-1] < currOp {
				opStack = append(opStack, currOp)
			} else {
				for o := len(opStack) - 1; o >= 0; o-- {
					expr = append(expr, decode[opStack[o]])
				}
				opStack = []int{}
				if j < numOperands-1 {
					opStack = append(opStack, currOp)
				}
			}
			// fmt.Println(expr, currOp, opStack)
		}
		// fmt.Println(j, operands, currOp, i, j)
		// fmt.Println(expr, i, numOperators)
		fmt.Println(result, ":", expr, "--", calculate(expr))
		if calculate(expr) == result {
			// fmt.Println(result, ":", expr)
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

func calculate(expr []string) (res int64) {
	operands := []int64{}
	for _, item := range expr {
		if op := []string{"+", "*"}; !slices.Contains(op, item) {
			parsed, _ := strconv.ParseInt(item, 10, 64)
			operands = append(operands, parsed)
		} else {
			ops := len(operands)
			if item == "+" {
				operands[ops-2] += operands[ops-1]
			} else {
				operands[ops-2] *= operands[ops-1]
			}
			operands = operands[:ops-1]
		}
	}
	return operands[0]
}

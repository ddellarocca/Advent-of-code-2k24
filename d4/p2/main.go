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

	scanner := bufio.NewScanner(file)

	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	totalXmases := 0
	for line := 1; line < len(matrix)-1; line++ {
		chars := matrix[line]
		for index := 1; index < len(chars)-1; index++ {
			if chars[index] == "A" {
				totalXmases += findXmas(matrix, line, index)
			}
		}
	}

	fmt.Println("total:", totalXmases)
}

func findXmas(matrix [][]string, line int, index int) (xmases int) {
	// we found an A we need to check whether the neighbours can compose a MAS

	// check diagonal \
	dec := checkXmasFromIncrement(matrix, line, -1, index, -1)
	// check diagonal /
	inc := checkXmasFromIncrement(matrix, line, -1, index, 1)

	return btoi(dec && inc) // if there are both diagonals then it's an x-mas
}

func checkXmasFromIncrement(matrix [][]string, startLine int, incLine int, startIndex int, incIndex int) bool {
	text := matrix[startLine+incLine][startIndex+incIndex]
	text += matrix[startLine][startIndex]
	text += matrix[startLine-incLine][startIndex-incIndex]

	return (text == "MAS" || text == "SAM")
}

func btoi(b bool) (i int) {
	if b {
		return 1
	}
	return 0
}

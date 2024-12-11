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
	for line, chars := range matrix {
		for index, char := range chars {
			if char == "X" {
				totalXmases += findXmas(matrix, line, index)
			}
		}
	}

	fmt.Println("total:", totalXmases)
}

func findXmas(matrix [][]string, line int, index int) (xmases int) {
	// we found an X we need to check whether the neighbours can compose an XMAS

	heightLenght := len(matrix)
	lineLenght := len(matrix[line])

	searchLeft := index >= 3
	searchRight := index < lineLenght-3
	searchTop := line >= 3
	searchBottom := line < heightLenght-3

	if searchLeft {
		// check left
		xmases += checkXmasFromIncrement(matrix, line, 0, index, -1)

		if searchTop {
			// check top-left
			xmases += checkXmasFromIncrement(matrix, line, -1, index, -1)
		}

		if searchBottom {
			// check bottom-left
			xmases += checkXmasFromIncrement(matrix, line, 1, index, -1)
		}
	}

	if searchTop {
		// check top
		xmases += checkXmasFromIncrement(matrix, line, -1, index, 0)
	}

	if searchRight {
		// check right
		xmases += checkXmasFromIncrement(matrix, line, 0, index, 1)

		if searchTop {
			// check top-right
			xmases += checkXmasFromIncrement(matrix, line, -1, index, 1)
		}

		if searchBottom {
			// check bottom-right
			xmases += checkXmasFromIncrement(matrix, line, 1, index, 1)
		}
	}

	if searchBottom {
		// check bottom
		xmases += checkXmasFromIncrement(matrix, line, 1, index, 0)
	}

	return
}

func checkXmasFromIncrement(matrix [][]string, startLine int, incLine int, startIndex int, incIndex int) int {
	var text string
	for i := 0; i < 4; i++ {
		moveLine := startLine + incLine*i
		moveIndex := startIndex + incIndex*i
		text += matrix[moveLine][moveIndex]
	}

	return btoi(text == "XMAS")
}

func btoi(b bool) (i int) {
	if b {
		return 1
	}
	return 0
}

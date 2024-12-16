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
	fsmap := []int{}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range strings.Split(line, "") {
			vp, _ := strconv.Atoi(v)
			fsmap = append(fsmap, vp)
			total += vp
		}
	}

	fsrep := slices.Repeat([]int{-1}, total)
	shift := 0
	for i, v := range fsmap {
		for s := range v {
			// if it is not empty space
			if i%2 == 0 {
				fsrep[s+shift] = i / 2
			}
		}
		shift += v
	}

	// until we have an empty space
	i := 0
	for {
		if i == total {
			break
		}

		if fsrep[i] == -1 {
			for j := total - 1; j > i; j-- {
				v := fsrep[j]
				if v != -1 {
					fsrep[i] = v
					fsrep[j] = -1
					break
				}
			}
		}

		i++
	}

	fmt.Println(fsrep)

	checksum := 0
	for i, v := range fsrep {
		if v == -1 {
			break
		}
		checksum += i * v
	}

	fmt.Println(checksum)
}

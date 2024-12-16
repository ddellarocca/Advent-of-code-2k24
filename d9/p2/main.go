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
	file, err := os.Open("full.txt") //6415163624282
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

	j := total - 1

	// fmt.Println(fsrep)

	for {
		i := 0
		if j <= 0 {
			break
		}

		// find first file to move
		fi, fl := findLastFull(fsrep, j)
		// fmt.Println("found file", fi+1, fl)

		ei, el := findFirstEmpty(fsrep, i)
		// fmt.Println("first empty block", ei+1, el, i+1)

		if ei >= fi {
			// fmt.Println("first empty block is after the next file. aka no space left")
			break
		}

		for {
			if el >= fl {
				// fmt.Println("file fit, inserting file", ei+1, fi+1, el, fl)
				for fn := 0; fn < fl; fn++ {
					fsrep[ei+fn] = fsrep[fi-fn]
					fsrep[fi-fn] = -1
				}
				// fmt.Println(fsrep)
				break
			} else {
				// fmt.Println("block too small")
			}

			i = (el + ei + 1)

			if i >= fi {
				break
			}

			if i >= total {
				break
			}

			ei, el = findFirstEmpty(fsrep, i)
			if ei < 0 {
				break
			}
			// fmt.Println("next empty block", ei+1, el, i+1)

			if ei >= fi {
				break
			}
		}

		j = (fi - fl)

		if j < 1 {
			break
		}
	}

	checksum := 0
	for i, v := range fsrep {
		if v == -1 {
			continue
		}
		checksum += i * v
	}

	fmt.Println(checksum)

}

func findFirstEmpty(values []int, starting int) (index int, lenght int) {
	index = starting
	for {
		if values[index] == -1 {
			break
		}
		index++
	}

	lenght = 1
	for {
		if values[index+lenght] != -1 {
			return
		}
		lenght++

		if lenght+index >= len(values) {
			return -1, -1
		}
	}
}

func findLastFull(values []int, starting int) (index int, lenght int) {
	index = starting
	findex := -1
	for {
		findex = values[index]
		if findex != -1 {
			// fmt.Println(findex)
			break
		}
		index--
	}

	lenght = 1
	for {
		if values[index-lenght] != findex {
			return
		}
		lenght++
	}
}

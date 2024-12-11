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
	isRule := true
	rules := make(map[string][]string)
	prints := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRule = false
			continue
		}

		if isRule {
			rule := strings.Split(line, "|")

			rules[rule[1]] = append(rules[rule[1]], rule[0])
		} else {
			prints = append(prints, strings.Split(line, ","))
		}
	}

	fmt.Println(rules)

	sumMiddle := 0
	for _, print := range prints {
		if res, i, j := checkValidity(print, rules); !res {
			fixed := fix(print, rules, i, j)
			sumMiddle += getMiddle(fixed)
		}
	}

	fmt.Println("total:", sumMiddle)
}

// basically we switch position of the ones making the check to fail until a valid sequence is found
func fix(print []string, rules map[string][]string, i int, j int) (fixed []string) {
	res := false
	fixed = make([]string, len(print))
	fmt.Println()
	si := i
	sj := j
	copy(fixed, print)
	for {
		tmp := fixed[sj]
		fixed[sj] = fixed[si]
		fixed[si] = tmp

		res, si, sj = checkValidity(fixed, rules)
		fmt.Println(res, si, sj, fixed)

		if res {
			break
		}
	}
	return
}

func checkValidity(print []string, rules map[string][]string) (res bool, i int, j int) {
	for ip, page := range print {
		// skip the last one since it doesn't have a next
		if ip == len(print)-1 {
			break
		}
		rulesForPage := rules[page]

		// if the next element should be before the current one then is not valid
		for it := ip + 1; it < len(print); it++ {
			for _, rule := range rulesForPage {
				if rule == print[it] {
					return false, ip, it
				}
			}
		}
	}

	return true, 0, 0
}

func getMiddle(slc []string) (value int) {
	value, _ = strconv.Atoi(slc[(len(slc)-1)/2])
	return
}

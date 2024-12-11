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
	fmt.Println(prints)

	sumMiddle := 0
	for _, print := range prints {
		if checkValidity(print, rules) {
			sumMiddle += getMiddle(print)
		}
	}

	fmt.Println("total:", sumMiddle)
}

func checkValidity(print []string, rules map[string][]string) bool {
	for index, page := range print {
		// skip the last one since it doesn't have a next
		if index == len(print)-1 {
			break
		}
		rulesForPage := rules[page]

		// if the next element should be before the current one then is not valid
		for i := index + 1; i < len(print); i++ {
			for _, rule := range rulesForPage {
				if rule == print[i] {
					return false
				}
			}
		}
	}

	return true
}

func getMiddle(slc []string) (value int) {
	value, _ = strconv.Atoi(slc[(len(slc)-1)/2])
	return
}

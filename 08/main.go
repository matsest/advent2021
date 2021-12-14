package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2021/utils"
)

func readNumber(s string) int {

	// cf len 2 -> 1
	// acf len 3 -> 7
	// bcdf len 4 -> 4
	// abcdefg len 7 -> 8

	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}

	return 0
}

func part1(entries []string) int {
	count := 0

	for _, e := range entries {
		digits := strings.Split(e, " ")
		for _, d := range digits {
			if readNumber(d) != 0 {
				count += 1
			}
		}
	}
	return count
}

func main() {
	lines, _ := utils.ReadLines("test.txt")

	// Parse input
	var outputvalueentries []string
	for _, v := range lines {
		parts := strings.Split(v, " | ")
		outputvalueentries = append(outputvalueentries, parts[1])
	}
	//fmt.Println(outputvalues)

	// Part 1
	p1 := part1(outputvalueentries)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(numbers)
	//fmt.Println("Part 2: ", p2)
}

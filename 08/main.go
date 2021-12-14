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

func part1(digits []string) int {
	count := 0
	for _, d := range digits {
		if readNumber(d) != 0 {
			count += 1
		}
	}
	return count
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Parse input
	var outputvalues []string
	for _, v := range lines {
		parts := strings.Split(v, " | ")
		digits := strings.Split(parts[1], " ")
		outputvalues = append(outputvalues, digits...)
	}
	//fmt.Println(outputvalues)

	// Part 1
	p1 := part1(outputvalues)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(numbers)
	//fmt.Println("Part 2: ", p2)
}

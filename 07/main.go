package main

import (
	"fmt"

	"github.com/matsest/advent2021/utils"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(crabs []int) int {

	fmt.Println(crabs)

	shortestRange := 9999999

	for i, _ := range crabs {
		sum := 0
		for j, v := range crabs {
			if i != j {
				sum += Abs((v - crabs[i]))
			}
		}
		if sum < shortestRange {
			shortestRange = sum
		}
		//fmt.Println(crabs[i], sum)
	}

	return shortestRange
}

func main() {
	lines, _ := utils.ReadInts("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(lines)
	//fmt.Println("Part 2: ", p2)
}

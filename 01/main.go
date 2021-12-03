package main

import (
	"fmt"
	"strconv"

	"github.com/matsest/advent2021/utils"
)

func part1(numbers []int) int {
	count := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			count++
		}
	}
	return count
}

func sumWindow(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func part2(numbers []int) int {
	count := 0
	for i := 1; i < len(numbers)-2; i++ {
		previousWindow := sumWindow(numbers[i-1 : i+2])
		currentWindow := sumWindow(numbers[i : i+3])
		if currentWindow > previousWindow {
			count++
		}
	}
	return count
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Convert to ints
	var numbers []int
	for _, num := range lines {
		number, _ := strconv.Atoi(num)
		numbers = append(numbers, number)
	}

	// Part 1
	p1 := part1(numbers)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(numbers)
	fmt.Println("Part 2: ", p2)
}

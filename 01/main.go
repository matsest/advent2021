package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

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
	lines, _ := readLines("input.txt")

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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func move(x int, y int, direction string, lengthString string) (int, int) {
	length, _ := strconv.Atoi(lengthString)
	switch direction {
	case "forward":
		return (x + length), y
	case "up":
		return x, (y + length)
	case "down":
		return x, (y - length)
	default:
		return x, y
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(directions []string) int {
	x, y := 0, 0
	for _, line := range directions {
		arr := strings.Split(line, " ")
		direction, length := arr[0], arr[1]
		x, y = move(x, y, direction, length)
	}

	return Abs(x) * Abs(y)
}

func main() {
	lines, _ := readLines("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(numbers)
	//fmt.Println("Part 2: ", p2)
}

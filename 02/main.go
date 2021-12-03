package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matsest/advent2021/utils"
)

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

func move2(x int, y int, direction string, lengthString string, aim int) (int, int, int) {
	length, _ := strconv.Atoi(lengthString)
	switch direction {
	case "forward":
		return (x + length), (y-(aim*length)), aim
	case "up":
		return x, y, (aim - length)
	case "down":
		return x, y, (aim + length)
	default:
		return x, y, aim
	}
}

func part2(directions []string) int {
	x, y, aim := 0, 0, 0
	for _, line := range directions {
		arr := strings.Split(line, " ")
		direction, length := arr[0], arr[1]
		x, y, aim = move2(x, y, direction, length, aim)
	}

	return Abs(x) * Abs(y)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(lines)
	fmt.Println("Part 2: ", p2)
}

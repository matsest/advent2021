package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2021/utils"
)

type Entry struct {
	this       int // value
	over       int
	under      int
	left       int
	right      int
	isLowPoint bool
}

func (e *Entry) SetLowPoint() bool {
	if (e.this < e.over) &&
		(e.this < e.under) &&
		(e.this < e.left) &&
		(e.this < e.right) {
		e.isLowPoint = true
	}
	return e.isLowPoint
}

func part1(numbers [][]int) int {

	//for _, r := range numbers {
	//	fmt.Println(r)
	//}

	riskLevelSum := 0

	for i, row := range numbers {
		for j := range row {
			// set neighbours to 10 by default have a simple check
			test := Entry{numbers[i][j], 10, 10, 10, 10, false}

			if i > 0 {
				test.over = numbers[i-1][j]
			}
			if i < len(numbers)-1 {
				test.under = numbers[i+1][j]
			}

			if j > 0 {
				test.left = numbers[i][j-1]
			}
			if j < len(row)-1 {
				test.right = numbers[i][j+1]
			}

			test.SetLowPoint()
			if test.isLowPoint {
				riskLevelSum += (test.this + 1)
			}
		}
	}

	return riskLevelSum
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Parse input
	var numbers [][]int
	for _, v := range lines {
		row := strings.Split(v, "")
		rowInt, _ := utils.SliceAtoi(row)
		numbers = append(numbers, rowInt)
	}

	// Part 1
	p1 := part1(numbers)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(numbers)
	//fmt.Println("Part 2: ", p2)
}

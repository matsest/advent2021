package main

import (
	"fmt"
	"strings"

	"github.com/matsest/advent2021/utils"
)

type Coordinate struct {
	x int
	y int
}

func convertCoordinates(coordinates []string) (out [][]Coordinate) {

	for _, v := range coordinates {
		//fmt.Println(v)
		split := strings.Split(v, " -> ")
		c1, c2 := strings.Split(split[0], ","), strings.Split(split[1], ",")

		c1Ints, _ := utils.SliceAtoi(c1)
		c2Ints, _ := utils.SliceAtoi(c2)

		x1, y1 := c1Ints[0], c1Ints[1]
		x2, y2 := c2Ints[0], c2Ints[1]
		//fmt.Println(x1,y1, "----", x2, y2)
		out = append(out, []Coordinate{{x1, y1}, {x2, y2}})
	}

	return out
}

func MaxInt(x int, y int) int {
	if y > x {
		return y
	}
	return x
}

func MinInt(x int, y int) int {
	if y > x {
		return x
	}
	return y
}

func part1(coordinates []string) int {

	coords := convertCoordinates(coordinates)
	//fmt.Println(coords)

	// Create diagram
	diagram := [999][999]int{}
	for i := 0; i < 999; i++ {
		for j := 0; j < 999; j++ {
			diagram[i][j] = 0
		}
	}
	//fmt.Println(diagram)

	for _, c := range coords {
		//fmt.Println(c)
		//fmt.Println("line (x2-x1; y2-y1): ", c[1].x - c[0].x, c[1].y - c[1].y)
		if c[1].x-c[0].x == 0 {
			//fmt.Println("found straight vertical line", c)
			yMax := MaxInt(c[1].y, c[0].y)
			yMin := MinInt(c[1].y, c[0].y)
			for y := yMin; y <= yMax; y++ {
				diagram[c[1].x][y] += 1
			}
		} else if c[1].y-c[0].y == 0 {
			//fmt.Println("found straight horizontal line", c)
			xMax := MaxInt(c[1].x, c[0].x)
			xMin := MinInt(c[1].x, c[0].x)
			for x := xMin; x <= xMax; x++ {
				diagram[x][c[1].y] += 1
			}
		}
	}
	//fmt.Println(diagram)

	sum := 0
	for _, row := range diagram {
		//fmt.Println(row)
		for _, e := range row {
			if e > 1 {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	lines, _ := utils.ReadLines("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	//p2 := part2(lines)
	//fmt.Println("Part 2: ", p2)
}

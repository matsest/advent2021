package main

import (
	"fmt"
	"regexp"

	"github.com/matsest/advent2021/utils"
)

func shootTrajectory(xMin, xMax, yMin, yMax, xv, yv int) (bool, int) {

	x, y := 0, 0
	yRecord := y

	for {
		x = x + xv
		y = y + yv

		// update x velo
		if xv > 0 {
			xv -= 1
		} else if xv < 0 {
			xv += 1
		}

		// update yv
		yv -= 1

		// update yRecord
		if y > yRecord {
			yRecord = y
		}

		//fmt.Println("x, y: ", x, y, "xv, yv: ", xv, yv)
		if (xMin <= x) && (x <= xMax) && (yMin <= y) && (y <= yMax) {
			return true, yRecord
		}

		// break if it does not go anywhere.. (along x)
		if xv == 0 && (x > xMax || x < xMin) {
			return false, 0
		}

		// break if it is does not go anywhere (along y)
		if yv < 0 && y < yMin {
			return false, 0
		}
	}
}

func part1(xMin, xMax, yMin, yMax int) int {

	yRecord := -10
	xvRecord, yvRecord := 1,1
	for xv:= 1; xv < 300; xv++ {
		for yv:=-10; yv < 300; yv++ {
			ok, res := shootTrajectory(xMin, xMax, yMin, yMax, xv, yv)
			if  ok && res > yRecord {
				yRecord = res
				xvRecord, yvRecord = xv, yv
			}
		}
	}
	fmt.Println("Using starting xv, yv: ", xvRecord, yvRecord)
	return yRecord
}

type StartingValue struct {
	xv, yv int
}

func part2(xMin, xMax, yMin, yMax int) int {

	startingValues := []StartingValue{}
	for xv:= 1; xv < 300; xv++ {
		for yv:=-300; yv < 300; yv++ {
			res, _ := shootTrajectory(xMin, xMax, yMin, yMax, xv, yv)
			if  res {
				startingValues = append(startingValues, StartingValue{xv, yv})
			}
		}
	}
	return len(startingValues)
}

func main() {
	lines, _ := utils.ReadLines("input.txt")
	re := regexp.MustCompile("-?[0-9]+")
	ns := re.FindAllString(lines[0], -1)
	n, _ := utils.SliceAtoi(ns)
	xMin, xMax, yMin, yMax := n[0], n[1], n[2], n[3]

	// Part 1
	p1 := part1(xMin, xMax, yMin, yMax)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(xMin, xMax, yMin, yMax)
	fmt.Println("Part 2: ", p2)
}

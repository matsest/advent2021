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

	//fmt.Println(crabs)

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

func AddMoveCost(n int, m int) int {

	diff := Abs(n - m)

	sum := 0

	for i := 1; i < diff; i++ {
		sum += i
	}
	//fmt.Println("diff ", diff, "additional ", sum)

	return diff + sum
}


func part2(crabs []int) int {

	//fmt.Println(crabs)

	shortestRange := 99999999999

	// Find max element in array
    max := crabs[0]
    for i := 1; i < len(crabs); i++ {
        if max < crabs[i] {
            max = crabs[i]
        }
    }
	//fmt.Println("largest number ", max)

	//for i := range crabs {
	for i:= 0; i < max; i++ {
		sum := 0
		//for j:= 0; j < max; j++ {
		for _, v := range crabs {
			if i != v {
				moveCost := AddMoveCost(v, i)
				//fmt.Println("Move ", j, "to ", crabs[i], ":", moveCost)
				sum += moveCost
			}
		}
		if sum < shortestRange {
			shortestRange = sum
		}
		//fmt.Println(i, sum)
	}

	return shortestRange
}

func main() {
	lines, _ := utils.ReadInts("input.txt")

	// Part 1
	p1 := part1(lines)
	fmt.Println("Part 1: ", p1)

	// Part 2
	p2 := part2(lines)
	fmt.Println("Part 2: ", p2)
}

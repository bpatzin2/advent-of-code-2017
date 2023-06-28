package main

import (
	"advent-of-code-2017/day1"
	"advent-of-code-2017/day7"
	"fmt"
)

func main() {
	fmt.Println("Day1 Part1: ", day1pt1())
	fmt.Println("Day1 Part2: ", day1pt2())
}

func day1pt1() int {
	return day1.Part1("inputs/day1.txt")
}

func day1pt2() int {
	return day1.Part2("inputs/day1.txt")
}

func day7pt1() string {
	return day7.Part1()
}

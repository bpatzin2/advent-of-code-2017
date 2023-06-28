package day1

import (
	"os"
	"strings"
)

func Part1(filename string) int {
	slice := intSliceFromFile(filename)
	return sumNumbersThatMatchNext(slice)
}

func Part2(filename string) int {
	slice := intSliceFromFile(filename)
	return sumNumbersThatMatchHalfway(slice)
}

func sumNumbersThatMatchNext(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == slice[(i+1)%len(slice)] {
			sum += slice[i]
		}
	}
	return sum
}

func sumNumbersThatMatchHalfway(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == slice[(i+len(slice)/2)%len(slice)] {
			sum += slice[i]
		}
	}
	return sum
}

func intSliceFromFile(filename string) []int {
	dat, err := os.ReadFile(filename)
	check(err)
	s := strings.TrimSpace(string(dat))
	return strToIntSlice(s)
}

func strToIntSlice(s string) []int {
	slice := make([]int, len(s))
	for i, r := range s {
		slice[i] = runeToInt(r)
	}
	return slice
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

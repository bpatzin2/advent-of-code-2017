package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	slice := intSliceFromFile("../inputs/day1.txt")
	sum := sumNumbersThatMatchNext(slice)
	fmt.Println(sum)
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

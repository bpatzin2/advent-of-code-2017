package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	rows := intSlicesFromFile("../inputs/day2.txt")
	c := checksum(rows)
	fmt.Println("Part1: ", c)
}

func checksum(rows [][]int) int {
	return lo.SumBy(rows, rowChecksum)
}

func rowChecksum(row []int) int {
	return lo.Max(row) - lo.Min(row)
}

func intSlicesFromFile(filename string) [][]int {
	lines := readLinesFromFile(filename)
	slices := make([][]int, len(lines))
	for i, line := range lines {
		slices[i] = strToIntSlice(line)
	}
	return slices
}

func readLinesFromFile(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	check(scanner.Err())
	return lines
}

func strToIntSlice(s string) []int {
	intsAsStrs := strings.Fields(s)
	slice := make([]int, len(intsAsStrs))
	for i, intAsStr := range intsAsStrs {
		slice[i], _ = strconv.Atoi(intAsStr)
	}
	return slice
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

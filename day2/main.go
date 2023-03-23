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
	dvq := divibleValueQuotent(rows)
	fmt.Println("Part2: ", dvq)
}

func checksum(rows [][]int) int {
	return lo.SumBy(rows, rowChecksum)
}

func divibleValueQuotent(rows [][]int) int {
	return lo.SumBy(rows, rowDivibleValueQuotent)
}

func rowDivibleValueQuotent(row []int) int {
	for i, a := range row {
		for _, b := range row[i+1:] {
			if a%b == 0 {
				return a / b
			}
			if b%a == 0 {
				return b / a
			}
		}
	}
	return 0
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, err := intSliceFromFile("../inputs/day5.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Part1:", executeSteps(numbers))
}

func executeSteps(ints []int) int {
	length := len(ints)
	stepCount := 0
	for instrPtr := 0; instrPtr >= 0 && instrPtr < length; {
		jumpLen := ints[instrPtr]
		ints[instrPtr] += 1
		instrPtr += jumpLen
		stepCount++
	}
	return stepCount
}

func intSliceFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

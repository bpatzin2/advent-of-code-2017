package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

func main() {
	numbers, err := intSliceFromFile("../inputs/day6.txt")
	if err != nil {
		log.Fatalf("intSliceFromFile: %s", err)
	}

	fmt.Println("Part1 test:", distribute([]int{0, 2, 7, 0}))
	fmt.Println("Part1:", distribute(numbers))
}

func distribute(memoryBanks []int) int {
	stepCount := 0
	bankHistory := mapset.NewSet[string]()
	bankHistory.Add(fmt.Sprint(memoryBanks))
	for {
		memoryBanks = distributeOnce(memoryBanks)
		stepCount++
		if bankHistory.Contains(fmt.Sprint(memoryBanks)) {
			return stepCount
		}
		bankHistory.Add(fmt.Sprint(memoryBanks))
	}
}

func distributeOnce(memoryBanks []int) []int {
	numBanks := len(memoryBanks)
	newMemory := make([]int, numBanks)
	copy(newMemory, memoryBanks)
	max := lo.Max(newMemory)
	idx := lo.IndexOf(newMemory, max)
	toDistribute := newMemory[idx]
	newMemory[idx] = 0
	for ; toDistribute > 0; toDistribute-- {
		idx = (idx + 1) % numBanks
		newMemory[idx] = newMemory[idx] + 1
	}
	return newMemory
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
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			var num int
			_, err := fmt.Sscanf(word, "%d", &num)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

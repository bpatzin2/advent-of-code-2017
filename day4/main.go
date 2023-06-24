package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func main() {
	lines, err := strSlicesFromFile("../inputs/day4.txt") // Replace "test.txt" with your file path
	if err != nil {
		log.Fatalf("readFile: %s", err)
	}
	fmt.Println("Part 1", countValid(lines))
	fmt.Println("Part 2", countValid2(lines))
}

func countValid(passphrases [][]string) int {
	return lo.CountBy(passphrases, func(passphrase []string) bool {
		return isValid(passphrase)
	})
}

func countValid2(passphrases [][]string) int {
	return lo.CountBy(passphrases, func(passphrase []string) bool {
		return isValid2(passphrase)
	})
}

func isValid(passphrase []string) bool {
	dups := lo.FindDuplicates(passphrase)
	return len(dups) == 0
}

func isValid2(passphrase []string) bool {
	letterCounts := lo.Map(passphrase, func(word string, index int) string {
		chars := []byte(word)
		countMap := lo.CountValues(chars)
		return mapToSortedJSONString(countMap)
	})
	dups := lo.FindDuplicates(letterCounts)
	return len(dups) == 0
}

func mapToSortedJSONString(m map[byte]int) string {
	// Get the keys and sort them to ensure a consistent order
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	// Build a new map with the sorted keys
	sortedMap := make(map[byte]int)
	for _, k := range keys {
		sortedMap[byte(k)] = m[byte(k)]
	}

	// Convert the sorted map to JSON
	jsonString, err := json.Marshal(sortedMap)
	if err != nil {
		// Handle the error in a way that's appropriate for your application
		fmt.Println(err)
		return ""
	}

	return string(jsonString)
}

func strSlicesFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		lines = append(lines, words)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samber/lo"
)

func main() {
	lines, err := strSlicesFromFile("../inputs/day4.txt") // Replace "test.txt" with your file path
	if err != nil {
		log.Fatalf("readFile: %s", err)
	}
	fmt.Println("Part 1", countValid(lines))
}

func countValid(passphrases [][]string) int {
	return lo.CountBy(passphrases, func(passphrase []string) bool {
		return isValid(passphrase)
	})
}

func isValid(passphrase []string) bool {
	dups := lo.FindDuplicates(passphrase)
	return len(dups) == 0
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

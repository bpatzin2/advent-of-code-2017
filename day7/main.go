package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	programs, err := parseInput("../inputs/day7.txt")
	if err != nil {
		log.Fatalf("intSliceFromFile: %s", err)
	}

	result, ok := findRoot(programs)
	fmt.Println("Part1:", result, ok)
}

type Program struct {
	name     string
	weight   int
	children []string
}

func findRoot(programs []Program) (Program, bool) {
	parents := indexByParent(programs)
	return lo.Find(programs, func(program Program) bool {
		_, ok := parents[program.name]
		if ok {
			return false
		}
		return len(program.children) > 0
	})
}

func indexByParent(programs []Program) map[string]string {
	parents := make(map[string]string)
	for _, program := range programs {
		for _, child := range program.children {
			parents[child] = program.name
		}
	}
	return parents
}

func parseInput(filename string) ([]Program, error) {
	lines, error := readLines(filename)
	if error != nil {
		return nil, error
	}
	return lo.Map(lines, func(line string, index int) Program {
		return parseLine(line)
	}), nil
}

func parseLine(line string) Program {
	parts := strings.Split(line, "->")

	nameAndWeight := strings.Fields(parts[0])
	name := nameAndWeight[0]
	weightStr := strings.Trim(nameAndWeight[1], "()")
	weight, err := strconv.Atoi(weightStr)
	if err != nil {
		return Program{}
	}

	var children []string
	if len(parts) > 1 {
		childrenParts := strings.Split(parts[1], ",")
		for _, child := range childrenParts {
			children = append(children, strings.TrimSpace(child))
		}
	}

	return Program{
		name:     name,
		weight:   weight,
		children: children,
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

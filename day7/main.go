package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Part1() string {
	programInfos, err := parseInput("inputs/day7.txt")
	if err != nil {
		log.Fatalf("intSliceFromFile: %s", err)
	}

	root := makeTree(programInfos)
	return root.name
}

func Part2() int {
	programInfos, err := parseInput("inputs/day7.txt")
	if err != nil {
		log.Fatalf("intSliceFromFile: %s", err)
	}

	root := makeTree(programInfos)
	rootWithTotalWeight := calculateTotalWeight(root)
	return calculateExpectedWeight(rootWithTotalWeight, rootWithTotalWeight.totalWeight)
}

func calculateExpectedWeight(program *ProgramWithTotalWeight, totalExpectedWeight int) int {
	if program.isBalanced {
		panic("program is already balanced")
	}

	balancedChildCount := lo.CountBy(program.children, func(child *ProgramWithTotalWeight) bool {
		return child.isBalanced
	})

	if balancedChildCount == len(program.children) {
		if len(program.children) >= 3 {
			weights := lo.GroupBy(program.children, func(child *ProgramWithTotalWeight) int {
				return child.totalWeight
			})
			weightGroups := lo.Values(weights)
			modalWeightGroup := lo.MaxBy(weightGroups, func(a []*ProgramWithTotalWeight, b []*ProgramWithTotalWeight) bool {
				return len(a) > len(b)
			})
			modalWeightGroupSize := len(modalWeightGroup)

			modalWeight := 0
			for weight, group := range weights {
				if len(group) == modalWeightGroupSize {
					modalWeight = weight
					break
				}
			}

			unbalancedChild, _ := lo.Find(program.children, func(child *ProgramWithTotalWeight) bool {
				return child.totalWeight != modalWeight
			})

			unbalancedDiff := modalWeight - unbalancedChild.totalWeight
			return unbalancedChild.weight + unbalancedDiff
		} else {
			panic("program is not balanced and has less than 3 children")
			// todo - handle this case
			// two children with different total weights
			// you can take the expected weight of the subtree
			// subtract the wight of the parent node
			// that gives you the expected weight of the combined children
			// but then it's unclear which child to adjust?
		}
	}

	expectedTotalOfUnbalancedChild := totalExpectedWeight - program.weight
	for _, child := range program.children {
		if child.isBalanced {
			expectedTotalOfUnbalancedChild -= child.totalWeight
		}
	}
	unbalancedChild, _ := lo.Find(program.children, func(child *ProgramWithTotalWeight) bool {
		return !child.isBalanced
	})

	return calculateExpectedWeight(unbalancedChild, expectedTotalOfUnbalancedChild)
}

func calculateTotalWeight(program *Program) *ProgramWithTotalWeight {
	weight := program.weight
	children := make([]*ProgramWithTotalWeight, 0, len(program.children)) // make it with length 0 but capacity equal to len(program.children)
	for _, child := range program.children {
		childWithTotalWeight := calculateTotalWeight(child)
		children = append(children, childWithTotalWeight)
		weight += childWithTotalWeight.totalWeight
	}
	return &ProgramWithTotalWeight{
		name:        program.name,
		weight:      program.weight,
		children:    children,
		totalWeight: weight,
		isBalanced:  isBalanced(program, children),
	}
}

func isBalanced(program *Program, children []*ProgramWithTotalWeight) bool {
	if len(children) == 0 {
		return true
	}

	weight := children[0].totalWeight
	for _, child := range children {
		if child.totalWeight != weight {
			return false
		}
	}
	return true
}

func makeTree(programInfos []ProgramInfo) *Program {
	programs := make([]*Program, len(programInfos))
	for i, programInfo := range programInfos {
		programs[i] = &Program{
			name:     programInfo.name,
			weight:   programInfo.weight,
			children: []*Program{},
		}
	}

	programsByName := make(map[string]*Program, len(programInfos))
	for _, program := range programs {
		programsByName[program.name] = program
	}

	for _, programInfo := range programInfos {
		program := programsByName[programInfo.name]
		for _, childName := range programInfo.children {
			child := programsByName[childName]
			program.children = append(program.children, child)
		}
	}

	result, _ := findRoot(programs)
	return result
}

type ProgramInfo struct {
	name     string
	weight   int
	children []string
}

type Program struct {
	name     string
	weight   int
	children []*Program
}

type ProgramWithTotalWeight struct {
	name        string
	weight      int
	children    []*ProgramWithTotalWeight
	totalWeight int
	isBalanced  bool
}

func findRoot(programs []*Program) (*Program, bool) {
	parents := indexByParent(programs)
	return lo.Find(programs, func(program *Program) bool {
		_, ok := parents[program.name]
		if ok {
			return false
		}
		return len(program.children) > 0
	})
}

func indexByParent(programs []*Program) map[string]string {
	parents := make(map[string]string)
	for _, program := range programs {
		for _, child := range program.children {
			parents[child.name] = program.name
		}
	}
	return parents
}

func parseInput(filename string) ([]ProgramInfo, error) {
	lines, error := readLines(filename)
	if error != nil {
		return nil, error
	}
	return lo.Map(lines, func(line string, index int) ProgramInfo {
		return parseLine(line)
	}), nil
}

func parseLine(line string) ProgramInfo {
	parts := strings.Split(line, "->")

	nameAndWeight := strings.Fields(parts[0])
	name := nameAndWeight[0]
	weightStr := strings.Trim(nameAndWeight[1], "()")
	weight, err := strconv.Atoi(weightStr)
	if err != nil {
		return ProgramInfo{}
	}

	var children []string
	if len(parts) > 1 {
		childrenParts := strings.Split(parts[1], ",")
		for _, child := range childrenParts {
			children = append(children, strings.TrimSpace(child))
		}
	}

	return ProgramInfo{
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

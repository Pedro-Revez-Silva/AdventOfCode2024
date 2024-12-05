package day05

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rules struct {
	x, y int
}

func Run() {
	fmt.Println("\nDay 5:")
	puzzle1()
}

func puzzle1() {
	file, err := os.Open("input5.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)

	var rules []Rules
	var updates [][]int
	var validUpdates [][]int
	var invalidUpdates [][]int

	rules, updates = parseInput(scanner)
	for _, update := range updates {
		if isValidUpdateSequence(update, rules) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	sum := sumValidUpdates(validUpdates)

	fmt.Println("Answer to first puzzle:", sum)

	// Part two
	var sortedUpdates [][]int
	for _, invalidUpdate := range invalidUpdates {
		sortedUpdate := sortUpdate(invalidUpdate, rules)
		sortedUpdates = append(sortedUpdates, sortedUpdate)
	}

	sumInvalidUpdates := sumValidUpdates(sortedUpdates)
	fmt.Println("Answer to second puzzle:", sumInvalidUpdates)
}

func parseInput(scanner *bufio.Scanner) ([]Rules, [][]int) {
	rules := make([]Rules, 0)
	updates := make([][]int, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "|") {
			numbers := strings.Split(text, "|")
			x, err := strconv.Atoi(strings.TrimSpace(numbers[0]))
			utils.Check(err)
			y, err := strconv.Atoi(strings.TrimSpace(numbers[1]))
			utils.Check(err)

			rules = append(rules, Rules{x: x, y: y})
		} else if strings.Contains(text, ",") {
			numbers := strings.Split(text, ",")
			var convertedNumbers []int
			for _, number := range numbers {
				convertedNumber, err := strconv.Atoi(number)
				utils.Check(err)
				convertedNumbers = append(convertedNumbers, convertedNumber)
			}
			updates = append(updates, convertedNumbers)
		}
	}

	return rules, updates
}

func isValidUpdateSequence(update []int, rules []Rules) bool {
	for _, rule := range rules {
		if containsBothPages(update, rule.x, rule.y) {
			xPosition := slices.Index(update, rule.x)
			yPosition := slices.Index(update, rule.y)

			if xPosition > yPosition {
				return false
			}
		}
	}
	return true
}

func containsBothPages(update []int, x int, y int) bool {
	containsX := slices.Contains(update, x)
	containsY := slices.Contains(update, y)

	if containsX && containsY {
		return true
	} else {
		return false
	}
}

func sumValidUpdates(validUpdates [][]int) int {
	sum := 0
	for _, validUpdate := range validUpdates {
		middleIndex := (len(validUpdate) - 1) / 2
		sum += validUpdate[middleIndex]
	}
	return sum
}

func buildGraph(numbers []int, rules []Rules) map[int][]int {
	graph := make(map[int][]int)

	for _, num := range numbers {
		graph[num] = make([]int, 0)
	}

	for _, rule := range rules {
		if slices.Contains(numbers, rule.x) && slices.Contains(numbers, rule.y) {
			graph[rule.x] = append(graph[rule.x], rule.y)
		}
	}

	return graph
}

func topologicalSort(numbers []int, graph map[int][]int) []int {
	visited := make(map[int]bool)
	temp := make(map[int]bool)
	order := make([]int, 0)

	var visit func(int) bool
	visit = func(n int) bool {
		if temp[n] {
			return false
		}
		if visited[n] {
			return true
		}

		temp[n] = true

		for _, m := range graph[n] {
			if !visit(m) {
				return false
			}
		}

		temp[n] = false
		visited[n] = true
		order = append([]int{n}, order...)
		return true
	}

	for _, n := range numbers {
		if !visited[n] {
			if !visit(n) {
				return nil
			}
		}
	}

	return order
}

func sortUpdate(update []int, rules []Rules) []int {
	graph := buildGraph(update, rules)
	return topologicalSort(update, graph)
}

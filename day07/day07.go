package day07

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Functions struct {
	result  int
	numbers []int
}

func Run() {
	fmt.Println("\nDay 7:")
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("input7.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	functions := parseFunctions(lines)
	sum := 0
	for _, f := range functions {
		if canMakeTarget(f) {
			sum += f.result
		}
	}

	fmt.Printf("Answer to first puzzle: %d\n", sum)
}

func parseFunctions(lines []string) []Functions {
	var functions []Functions
	for _, line := range lines {
		parts := strings.Split(line, ":")

		targetResult, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Printf("Error parsing target result: %s", err)
			continue
		}

		numbersStrs := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, 0, len(numbersStrs))
		for _, numStr := range numbersStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Printf("Error parsing number: %s", err)
				continue
			}
			numbers = append(numbers, num)
		}

		functions = append(functions, Functions{
			result:  targetResult,
			numbers: numbers,
		})
	}

	return functions
}

func evaluateExpression(numbers []int, operators []rune) int {
	result := numbers[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == '+' {
			result += numbers[i+1]
		} else {
			result *= numbers[i+1]
		}
	}
	return result
}

func canMakeTarget(f Functions) bool {
	numOperators := len(f.numbers) - 1
	maxCombinations := 1 << numOperators

	for i := 0; i < maxCombinations; i++ {
		operators := make([]rune, numOperators)
		for j := 0; j < numOperators; j++ {
			if (i & (1 << j)) == 0 {
				operators[j] = '+'
			} else {
				operators[j] = '*'
			}
		}

		if evaluateExpression(f.numbers, operators) == f.result {
			return true
		}
	}
	return false
}

func puzzle2() {
	file, err := os.Open("input7.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	functions := parseFunctions(lines)
	sum := 0
	for _, f := range functions {
		if canMakeTargetV2(f) {
			sum += f.result
		}
	}

	fmt.Printf("Answer to second puzzle: %d\n", sum)
}

func canMakeTargetV2(f Functions) bool {
	numOperators := len(f.numbers) - 1
	maxCombinations := int(math.Pow(3, float64(numOperators)))

	for i := 0; i < maxCombinations; i++ {
		operators := make([]int, numOperators)
		temp := i
		for j := 0; j < numOperators; j++ {
			operators[j] = temp % 3
			temp /= 3
		}

		if evaluateExpressionWithConcat(f.numbers, operators) == f.result {
			return true
		}
	}
	return false
}

func evaluateExpressionWithConcat(numbers []int, operators []int) int {
	result := numbers[0]

	for i := 0; i < len(operators); i++ {
		currentNum := numbers[i+1]
		switch operators[i] {
		case 0: // addition
			result += currentNum
		case 1: // multiplication
			result *= currentNum
		case 2: // concatenation
			resultStr := strconv.Itoa(result) + strconv.Itoa(currentNum)
			var err error
			result, err = strconv.Atoi(resultStr)
			if err != nil {
				return -1
			}
		}
	}
	return result
}

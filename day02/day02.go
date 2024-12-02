package day02

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkIfSafe(line []int) bool {
	var previousDifference int = 0
	for i := 0; i < len(line)-1; i++ {
		difference := line[i] - line[i+1]
		if difference == 0 || intAbs(difference) > 3 {
			return false
		} else if difference < 0 && previousDifference > 0 {
			return false
		} else if difference > 0 && previousDifference < 0 {
			return false
		}
		previousDifference = difference
	}
	return true
}

func Run() {
	fmt.Println("\nDay 2:")
	file, err := os.Open("input2.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	var lines [][]int
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var lineNumbers []int
		for _, field := range fields {
			number, err := strconv.Atoi(field)
			utils.Check(err)
			lineNumbers = append(lineNumbers, number)
		}
		lines = append(lines, lineNumbers)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var safeReports int
	for _, line := range lines {
		if checkIfSafe(line) {
			safeReports++
		} else { // Add "if false" to get the answer to the first puzzle.
			for i := 0; i < len(line); i++ {
				tempLine := make([]int, len(line))
				copy(tempLine, line)
				tempLine = append(tempLine[:i], tempLine[i+1:]...)
				if checkIfSafe(tempLine) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Printf("Answer to puzzle: %d\n", safeReports)

}

package day03

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("\nDay 3:")
	file, err := os.Open("input3.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}
	utils.Check(scanner.Err())
	var sum int
	// Remove this bool and related conditions to get the answer to the first puzzle
	mulEnabled := true
	for i := 0; i < len(text)-7; i++ {
		if mulEnabled {
			if text[i:i+4] == "mul(" {
				start := i + 4
				for j := start; j < start+8 && j < len(text); j++ {
					if text[j] == ')' {
						numberPart := text[start:j]
						parts := strings.Split(numberPart, ",")
						if len(parts) != 2 {
							continue
						}
						firstNum, err1 := strconv.Atoi(parts[0])
						secondNum, err2 := strconv.Atoi(parts[1])
						if err1 != nil || err2 != nil {
							continue
						}
						sum += firstNum * secondNum
						break
					}
				}
			} else if text[i:i+7] == "don't()" {
				mulEnabled = false
				i += 5
			}
		} else {
			if text[i:i+4] == "do()" {
				mulEnabled = true
				i += 2
			}
		}
	}

	fmt.Printf("Answer to second puzzle: %d\n", sum)

}

package day01

import (
	"AdventOfCode2024/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func Run() {
	fmt.Println("Day 1:")
	file, err := os.Open("input.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	var lines []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		lines = append(lines, line)
	}
	var slice1 []string
	var slice2 []string
	for i, line := range lines {
		if i%2 == 0 {
			slice1 = append(slice1, line)
		} else {
			slice2 = append(slice2, line)
		}
	}

	for i := 0; i < len(slice1); i++ {
		for j := i + 1; j < len(slice1); j++ {
			if slice1[i] > slice1[j] {
				slice1[i], slice1[j] = slice1[j], slice1[i]
			}
			if slice2[i] > slice2[j] {
				slice2[i], slice2[j] = slice2[j], slice2[i]
			}
		}
	}

	var sum int64
	for i := 0; i < len(slice1); i++ {
		num1, err := strconv.ParseInt(slice1[i], 10, 64)
		utils.Check(err)
		num2, err := strconv.ParseInt(slice2[i], 10, 64)
		utils.Check(err)
		if num1 > num2 {
			sum += num1 - num2
		} else if num2 > num1 {
			sum += num2 - num1
		}
	}
	fmt.Printf("Answer to the first puzzle: %d\n", sum)

	// Part 2
	var slice3 []int
	for i := 0; i < len(slice1); i++ {
		count := 0
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				count++
			}
		}
		slice3 = append(slice3, count)
	}

	var sum2 int64
	for i := 0; i < len(slice1); i++ {
		num1, err := strconv.ParseInt(slice1[i], 10, 64)
		utils.Check(err)
		sum2 += num1 * int64(slice3[i])
	}
	fmt.Printf("Answer to the second puzzle: %d\n", sum2)
}

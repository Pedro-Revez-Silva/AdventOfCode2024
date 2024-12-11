package day11

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("\nDay 11:")
	puzzle()
}

func processNumber(n int64) []int64 {
	if n == 0 {
		return []int64{1}
	}

	digits := int(math.Floor(math.Log10(float64(n)))) + 1

	if digits%2 == 0 {
		divisor := int64(math.Pow(10, float64(digits/2)))
		left := n / divisor
		right := n % divisor
		return []int64{left, right}
	}

	return []int64{n * 2024}
}

func blink(stones map[int64]int64) map[int64]int64 {
	result := make(map[int64]int64)

	for stone, count := range stones {
		newStones := processNumber(stone)
		for _, newStone := range newStones {
			result[newStone] += count
		}
	}

	return result
}

func countStones(stones map[int64]int64) int64 {
	var total int64
	for _, count := range stones {
		total += count
	}
	return total
}

func puzzle() {
	file, err := os.Open("input11.txt")
	utils.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)
	stones := make(map[int64]int64)

	if scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		for _, numStr := range numbers {
			num, _ := strconv.ParseInt(numStr, 10, 64)
			stones[num]++
		}
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println("Answer to puzzle 1:", countStones(stones))

	for i := 0; i < 50; i++ {
		stones = blink(stones)
	}

	fmt.Println("Answer to puzzle 2:", countStones(stones))
}

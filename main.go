package main

import (
	"AdventOfCode2024/day11"
	"fmt"
	"time"
)

func timeFunction(name string, f func()) {
	start := time.Now()
	f()
	duration := time.Since(start)

	// Format duration based on magnitude
	var timeStr string
	if duration.Milliseconds() < 1 {
		timeStr = fmt.Sprintf("%d µs", duration.Microseconds())
	} else if duration.Seconds() < 1 {
		timeStr = fmt.Sprintf("%d ms", duration.Milliseconds())
	} else {
		timeStr = fmt.Sprintf("%.2f s", duration.Seconds())
	}

	fmt.Printf("Time taken for %s: %s\n", name, timeStr)
}

func main() {
	fmt.Println("Advent of Code 2024")
	fmt.Println("===================")

	//timeFunction("Day 1", day01.Run)
	//timeFunction("Day 2", day02.Run)
	//timeFunction("Day 3", day03.Run)
	//timeFunction("Day 4", day04.Run)
	//timeFunction("Day 5", day05.Run)
	//timeFunction("Day 6", day06.Run)
	//timeFunction("Day 7", day07.Run)
	//timeFunction("Day 8", day08.Run)
	//timeFunction("Day 9", day09.Run)
	//timeFunction("Day 10", day10.Run)
	timeFunction("Day 11", day11.Run)
}

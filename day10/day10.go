package day10

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	fmt.Println("\nDay 10:")
	grid := utils.ReadGrid("input10.txt")
	puzzle1(grid)
	puzzle2(grid)
}

func trailsCounter(grid []string, pos utils.Position, val int, reached map[utils.Position]bool) int {
	if pos.Y < 0 || pos.X < 0 || pos.Y >= len(grid) || pos.X >= len(grid[0]) {
		return 0
	}
	if int(grid[pos.Y][pos.X]-'0') != val {
		return 0
	}
	if val == 9 {
		if reached[pos] {
			return 0
		}
		reached[pos] = true
		return 1
	}
	return trailsCounter(grid, utils.Position{X: pos.X - 1, Y: pos.Y}, val+1, reached) +
		trailsCounter(grid, utils.Position{X: pos.X + 1, Y: pos.Y}, val+1, reached) +
		trailsCounter(grid, utils.Position{X: pos.X, Y: pos.Y - 1}, val+1, reached) +
		trailsCounter(grid, utils.Position{X: pos.X, Y: pos.Y + 1}, val+1, reached)
}

func puzzle1(grid []string) {
	totalScore := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '0' {
				reached := make(map[utils.Position]bool)
				totalScore += trailsCounter(grid, utils.Position{X: x, Y: y}, 0, reached)
			}
		}
	}

	fmt.Println("Answer to puzzle 1:", totalScore)
}

func pathCounter(grid []string, pos utils.Position, val int, visited map[utils.Position]bool) int {
	if pos.Y < 0 || pos.X < 0 || pos.Y >= len(grid) || pos.X >= len(grid[0]) {
		return 0
	}

	if int(grid[pos.Y][pos.X]-'0') != val {
		return 0
	}

	if visited[pos] {
		return 0
	}

	if val == 9 {
		return 1
	}

	visited[pos] = true

	paths := pathCounter(grid, utils.Position{X: pos.X - 1, Y: pos.Y}, val+1, visited) +
		pathCounter(grid, utils.Position{X: pos.X + 1, Y: pos.Y}, val+1, visited) +
		pathCounter(grid, utils.Position{X: pos.X, Y: pos.Y - 1}, val+1, visited) +
		pathCounter(grid, utils.Position{X: pos.X, Y: pos.Y + 1}, val+1, visited)

	visited[pos] = false

	return paths
}
func puzzle2(grid []string) {
	totalRating := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '0' {
				visited := make(map[utils.Position]bool)
				rating := pathCounter(grid, utils.Position{X: x, Y: y}, 0, visited)
				totalRating += rating
			}
		}
	}

	fmt.Println("Answer to puzzle 2:", totalRating)
}

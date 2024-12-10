package day10

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strings"
)

var directions = []utils.Position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func Run() {
	fmt.Println("\nDay 10:")
	grid := utils.ReadGrid("input10.txt")
	puzzle1(grid)
	puzzle2(grid)
}

func findTrailheads(grid []string) []utils.Position {
	var trailheads []utils.Position
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '0' {
				trailheads = append(trailheads, utils.Position{X: x, Y: y})
			}
		}
	}
	return trailheads
}

func getNeighbors(p utils.Position, grid []string) []utils.Position {
	var neighbors []utils.Position
	for _, dir := range directions {
		newX, newY := p.X+dir.X, p.Y+dir.Y
		if newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) {
			neighbors = append(neighbors, utils.Position{X: newX, Y: newY})
		}
	}
	return neighbors
}

func getHeight(p utils.Position, grid []string) int {
	return int(grid[p.Y][p.X] - '0')
}

func dfs(pos utils.Position, currentHeight int, path map[utils.Position]bool, peaks map[utils.Position]bool, grid []string) {
	if getHeight(pos, grid) == 9 {
		peaks[pos] = true
		return
	}

	for _, next := range getNeighbors(pos, grid) {
		nextHeight := getHeight(next, grid)
		if nextHeight == currentHeight+1 && !path[next] {
			newPath := make(map[utils.Position]bool)
			for k, v := range path {
				newPath[k] = v
			}
			newPath[next] = true
			dfs(next, nextHeight, newPath, peaks, grid)
		}
	}
}

func findReachablePeaks(start utils.Position, grid []string) map[utils.Position]bool {
	peaks := make(map[utils.Position]bool)
	path := map[utils.Position]bool{start: true}
	dfs(start, 0, path, peaks, grid)
	return peaks
}

func puzzle1(grid []string) {
	trailheads := findTrailheads(grid)
	totalScore := 0

	for _, trailhead := range trailheads {
		peaks := findReachablePeaks(trailhead, grid)
		totalScore += len(peaks)
	}

	fmt.Println("Answer to puzzle 1:", totalScore)
}

func getTrailString(path []utils.Position) string {
	var parts []string
	for _, p := range path {
		parts = append(parts, fmt.Sprintf("%d,%d", p.X, p.Y))
	}
	return strings.Join(parts, "|")
}

func dfsTrails(pos utils.Position, currentHeight int, path []utils.Position, trails map[string]bool, grid []string) {
	if getHeight(pos, grid) == 9 {
		trail := getTrailString(path)
		trails[trail] = true
		return
	}

	for _, next := range getNeighbors(pos, grid) {
		nextHeight := getHeight(next, grid)
		if nextHeight == currentHeight+1 {
			alreadyVisited := false
			for _, p := range path {
				if p == next {
					alreadyVisited = true
					break
				}
			}
			if !alreadyVisited {
				newPath := make([]utils.Position, len(path))
				copy(newPath, path)
				newPath = append(newPath, next)
				dfsTrails(next, nextHeight, newPath, trails, grid)
			}
		}
	}
}

func countUniqueTrails(start utils.Position, grid []string) int {
	trails := make(map[string]bool)
	initialPath := []utils.Position{start}
	dfsTrails(start, 0, initialPath, trails, grid)
	return len(trails)
}

func puzzle2(grid []string) {
	trailheads := findTrailheads(grid)
	totalRating := 0

	for _, trailhead := range trailheads {
		rating := countUniqueTrails(trailhead, grid)
		totalRating += rating
	}

	fmt.Println("Answer to puzzle 2:", totalRating)
}

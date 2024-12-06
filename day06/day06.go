package day06

import (
	"AdventOfCode2024/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type Position struct {
	x, y int
}

func Run() {
	fmt.Println("\nDay 6:")
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("input6.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	positionsCounter := 1
	var grid []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		grid = append(grid, line)
	}
	var position Position

	direction := Position{-1, 0}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				position = Position{i, j}
				break
			}
		}
	}

	visitedPositions := make(map[Position]bool)
	visitedPositions[position] = true
loopLabel:
	for isInbounds(position, grid) {
		moveResult := canMove(position, direction, grid)
		switch moveResult {
		case "yes":
			position = Position{position.x + direction.x, position.y + direction.y}
			if !visitedPositions[position] {
				positionsCounter++
				visitedPositions[position] = true
			}
		case "obstacle":
			direction = rotateRight(direction)
		case "out":
			break loopLabel
		}

	}

	fmt.Println("Answer to first puzzle:", positionsCounter)
}

func isInbounds(position Position, grid []string) bool {
	if position.x < 0 || position.x >= len(grid) {
		return false
	}
	if position.y < 0 || position.y >= len(grid[0]) {
		return false
	}
	return true
}

func canMove(position Position, direction Position, grid []string) string {
	newPosition := Position{position.x + direction.x, position.y + direction.y}
	if !isInbounds(newPosition, grid) {
		return "out"
	}
	if grid[newPosition.x][newPosition.y] == '#' {
		return "obstacle"
	}
	return "yes"
}

func rotateRight(direction Position) Position {
	return Position{
		x: direction.y,
		y: -direction.x,
	}
}

func puzzle2() {
	file, err := os.Open("input6.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	var grid []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		grid = append(grid, line)
	}
	var startPosition Position

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				startPosition = Position{i, j}
				break
			}
		}
	}

	validPositions := 0
	startDirection := Position{-1, 0}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' || (i == startPosition.x && j == startPosition.y) {
				continue
			}
			newGrid := make([]string, len(grid))
			copy(newGrid, grid)
			row := []rune(newGrid[i])
			row[j] = '#'
			newGrid[i] = string(row)

			if isLoop(newGrid, startPosition, startDirection) {
				validPositions++
			}
		}
	}

	fmt.Println("Answer to second puzzle:", validPositions)

}

func isLoop(grid []string, startPos Position, startDir Position) bool {
	position := startPos
	direction := startDir

	type state struct {
		pos Position
		dir Position
	}
	visited := make(map[state]bool)

	for {
		currentState := state{position, direction}
		if visited[currentState] {
			return true
		}
		visited[currentState] = true

		moveResult := canMove(position, direction, grid)
		switch moveResult {
		case "yes":
			position = Position{position.x + direction.x, position.y + direction.y}
		case "obstacle":
			direction = rotateRight(direction)
		case "out":
			return false
		}
	}
}

package day08

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"os"
)

type AntennaMap struct {
	Frequencies   map[rune][]utils.Position
	Width, Height int
}

func Run() {
	fmt.Println("\nDay 8:")
	grid := readGrid("input8.txt")
	antennaMap := parseAntennaMap(grid)

	antinodes := antennaMap.findAntinodes()
	fmt.Println("Answer to first puzzle:", len(antinodes))

	resonantAntinodes := antennaMap.findResonantAntinodes()
	fmt.Println("Answer to second puzzle:", len(resonantAntinodes))
}

func readGrid(filename string) []string {
	file, err := os.Open(filename)
	utils.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	utils.Check(scanner.Err())
	return grid
}

func parseAntennaMap(grid []string) *AntennaMap {
	aMap := &AntennaMap{
		Frequencies: make(map[rune][]utils.Position),
		Height:      len(grid),
		Width:       len(grid[0]),
	}

	for y, row := range grid {
		for x, char := range row {
			if char != '.' {
				aMap.Frequencies[char] = append(aMap.Frequencies[char], utils.Position{X: x, Y: y})
			}
		}
	}
	return aMap
}

func (a *AntennaMap) findAntinodes() map[utils.Position]bool {
	antinodes := make(map[utils.Position]bool)

	for _, positions := range a.Frequencies {
		if len(positions) < 2 {
			continue
		}

		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				a.calculatePairAntinodes(positions[i], positions[j], antinodes)
			}
		}
	}
	return antinodes
}

func (a *AntennaMap) calculatePairAntinodes(pos1, pos2 utils.Position, antinodes map[utils.Position]bool) {
	vector := utils.Position{
		X: pos2.X - pos1.X,
		Y: pos2.Y - pos1.Y,
	}

	antinode1 := utils.Position{X: pos1.X - vector.X, Y: pos1.Y - vector.Y}
	antinode2 := utils.Position{X: pos2.X + vector.X, Y: pos2.Y + vector.Y}

	if a.isValidPosition(antinode1) {
		antinodes[antinode1] = true
	}
	if a.isValidPosition(antinode2) {
		antinodes[antinode2] = true
	}
}

func (a *AntennaMap) isValidPosition(p utils.Position) bool {
	return p.X >= 0 && p.X < a.Width && p.Y >= 0 && p.Y < a.Height
}

func (a *AntennaMap) findResonantAntinodes() map[utils.Position]bool {
	antinodes := make(map[utils.Position]bool)

	for _, positions := range a.Frequencies {
		if len(positions) < 2 {
			continue
		}

		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				a.findLinearAntinodes(positions[i], positions[j], antinodes)
			}
		}
	}
	return antinodes
}

func (a *AntennaMap) findLinearAntinodes(pos1, pos2 utils.Position, antinodes map[utils.Position]bool) {
	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			p := utils.Position{X: x, Y: y}
			if isPointOnLine(p, pos1, pos2) {
				antinodes[p] = true
			}
		}
	}
}

func isPointOnLine(p, p1, p2 utils.Position) bool {
	if p1.X == p2.X {
		return p.X == p1.X
	}
	if p1.Y == p2.Y {
		return p.Y == p1.Y
	}

	dx1 := p.X - p1.X
	dy1 := p.Y - p1.Y
	dx2 := p2.X - p1.X
	dy2 := p2.Y - p1.Y

	return dx1*dy2 == dx2*dy1
}

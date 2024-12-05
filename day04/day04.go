package day04

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	dx, dy int
}

func Run() {
	fmt.Println("\nDay 4:")

	puzzle1()
	puzzle2()
}

func puzzle1() {
	wordToSearch := "XMAS"
	file, err := os.Open("input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
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

	directions := []Point{
		{-1, 0},  // North
		{1, 0},   // South
		{0, 1},   // East
		{0, -1},  // West
		{-1, 1},  // North-East
		{-1, -1}, // North-West
		{1, 1},   // South-East
		{1, -1},  // South-West
	}

	wordCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == wordToSearch[0] { // If we find 'X'
				for _, dir := range directions {
					if checkWord(grid, i, j, wordToSearch, dir) {
						wordCount++
					}
				}
			}
		}
	}
	fmt.Println("Answer to first puzzle:", wordCount)
}

func puzzle2() {
	file, err := os.Open("input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
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

	directions := []Point{
		{-1, -1}, // North-West
		{-1, 1},  // North-East
		{1, -1},  // South-West
		{1, 1},   // South-East
	}

	xCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'A' {
				if checkXPattern(grid, i, j, directions) {
					xCount++
				}
			}
		}
	}
	fmt.Println("Answer to second puzzle:", xCount)
}

func checkWord(grid []string, x, y int, word string, dir Point) bool {
	lastX, lastY := x+(dir.dx*(len(word)-1)), y+(dir.dy*(len(word)-1))
	if lastX < 0 || lastY < 0 || lastX >= len(grid) || lastY >= len(grid[0]) {
		return false
	}

	for i := 0; i < len(word); i++ {
		newX, newY := x+(dir.dx*i), y+(dir.dy*i)
		if grid[newX][newY] != word[i] {
			return false
		}
	}
	return true
}

func checkXPattern(grid []string, x, y int, dirs []Point) bool {
	if x-1 < 0 || x+1 >= len(grid) || y-1 < 0 || y+1 >= len(grid[0]) {
		return false
	}

	nw := grid[x+dirs[0].dx][y+dirs[0].dy]
	ne := grid[x+dirs[1].dx][y+dirs[1].dy]
	sw := grid[x+dirs[2].dx][y+dirs[2].dy]
	se := grid[x+dirs[3].dx][y+dirs[3].dy]

	pattern1 := nw == 'M' && ne == 'M' && sw == 'S' && se == 'S'
	pattern2 := nw == 'S' && ne == 'S' && sw == 'M' && se == 'M'
	pattern3 := nw == 'M' && se == 'S' && ne == 'S' && sw == 'M'
	pattern4 := nw == 'S' && se == 'M' && ne == 'M' && sw == 'S'

	return pattern1 || pattern2 || pattern3 || pattern4
}

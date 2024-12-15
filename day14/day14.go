package day14

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Width  = 101
	Height = 103
	Steps  = 100
)

type robot struct {
	x, y   int
	dx, dy int
}

func (r *robot) simulate(steps int) {
	for i := 0; i < steps; i++ {
		r.x = (r.x + r.dx) % Width
		if r.x < 0 {
			r.x += Width
		}
		r.y = (r.y + r.dy) % Height
		if r.y < 0 {
			r.y += Height
		}
	}
}

func Run() {
	fmt.Println("\nDay 14:")
	file, err := os.Open("input14.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	scanner := bufio.NewScanner(file)

	var robots []robot

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		r, err := parseLine(line)
		if err != nil {
			log.Fatalf("Failed to parse line %q: %v", line, err)
		}
		robots = append(robots, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	robotsForPuzzle1 := make([]robot, len(robots))
	copy(robotsForPuzzle1, robots)

	puzzle1(robotsForPuzzle1)

	var robotData [][]int
	for _, r := range robots {
		robotData = append(robotData, []int{r.x, r.y, r.dx, r.dy})
	}
	puzzle2(robotData)
}

func puzzle1(robots []robot) {
	for i := range robots {
		robots[i].simulate(Steps)
	}

	topLeftCount, topRightCount, bottomLeftCount, bottomRightCount := countQuadrants(robots)

	safetyFactor := topLeftCount * topRightCount * bottomLeftCount * bottomRightCount

	fmt.Println("Answer to first puzzle: ", safetyFactor)
}

func parseLine(line string) (robot, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return robot{}, fmt.Errorf("invalid line format")
	}

	// Parse position
	posPart := strings.TrimPrefix(parts[0], "p=")
	var px, py int
	_, err := fmt.Sscanf(posPart, "%d,%d", &px, &py)
	if err != nil {
		return robot{}, fmt.Errorf("failed to parse position: %w", err)
	}

	// Parse velocity
	velPart := strings.TrimPrefix(parts[1], "v=")
	var dx, dy int
	_, err = fmt.Sscanf(velPart, "%d,%d", &dx, &dy)
	if err != nil {
		return robot{}, fmt.Errorf("failed to parse velocity: %w", err)
	}

	return robot{x: px, y: py, dx: dx, dy: dy}, nil
}

func countQuadrants(robots []robot) (topLeft, topRight, bottomLeft, bottomRight int) {
	for _, r := range robots {
		if r.x == 50 || r.y == 51 {
			continue
		}

		if r.x < 50 && r.y < 51 {
			topLeft++
		} else if r.x > 50 && r.y < 51 {
			topRight++
		} else if r.x < 50 && r.y > 51 {
			bottomLeft++
		} else if r.x > 50 && r.y > 51 {
			bottomRight++
		}
	}
	return
}

func puzzle2(robots [][]int) {

	i := 1
	for {
		filledIn := make(map[int]map[int]bool)
		for idx := range robots {
			px, py, dx, dy := robots[idx][0], robots[idx][1], robots[idx][2], robots[idx][3]

			px = (px + dx) % Width
			if px < 0 {
				px += Width
			}
			py = (py + dy) % Height
			if py < 0 {
				py += Height
			}

			robots[idx][0] = px
			robots[idx][1] = py

			if filledIn[py] == nil {
				filledIn[py] = make(map[int]bool)
			}
			filledIn[py][px] = true
		}

		found := false
		for y := 0; y < Height && !found; y++ {
			filledRow := filledIn[y]
			if filledRow == nil {
				continue
			}

			for x := 0; x <= Width-10 && !found; x++ {
				allFilled := true
				for offset := 0; offset < 10; offset++ {
					if !filledRow[x+offset] {
						allFilled = false
						break
					}
				}
				if allFilled {
					fmt.Println("Answer for second puzzle: ", i)
					found = true
				}
			}
		}

		if found {
			break
		}

		i++
	}
}

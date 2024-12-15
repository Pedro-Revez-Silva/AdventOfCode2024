package day15

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Cell struct {
	wall  bool
	robot bool
	box   bool
	lhs   bool
	rhs   bool
}

type Grid struct {
	cells [][]Cell
	w, h  int
}

type Vec2 struct {
	x, y int
}

var directions = map[rune]Vec2{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func Run() {
	fmt.Println("\nDay 15:")
	warehouseLines, moveLines := readInput()
	warehouse, robotPos := parseGrid(warehouseLines, false)
	moves := parseMoves(moveLines)

	// Puzzle 1
	p1Grid := copyGrid(warehouse)
	p1Pos := robotPos
	puzzle1(p1Grid, &p1Pos, moves)
	fmt.Printf("Answer to first puzzle: %d\n", calculateGPS(p1Grid, false))

	// Puzzle 2
	p2Grid, p2Pos := parseGrid(warehouseLines, true)
	puzzle2(p2Grid, &p2Pos, moves)
	fmt.Printf("Answer to second puzzle: %d\n", calculateGPS(p2Grid, true))
}

func puzzle1(grid *Grid, robotPos *Vec2, moves []Vec2) {
	executeAllMoves(grid, robotPos, moves)
}

func puzzle2(grid *Grid, robotPos *Vec2, moves []Vec2) {
	executeAllMoves(grid, robotPos, moves)
}

func executeAllMoves(grid *Grid, robotPos *Vec2, moves []Vec2) {
	for _, m := range moves {
		if checkMove(grid, *robotPos, m) {
			executeMove(grid, *robotPos, m)
			robotPos.x += m.x
			robotPos.y += m.y
		}
	}
}

func inBounds(grid *Grid, pos Vec2) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.y < grid.h && pos.x < grid.w
}

func checkMove(grid *Grid, pos Vec2, m Vec2) bool {
	if !inBounds(grid, pos) {
		return false
	}
	cell := &grid.cells[pos.y][pos.x]

	if !cell.wall && !cell.box && !cell.robot {
		return true
	}

	if cell.wall {
		return false
	}

	if cell.box {
		return checkMoveBox(grid, pos, m)
	}

	if cell.robot {
		nextPos := Vec2{pos.x + m.x, pos.y + m.y}
		return checkMove(grid, nextPos, m)
	}

	return false
}

func checkMoveBox(grid *Grid, pos Vec2, m Vec2) bool {
	cell := &grid.cells[pos.y][pos.x]
	vertical := m.y != 0

	if vertical && cell.rhs {
		np1 := Vec2{pos.x - 1, pos.y + m.y}
		np2 := Vec2{pos.x, pos.y + m.y}
		return checkMove(grid, np1, m) && checkMove(grid, np2, m)
	} else if vertical && cell.lhs {
		np1 := Vec2{pos.x, pos.y + m.y}
		np2 := Vec2{pos.x + 1, pos.y + m.y}
		return checkMove(grid, np1, m) && checkMove(grid, np2, m)
	} else {
		newPos := Vec2{pos.x + m.x, pos.y + m.y}
		return checkMove(grid, newPos, m)
	}
}

func executeMove(grid *Grid, pos Vec2, m Vec2) {
	if !inBounds(grid, pos) {
		return
	}
	cell := &grid.cells[pos.y][pos.x]

	if !cell.wall && !cell.box && !cell.robot {
		return
	}

	if cell.box {
		executeMoveBox(grid, pos, m)
	} else if cell.robot {
		newPos := Vec2{pos.x + m.x, pos.y + m.y}
		executeMove(grid, newPos, m)
		grid.cells[pos.y][pos.x].robot = false
		grid.cells[newPos.y][newPos.x].robot = true
	}
}

func executeMoveBox(grid *Grid, pos Vec2, m Vec2) {
	cell := &grid.cells[pos.y][pos.x]
	vertical := m.y != 0

	if vertical && cell.rhs {
		np1 := Vec2{pos.x - 1, pos.y + m.y}
		np2 := Vec2{pos.x, pos.y + m.y}
		executeMove(grid, np1, m)
		executeMove(grid, np2, m)

		grid.cells[pos.y][pos.x].box = false
		grid.cells[pos.y][pos.x].rhs = false
		grid.cells[pos.y][pos.x-1].box = false
		grid.cells[pos.y][pos.x-1].lhs = false

		grid.cells[np1.y][np1.x].box = true
		grid.cells[np1.y][np1.x].lhs = true
		grid.cells[np2.y][np2.x].box = true
		grid.cells[np2.y][np2.x].rhs = true

	} else if vertical && cell.lhs {
		np1 := Vec2{pos.x, pos.y + m.y}
		np2 := Vec2{pos.x + 1, pos.y + m.y}
		executeMove(grid, np1, m)
		executeMove(grid, np2, m)

		grid.cells[pos.y][pos.x].box = false
		grid.cells[pos.y][pos.x].lhs = false
		grid.cells[pos.y][pos.x+1].box = false
		grid.cells[pos.y][pos.x+1].rhs = false

		grid.cells[np1.y][np1.x].box = true
		grid.cells[np1.y][np1.x].lhs = true
		grid.cells[np2.y][np2.x].box = true
		grid.cells[np2.y][np2.x].rhs = true

	} else {
		newpos := Vec2{pos.x + m.x, pos.y + m.y}
		executeMove(grid, newpos, m)

		grid.cells[newpos.y][newpos.x].box = true
		grid.cells[newpos.y][newpos.x].rhs = cell.rhs
		grid.cells[newpos.y][newpos.x].lhs = cell.lhs
		cell.box = false
		cell.rhs = false
		cell.lhs = false
	}
}

func calculateGPS(grid *Grid, wide bool) uint32 {
	var score uint32
	for y := 0; y < grid.h; y++ {
		for x := 0; x < grid.w; x++ {
			c := &grid.cells[y][x]
			if wide {
				if c.lhs {
					score += uint32(y)*100 + uint32(x)
				}
			} else {
				if c.box && !c.lhs && !c.rhs {
					score += uint32(y)*100 + uint32(x)
				}
			}
		}
	}
	return score
}

func readInput() (mapLines []string, moveLines []string) {
	file, err := os.Open("input15.txt")
	utils.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		utils.Check(err)
	}(file)

	reader := bufio.NewReader(file)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		if err == io.EOF {
			if line != "" {
				lines = append(lines, line)
			}
			break
		} else if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}

	i := 0
	for i < len(lines) {
		if strings.TrimSpace(lines[i]) == "" {
			i++
			break
		}
		mapLines = append(mapLines, lines[i])
		i++
	}
	for ; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) != "" {
			moveLines = append(moveLines, strings.TrimSpace(lines[i]))
		}
	}

	return mapLines, moveLines
}

func parseGrid(lines []string, wide bool) (*Grid, Vec2) {
	var robotPos Vec2
	robotPos.x, robotPos.y = -1, -1

	h := len(lines)
	var w int
	if h > 0 {
		w = len(lines[0])
	}
	if wide {
		w = w * 2
	}
	grid := &Grid{w: w, h: h, cells: make([][]Cell, h)}
	for y := 0; y < h; y++ {
		grid.cells[y] = make([]Cell, w)
	}

	for y, line := range lines {
		for x, ch := range line {
			if !wide {
				switch ch {
				case '#':
					grid.cells[y][x].wall = true
				case 'O':
					grid.cells[y][x].box = true
				case '@':
					grid.cells[y][x].robot = true
					robotPos = Vec2{x, y}
				case '.':
				}
			} else {
				baseX := x * 2
				switch ch {
				case '#':
					grid.cells[y][baseX].wall = true
					grid.cells[y][baseX+1].wall = true
				case 'O':
					grid.cells[y][baseX].box = true
					grid.cells[y][baseX].lhs = true
					grid.cells[y][baseX+1].box = true
					grid.cells[y][baseX+1].rhs = true
				case '@':
					grid.cells[y][baseX].robot = true
					robotPos = Vec2{baseX, y}
				case '.':

				}
			}
		}
	}

	return grid, robotPos
}

func parseMoves(moveLines []string) []Vec2 {
	var sb strings.Builder
	for _, l := range moveLines {
		sb.WriteString(l)
	}
	movesStr := sb.String()
	moves := make([]Vec2, 0, len(movesStr))
	for _, ch := range movesStr {
		moves = append(moves, directions[ch])
	}
	return moves
}

func copyGrid(src *Grid) *Grid {
	g := &Grid{w: src.w, h: src.h, cells: make([][]Cell, src.h)}
	for i := 0; i < src.h; i++ {
		g.cells[i] = make([]Cell, src.w)
		copy(g.cells[i], src.cells[i])
	}
	return g
}

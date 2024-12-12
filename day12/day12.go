package day12

import (
	"AdventOfCode2024/utils"
	"fmt"
)

type Position = utils.Position // X=horizontal, Y=vertical

type Edge struct{ R1, C1, R2, C2 int }

var (
	// Directions: Up, Right, Down, Left
	// Up means Y-1, Right means X+1, Down means Y+1, Left means X-1.
	dirs = []Position{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

	clockwiseNext = map[int][]int{
		0: {1, 2, 3}, // up -> prefer right, down, left
		1: {2, 3, 0}, // right -> prefer down, left, up
		2: {3, 0, 1}, // down -> prefer left, up, right
		3: {0, 1, 2}, // left -> prefer up, right, down
	}
)

func Run() {
	fmt.Println("\nDay 12:")
	grid := utils.ReadGrid("input12.txt")
	puzzle(grid)
}

func puzzle(grid []string) {
	H := len(grid)
	if H == 0 {
		fmt.Println("No input")
		return
	}
	W := len(grid[0])

	field := make([][]rune, H)
	for y := 0; y < H; y++ {
		field[y] = []rune(grid[y])
	}

	visited := make([][]bool, H)
	for i := range visited {
		visited[i] = make([]bool, W)
	}

	regions := findRegions(field, visited, H, W)

	totalPart1 := 0
	totalPart2 := 0
	for _, region := range regions {
		boundaryEdges, perimeter := buildBoundaryEdges(region, H, W)
		area := len(region)
		loopsSides := findLoopsSides(boundaryEdges)
		totalPart1 += area * perimeter
		totalPart2 += area * loopsSides
	}

	fmt.Println("Answer to puzzle 1:", totalPart1)
	fmt.Println("Answer to puzzle 2:", totalPart2)
}

func findRegions(field [][]rune, visited [][]bool, H, W int) [][]Position {
	var regions [][]Position
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if visited[y][x] {
				continue
			}
			ch := field[y][x]
			stack := []Position{{X: x, Y: y}}
			visited[y][x] = true
			var region []Position
			for len(stack) > 0 {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				region = append(region, cur)
				for _, d := range dirs {
					nx, ny := cur.X+d.X, cur.Y+d.Y
					if ny >= 0 && ny < H && nx >= 0 && nx < W &&
						!visited[ny][nx] && field[ny][nx] == ch {
						visited[ny][nx] = true
						stack = append(stack, Position{X: nx, Y: ny})
					}
				}
			}
			regions = append(regions, region)
		}
	}
	return regions
}

func buildBoundaryEdges(region []Position, H, W int) (map[Edge]bool, int) {
	regionSet := make(map[Position]bool, len(region))
	for _, p := range region {
		regionSet[p] = true
	}

	boundaryEdges := make(map[Edge]bool)
	perimeter := 0
	for _, p := range region {
		x, y := p.X, p.Y
		for i, d := range dirs {
			nx, ny := x+d.X, y+d.Y
			if ny < 0 || ny >= H || nx < 0 || nx >= W || !regionSet[Position{X: nx, Y: ny}] {
				perimeter++
				p1r, p1c, p2r, p2c := edgeCorners(x, y, i)
				if p1r > p2r || (p1r == p2r && p1c > p2c) {
					p1r, p2r = p2r, p1r
					p1c, p2c = p2c, p1c
				}
				boundaryEdges[Edge{p1r, p1c, p2r, p2c}] = true
			}
		}
	}
	return boundaryEdges, perimeter
}

// i: direction index in dirs: 0=up,1=right,2=down,3=left
func edgeCorners(x, y, dirIdx int) (p1r, p1c, p2r, p2c int) {
	switch dirIdx {
	case 0: // up
		return y, x, y, x + 1
	case 1: // right
		return y, x + 1, y + 1, x + 1
	case 2: // down
		return y + 1, x, y + 1, x + 1
	default: // left
		return y, x, y + 1, x
	}
}

func findLoopsSides(boundaryEdges map[Edge]bool) int {
	cornerAdj := make(map[[2]int][][2]int)
	var allEdges [][2][2]int
	for e := range boundaryEdges {
		p1 := [2]int{e.R1, e.C1}
		p2 := [2]int{e.R2, e.C2}
		cornerAdj[p1] = append(cornerAdj[p1], p2)
		cornerAdj[p2] = append(cornerAdj[p2], p1)
		allEdges = append(allEdges, [2][2]int{p1, p2})
	}

	used := make(map[[2][2]int]bool)
	isUsed := func(a, b [2]int) bool { return used[normalizeEdge(a, b)] }
	useEdge := func(a, b [2]int) { used[normalizeEdge(a, b)] = true }

	totalSides := 0

	for {
		startFound, startA, startB := findStartEdge(allEdges, isUsed)
		if !startFound {
			break
		}

		loopPath := traceLoop(startA, startB, cornerAdj, isUsed, useEdge)
		totalSides += countSides(loopPath)
	}

	return totalSides
}

func findStartEdge(allEdges [][2][2]int, isUsed func([2]int, [2]int) bool) (bool, [2]int, [2]int) {
	found := false
	var bestA, bestB [2]int
	for _, E := range allEdges {
		a, b := E[0], E[1]
		if !isUsed(a, b) {
			if !found {
				bestA, bestB = a, b
				found = true
			} else {
				if lexLess(a, bestA) || (a == bestA && lexLess(b, bestB)) {
					bestA, bestB = a, b
				}
			}
		}
	}
	return found, bestA, bestB
}

func traceLoop(startA, startB [2]int, cornerAdj map[[2]int][][2]int,
	isUsed func([2]int, [2]int) bool, useEdge func([2]int, [2]int)) [][2]int {

	loopPath := [][2]int{startA, startB}
	useEdge(startA, startB)
	prev, cur := startA, startB
	d := getDir(prev, cur)

	for cur != startA {
		candidates := unusedCandidates(cur, prev, cornerAdj, isUsed)
		nextP := pickNextCorner(cur, d, candidates)
		loopPath = append(loopPath, nextP)
		useEdge(cur, nextP)
		d = getDir(cur, nextP)
		prev, cur = cur, nextP
	}
	return loopPath
}

func unusedCandidates(cur, prev [2]int, cornerAdj map[[2]int][][2]int,
	isUsed func([2]int, [2]int) bool) [][2]int {

	var candidates [][2]int
	for _, nb := range cornerAdj[cur] {
		if nb != prev && !isUsed(cur, nb) {
			candidates = append(candidates, nb)
		}
	}
	return candidates
}

func pickNextCorner(cur [2]int, d int, candidates [][2]int) [2]int {
	if len(candidates) == 1 {
		return candidates[0]
	}
	for _, ndir := range clockwiseNext[d] {
		for _, cand := range candidates {
			if getDir(cur, cand) == ndir {
				return cand
			}
		}
	}
	return candidates[0]
}

func countSides(loop [][2]int) int {
	directions := make([]int, len(loop)-1)
	for i := 0; i < len(loop)-1; i++ {
		directions[i] = getDir(loop[i], loop[i+1])
	}
	sides := 1
	for i := 1; i < len(directions); i++ {
		if directions[i] != directions[i-1] {
			sides++
		}
	}
	return sides
}

func getDir(a, b [2]int) int {
	switch {
	case a[0] == b[0] && b[1] > a[1]:
		return 1 // right
	case a[0] == b[0] && b[1] < a[1]:
		return 3 // left
	case b[1] > a[1]:
		return 2 // down
	default:
		return 0 // up
	}
}

func normalizeEdge(a, b [2]int) [2][2]int {
	if lexLess(a, b) {
		return [2][2]int{a, b}
	}
	return [2][2]int{b, a}
}

func lexLess(a, b [2]int) bool {
	return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
}

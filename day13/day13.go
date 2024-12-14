package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	Px, Py int
}

func Run() {
	fmt.Println("\nDay 13:")
	puzzle1()
}

func puzzle1() {

	machines, err := parseInput()
	if err != nil {
		log.Fatal(err)
	}

	totalCost := 0
	solvableCount := 0

	for _, m := range machines {
		cost, ok := solveMachine(m)
		if ok {
			solvableCount++
			totalCost += cost
		}
	}

	fmt.Println("Max prizes: ", solvableCount)
	fmt.Println("Answer to puzzle: ", totalCost)
}

// parseInput parses the text file input into a slice of Machines
func parseInput() ([]Machine, error) {
	file, err := os.Open("input13.txt")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var machines []Machine
	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			// Blank line indicates a separation between machines
			if len(lines) == 3 {
				m, err := parseMachineLines(lines)
				if err != nil {
					return nil, err
				}
				machines = append(machines, m)
				lines = nil
			}
			continue
		}
		lines = append(lines, line)
	}

	// Handle last machine if no trailing blank line
	if len(lines) == 3 {
		m, err := parseMachineLines(lines)
		if err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

// parseMachineLines parses three lines describing a single machine
func parseMachineLines(lines []string) (Machine, error) {
	var m Machine
	// parse button A line
	if err := parseButtonLine(lines[0], &m.Ax, &m.Ay); err != nil {
		return m, err
	}
	// parse button B line
	if err := parseButtonLine(lines[1], &m.Bx, &m.By); err != nil {
		return m, err
	}
	// parse prize line
	if err := parsePrizeLine(lines[2], &m.Px, &m.Py); err != nil {
		return m, err
	}

	return m, nil
}

func parseButtonLine(line string, X, Y *int) error {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return fmt.Errorf("invalid button line: %s", line)
	}

	coords := strings.TrimSpace(parts[1])
	cparts := strings.Split(coords, ",")
	if len(cparts) != 2 {
		return fmt.Errorf("invalid button line coords: %s", coords)
	}

	var xVal int
	_, err := fmt.Sscanf(strings.TrimSpace(cparts[0]), "X+%d", &xVal)
	if err != nil {
		return fmt.Errorf("failed to parse x from %s: %v", cparts[0], err)
	}

	var yVal int
	_, err = fmt.Sscanf(strings.TrimSpace(cparts[1]), "Y+%d", &yVal)
	if err != nil {
		return fmt.Errorf("failed to parse y from %s: %v", cparts[1], err)
	}

	*X = xVal
	*Y = yVal

	return nil
}

func parsePrizeLine(line string, Px, Py *int) error {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return fmt.Errorf("invalid prize line: %s", line)
	}

	coords := strings.TrimSpace(parts[1])
	cparts := strings.Split(coords, ",")
	if len(cparts) != 2 {
		return fmt.Errorf("invalid prize line coords: %s", coords)
	}

	var xVal, yVal int
	if _, err := fmt.Sscanf(strings.TrimSpace(cparts[0]), "X=%d", &xVal); err != nil {
		return fmt.Errorf("failed to parse Px from %s: %v", cparts[0], err)
	}
	if _, err := fmt.Sscanf(strings.TrimSpace(cparts[1]), "Y=%d", &yVal); err != nil {
		return fmt.Errorf("failed to parse Py from %s: %v", cparts[1], err)
	}

	*Px = xVal
	*Py = yVal

	// Adjust prize coordinates for second part
	*Px += 10000000000000
	*Py += 10000000000000
	return nil
}

func solveMachine(m Machine) (int, bool) {
	D := m.Ax*m.By - m.Ay*m.Bx
	if D == 0 {
		return 0, false
	}

	numeratorX := m.Px*m.By - m.Bx*m.Py
	numeratorY := m.Ax*m.Py - m.Ay*m.Px

	if numeratorX%D != 0 || numeratorY%D != 0 {
		return 0, false
	}

	x := numeratorX / D
	y := numeratorY / D

	if x < 0 || y < 0 {
		return 0, false
	}

	cost := 3*x + y
	return cost, true
}

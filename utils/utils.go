package utils

import (
	"bufio"
	"log"
	"os"
)

type Position struct {
	X, Y int
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadGrid(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer func(file *os.File) {
		err := file.Close()
		Check(err)
	}(file)

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	Check(scanner.Err())
	return grid
}

package day09

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"os"
)

type Disk struct {
	blocks []int // -1 for space, >= 0 for file blocks
}

func Run() {
	fmt.Println("\nDay 9:")
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("input9.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	var disk Disk
	fileIndex := 0
	for index := range input {
		number := int(input[index] - '0')
		if index%2 == 0 {
			for n := 0; n < number; n++ {
				disk.blocks = append(disk.blocks, fileIndex)
			}
			fileIndex++
		} else {
			for n := 0; n < number; n++ {
				disk.blocks = append(disk.blocks, -1)
			}
		}
	}
	orderedDisk := move(disk)
	sum := checksum(orderedDisk)
	fmt.Println("Answer to first puzzle: ", sum)
}

func move(disk Disk) Disk {
	leftSpace := -1
	for i := 0; i < len(disk.blocks); i++ {
		if disk.blocks[i] == -1 {
			leftSpace = i
			break
		}
	}

	rightBlock := -1
	for i := len(disk.blocks) - 1; i >= 0; i-- {
		if disk.blocks[i] != -1 {
			rightBlock = i
			break
		}
	}

	disk.blocks[leftSpace] = disk.blocks[rightBlock]
	disk.blocks[rightBlock] = -1
	if completeOrganized(disk) {
		return disk
	} else {
		return move(disk)
	}
}

func completeOrganized(disk Disk) bool {
	var foundEmptySpace bool
	for i := 0; i < len(disk.blocks); i++ {
		if disk.blocks[i] >= 0 {
			if !foundEmptySpace {
				continue
			} else {
				return false
			}
		} else if disk.blocks[i] == -1 {
			foundEmptySpace = true
		}
	}
	return true
}
func checksum(disk Disk) int {
	sum := 0

	for i := 0; i < len(disk.blocks); i++ {
		if disk.blocks[i] != -1 {
			sum += i * disk.blocks[i]
		}
	}

	return sum
}

func puzzle2() {
	file, err := os.Open("input9.txt")
	utils.Check(err)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	var disk Disk
	fileIndex := 0
	for index := range input {
		number := int(input[index] - '0')
		if index%2 == 0 {
			for n := 0; n < number; n++ {
				disk.blocks = append(disk.blocks, fileIndex)
			}
			fileIndex++
		} else {
			for n := 0; n < number; n++ {
				disk.blocks = append(disk.blocks, -1)
			}
		}
	}
	orderedDisk := moveWholeFile(disk)
	sum := checksum(orderedDisk)
	fmt.Println("Answer to second puzzle: ", sum)
}

func moveWholeFile(disk Disk) Disk {
	files := make(map[int][]int)
	maxFileID := -1
	for pos, blockID := range disk.blocks {
		if blockID != -1 {
			files[blockID] = append(files[blockID], pos)
			if blockID > maxFileID {
				maxFileID = blockID
			}
		}
	}

	for fileID := maxFileID; fileID >= 0; fileID-- {
		filePositions, exists := files[fileID]
		if !exists {
			continue
		}

		fileSize := len(filePositions)

		spaceStart := -1
		spaceSize := 0

		for pos := 0; pos < len(disk.blocks); pos++ {
			if disk.blocks[pos] == -1 {
				if spaceStart == -1 {
					spaceStart = pos
				}
				spaceSize++

				if spaceSize >= fileSize {
					break
				}
			} else {
				spaceStart = -1
				spaceSize = 0
			}
		}

		if spaceSize >= fileSize && spaceStart < filePositions[0] {
			for _, oldPos := range filePositions {
				disk.blocks[oldPos] = -1
			}

			for i := 0; i < fileSize; i++ {
				disk.blocks[spaceStart+i] = fileID
			}
		}
	}

	return disk
}

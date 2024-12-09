package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"strconv"
)

// File is a struct that defines a file with an index and length
type File struct {
	index, length int
}

// Disk is a list of file-blocks where each int defines the index of the file it belongs to
type Disk []int

// CountedElement is a struct that defines a block of elements with a start index and count
type CountedElement struct {
	element, startIndex, count int
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	var files []File
	currentIndex, totalSize := 0, 0
	for _, line := range input {
		for i := 0; i < len(line); i += 2 {
			fileLength, _ := strconv.Atoi(string(line[i]))
			files = append(files, File{index: currentIndex, length: fileLength})
			totalSize += fileLength

			if i+1 < len(line) {
				emptySpace, _ := strconv.Atoi(string(line[i+1]))
				currentIndex += fileLength + emptySpace
				totalSize += emptySpace
			}
		}
	}

	// Create 2 disks (for part1 and part2), default all values to -1 (empty space)
	disk1, disk2 := make(Disk, totalSize), make(Disk, totalSize)
	for index, _ := range disk1 {
		disk1[index], disk2[index] = -1, -1
	}
	for index, file := range files {
		for i := range file.length {
			disk1[file.index+i], disk2[file.index+i] = index, index
		}
	}

	// Part 1
	fmt.Println(disk1.defragmentA().getCheckSum())

	// Part 2
	fmt.Println(disk2.defragmentB().getCheckSum())
}

func (d Disk) defragmentA() Disk {
	for leftIndex, _ := range d { // Find leftmost space (-1)
		for rightIndex := len(d) - 1; rightIndex > leftIndex; rightIndex-- { // Find rightmost non-space (other number)
			if d[leftIndex] == -1 && d[rightIndex] != -1 {
				d[leftIndex], d[rightIndex] = d[rightIndex], d[leftIndex] // Swap values
			}
		}
	}
	return d
}

func (d Disk) defragmentB() Disk {
	fileBlocks := d.getFileBlocks()

	// Keep track of the lowest processed index, limited to processing each once
	maxIndex := fileBlocks[len(fileBlocks)-1].element + 1

	for i := len(fileBlocks) - 1; i > 0; i-- {
		if fileBlocks[i].element == -1 {
			continue
		}

		if fileBlocks[i].element < maxIndex {
			// Find left-most empty space
			for j := 0; j < i; j++ {
				if fileBlocks[j].element == -1 && fileBlocks[j].count >= fileBlocks[i].count {
					// Move file to empty space
					for k := range fileBlocks[i].count {
						d[fileBlocks[j].startIndex+k], d[fileBlocks[i].startIndex+k] = fileBlocks[i].element, -1
					}
					maxIndex = fileBlocks[i].element
					fileBlocks = d.getFileBlocks()
					i = len(fileBlocks) - 1 // Start again from the back (list entries could have changed)
					break
				}
			}
		}
	}

	return d
}

func (d Disk) getFileBlocks() []CountedElement {
	currentFile, currentFileLength, currentFileStart := d[0], 0, 0
	var countedBlocks []CountedElement
	for index, file := range d {
		if file == currentFile { // More blocks from same file
			currentFileLength++
		} else { // Entering new file
			countedBlocks = append(countedBlocks, CountedElement{startIndex: currentFileStart, element: currentFile, count: currentFileLength})
			currentFile = file
			currentFileStart = index
			currentFileLength = 1
		}
	}
	countedBlocks = append(countedBlocks, CountedElement{startIndex: currentFileStart, element: currentFile, count: currentFileLength})

	return countedBlocks
}

func (d Disk) getCheckSum() int {
	fileBlocks := d.getFileBlocks()
	pos, result := 0, 0
	for _, fileBlock := range fileBlocks {
		for range fileBlock.count {
			if fileBlock.element != -1 {
				result += pos * fileBlock.element
			}
			pos++
		}
	}

	return result
}

func (d Disk) print() {
	for _, val := range d {
		fmt.Println(val)
	}
}

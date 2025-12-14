package main

import (
	"fmt"
	"strconv"

	"github.com/mzietara/advent-of-code/util"
)

const emptySpace = -1

func main() {
	util.Timer()()
	part1()
	part2()
}

func part1() {
	disk := make([]int, 0)
	var curEmptyPosition int

	util.ProcessInputFile(func(_ int, line string) {

		for i := 0; i < len(line); i += 2 {
			fileSize, _ := strconv.Atoi(string(line[i]))
			if i == 0 {
				curEmptyPosition = fileSize
			}
			for j := 0; j < fileSize; j++ {
				disk = append(disk, i/2)
			}
			if i+1 < len(line) {
				emptySize, _ := strconv.Atoi(string(line[i+1]))
				for j := 0; j < emptySize; j++ {
					disk = append(disk, emptySpace)
				}
			}
		}
	})

	for i := len(disk) - 1; i >= 0 && curEmptyPosition < i; i-- {
		if disk[i] == emptySpace {
			continue
		}
		disk[curEmptyPosition], disk[i] = disk[i], emptySpace
		curEmptyPosition++
		for curEmptyPosition < len(disk) && disk[curEmptyPosition] != emptySpace {
			curEmptyPosition++
		}
	}

	output := ""
	for _, v := range disk {
		if v == emptySpace {
			output += "."
		} else {
			output += strconv.Itoa(v)
		}
	}

	//find checksum
	checksum := 0
	for i := 0; i < len(disk) && disk[i] != emptySpace; i++ {
		checksum += i * disk[i]
	}
	fmt.Println(checksum)
}

func part2() {
	disk := make([]int, 0)
	fileSizes := make([]int, 0)

	util.ProcessInputFile(func(_ int, line string) {

		for i := 0; i < len(line); i += 2 {
			fileSize, _ := strconv.Atoi(string(line[i]))
			fileSizes = append(fileSizes, fileSize)
			for j := 0; j < fileSize; j++ {
				disk = append(disk, i/2)
			}
			if i+1 < len(line) {
				emptySize, _ := strconv.Atoi(string(line[i+1]))
				for j := 0; j < emptySize; j++ {
					disk = append(disk, emptySpace)
				}
			}
		}
	})

	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == emptySpace {
			continue
		}
		size := fileSizes[disk[i]]
		//find a spot for it
		for j := 0; j < i; j++ {
			//calc size of empty spaces
			if disk[j] == emptySpace {
				emptySpaceSize := 0
				for k := j; k < len(disk) && disk[k] == emptySpace; k++ {
					emptySpaceSize++
				}
				if emptySpaceSize >= size {
					for k := j; k < j+size; k++ {
						disk[k], disk[i-k+j] = disk[i-k+j], emptySpace
					}
					// your are doing i-- already, so add one to this
					i -= size - 1
					break
				}
			}
		}
	}

	//find checksum
	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != emptySpace {
			checksum += i * disk[i]
		}
	}
	fmt.Println(checksum)
}

func printDisk(disk []int) {
	for i := 0; i < len(disk); i++ {
		if disk[i] == emptySpace {
			fmt.Print(".")
		} else {
			fmt.Print(disk[i])
		}
	}
	fmt.Println()
}

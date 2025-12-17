package main

import (
	"github.com/mzietara/advent-of-code/util"
)

type Object byte

const (
	Empty    = '.'
	Splitter = '^'
	Start    = 'S'
)

func part2() {
	grid, start := processInput()

	memo := make(map[[2]int]int)

	var countPaths func(row, col int) int
	countPaths = func(row, col int) int {
		if row == len(grid)-1 {
			return 1
		}

		if col < 0 || col >= len(grid[0]) {
			return 0
		}

		key := [2]int{row, col}
		if val, ok := memo[key]; ok {
			return val
		}

		paths := 0

		// Check current cell
		if grid[row][col] == Splitter {
			paths += countPaths(row+1, col-1)
			paths += countPaths(row+1, col+1)
		} else {
			paths += countPaths(row+1, col)
		}

		memo[key] = paths
		return paths
	}

	result := countPaths(start[0], start[1])
	println(result)
}

func part1() {
	grid, start := processInput()

	curBeams := [][2]int{start}
	splittersVisited := make(map[[2]int]bool)

	for len(curBeams) > 0 {
		beam := curBeams[0]
		curBeams = curBeams[1:]

		for i := beam[0]; i < len(grid); i++ {
			if grid[i][beam[1]] == Splitter {
				if !splittersVisited[[2]int{i, beam[1]}] {
					splittersVisited[[2]int{i, beam[1]}] = true
				}
				if beam[1] > 0 {
					if isUniqueBeam(grid, curBeams, [2]int{i, beam[1] - 1}) {
						curBeams = append(curBeams, [2]int{i, beam[1] - 1})
					}
				}
				if beam[1] < len(grid[i])-1 {
					if isUniqueBeam(grid, curBeams, [2]int{i, beam[1] + 1}) {
						curBeams = append(curBeams, [2]int{i, beam[1] + 1})
					}
				}
				break
			}
		}
	}
	println(len(splittersVisited))
}

func isUniqueBeam(grid [][]byte, curBeams [][2]int, beam [2]int) bool {
	for _, b := range curBeams {
		if beam[0] == b[0] && beam[1] == b[1] {
			return false
		}
		if beam[1] == b[1] {
			unique := false

			for k := min(beam[0], b[0]) + 1; k < max(beam[0], b[0]); k++ {
				if grid[k][beam[1]] == Splitter {
					unique = true
					break
				}
			}
			if !unique {
				return false
			}
		}
	}
	return true
}

func processInput() (input [][]byte, start [2]int) {
	input = [][]byte{}
	util.ProcessInputFile(func(i int, line string) {
		input = append(input, []byte(line))
		for j, ch := range line {
			if ch == rune(Start) {
				start = [2]int{i, j}
			}
		}
	})
	return
}

func main() {
	defer util.Timer()()
	part2()
}

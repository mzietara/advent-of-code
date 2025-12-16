package main

import (
	"github.com/mzietara/advent-of-code/util"
)

const rollChar = '@'

func part2() {
	grid := grid()
	sum := 0

	for rolls := rollsToRemove(grid); len(rolls) != 0; rolls = rollsToRemove(grid) {
		sum += len(rolls)
		for _, roll := range rolls {
			grid[roll[0]][roll[1]] = '.'
		}
	}
	println(sum)
}

func rollsToRemove(grid [][]byte) [][2]int {
	result := [][2]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == rollChar && numAdjacentRolls(grid, i, j) < 4 {
				result = append(result, [2]int{i, j})
			}

		}
	}
	return result
}

func part1() {
	grid := grid()
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == rollChar && numAdjacentRolls(grid, i, j) < 4 {
				sum++
			}

		}
	}
	println(sum)
}

func numAdjacentRolls(grid [][]byte, i, j int) int {
	sum := 0
	// check all 8 directions
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		ni := i + dir[0]
		nj := j + dir[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
			continue
		}
		if grid[ni][nj] == rollChar {
			sum++
		}
	}

	return sum
}

func grid() [][]byte {
	grid := [][]byte{}
	util.ProcessInputFile(func(_ int, line string) {
		grid = append(grid, []byte(line))
	})
	return grid
}

func main() {
	defer util.Timer()()
	part1()
	part2()
}

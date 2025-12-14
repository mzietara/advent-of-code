package main

import (
	"fmt"

	. "github.com/mzietara/advent-of-code/util"
)

func main() {
	Timer()()
	part1()
	part2()
}

func part1() {
	grid := ProcessInputMatrixInt("")
	trailheads := map[Coord][]Coord{}
	sum := 0

	grid.Iterate(func(y, x, cell int) {
		if cell == 0 {
			trailheads[Coord{X: x, Y: y}] = []Coord{}
		}
	})

	for trailhead := range trailheads {
		sum += getHikingTrailEnds(true, grid, trailhead).Size()
	}
	fmt.Println(sum)
}

func part2() {
	grid := ProcessInputMatrixInt("")
	trailheads := map[Coord][]Coord{}
	sum := 0

	grid.Iterate(func(y, x, cell int) {
		if cell == 0 {
			trailheads[Coord{X: x, Y: y}] = []Coord{}
		}
	})

	for trailhead := range trailheads {
		sum += getHikingTrailScores(grid, trailhead)
	}
	fmt.Println(sum)
}

func getHikingTrailEnds(goingUp bool, grid Matrix[int], coord Coord) Set[Coord] {
	peaks := NewSet[Coord]()

	curHeight := grid[coord.Y][coord.X]

	endLevel := 9
	if !goingUp {
		endLevel = 0
	}

	if curHeight == endLevel {
		peaks.Add(coord)
		return *peaks
	} else {
		nextHeight := curHeight + 1
		if !goingUp {
			nextHeight = curHeight - 1
		}
		// go up
		if coord.Y-1 >= 0 {
			next := grid[coord.Y-1][coord.X]
			if next == nextHeight {
				peaks.AddSet(getHikingTrailEnds(goingUp, grid, Coord{X: coord.X, Y: coord.Y - 1}))
			}
		}
		// go down
		if coord.Y+1 < len(grid) {
			next := grid[coord.Y+1][coord.X]
			if next == nextHeight {
				peaks.AddSet(getHikingTrailEnds(goingUp, grid, Coord{X: coord.X, Y: coord.Y + 1}))
			}
		}
		// go left
		if coord.X-1 >= 0 {
			next := grid[coord.Y][coord.X-1]
			if next == nextHeight {
				peaks.AddSet(getHikingTrailEnds(goingUp, grid, Coord{X: coord.X - 1, Y: coord.Y}))
			}
		}
		// go right
		if coord.X+1 < len(grid[0]) {
			next := grid[coord.Y][coord.X+1]
			if next == nextHeight {
				peaks.AddSet(getHikingTrailEnds(goingUp, grid, Coord{X: coord.X + 1, Y: coord.Y}))
			}
		}
	}

	return *peaks
}

func getHikingTrailScores(grid Matrix[int], coord Coord) int {
	sum := 0

	curHeight := grid[coord.Y][coord.X]

	endLevel := 9

	if curHeight == endLevel {
		return 1
	} else {
		nextHeight := curHeight + 1
		// go up
		if coord.Y-1 >= 0 {
			next := grid[coord.Y-1][coord.X]
			if next == nextHeight {
				sum += getHikingTrailScores(grid, Coord{X: coord.X, Y: coord.Y - 1})
			}
		}
		// go down
		if coord.Y+1 < len(grid) {
			next := grid[coord.Y+1][coord.X]
			if next == nextHeight {
				sum += (getHikingTrailScores(grid, Coord{X: coord.X, Y: coord.Y + 1}))
			}
		}
		// go left
		if coord.X-1 >= 0 {
			next := grid[coord.Y][coord.X-1]
			if next == nextHeight {
				sum += getHikingTrailScores(grid, Coord{X: coord.X - 1, Y: coord.Y})
			}
		}
		// go right
		if coord.X+1 < len(grid[0]) {
			next := grid[coord.Y][coord.X+1]
			if next == nextHeight {
				sum += getHikingTrailScores(grid, Coord{X: coord.X + 1, Y: coord.Y})
			}
		}
	}

	return sum
}

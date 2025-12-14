package main

import (
	"fmt"

	. "github.com/mzietara/advent-of-code/util"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type edgeMap map[direction]bool

func NewEdgeMap() map[direction]bool {
	return edgeMap{
		up:    false,
		down:  false,
		left:  false,
		right: false,
	}
}

var visited *Set[Coord]

func main() {
	defer Timer()()
	part1()
	part2()
}

func part1() {
	visited = NewSet[Coord]()
	sum := 0
	grid := ProcessInputMatrix[rune]("", func(cell string) rune {
		return rune(cell[0])
	})

	grid.Iterate(func(y, x int, value rune) {
		if visited.Contains(Coord{X: x, Y: y}) {
			return
		}
		set := NewSet[Coord]()
		edges, _, count := calcGardenDetails(grid, Coord{X: x, Y: y}, set)
		sum += edges * count
	})

	fmt.Println(sum)
}

func calcGardenDetails(grid Matrix[rune], start Coord, currentGarden *Set[Coord]) (edges, corners int, count int) {
	el := grid[start.Y][start.X]
	currentGarden.Add(start)
	visited.Add(start)

	//check up
	if !currentGarden.Contains(Coord{X: start.X, Y: start.Y - 1}) {
		if start.Y > 0 && grid[start.Y-1][start.X] == el {
			curEdge, curCorners, curCount := calcGardenDetails(grid, Coord{X: start.X, Y: start.Y - 1}, currentGarden)
			edges += curEdge
			count += curCount
			corners += curCorners
		} else {
			edges++
		}
	}

	//check down
	if !currentGarden.Contains(Coord{X: start.X, Y: start.Y + 1}) {
		if start.Y < len(grid)-1 && grid[start.Y+1][start.X] == el {
			curEdge, curCorners, curCount := calcGardenDetails(grid, Coord{X: start.X, Y: start.Y + 1}, currentGarden)
			edges += curEdge
			count += curCount
			corners += curCorners
		} else {
			edges++
		}
	}

	//check left
	if !currentGarden.Contains(Coord{X: start.X - 1, Y: start.Y}) {
		if start.X > 0 && grid[start.Y][start.X-1] == el {
			curEdge, curCorners, curCount := calcGardenDetails(grid, Coord{X: start.X - 1, Y: start.Y}, currentGarden)
			edges += curEdge
			count += curCount
			corners += curCorners
		} else {
			edges++
		}
	}

	//check right
	if !currentGarden.Contains(Coord{X: start.X + 1, Y: start.Y}) {
		if start.X < len(grid[0])-1 && grid[start.Y][start.X+1] == el {
			curEdge, curCorners, curCount := calcGardenDetails(grid, Coord{X: start.X + 1, Y: start.Y}, currentGarden)
			edges += curEdge
			count += curCount
			corners += curCorners
		} else {
			edges++
		}
	}

	corners += numCorners(grid, start)
	count++

	return edges, corners, count
}

func part2() {
	visited = NewSet[Coord]()
	sum := 0
	grid := ProcessInputMatrix[rune]("", func(cell string) rune {
		return rune(cell[0])
	})

	grid.Iterate(func(y, x int, value rune) {
		if visited.Contains(Coord{X: x, Y: y}) {
			return
		}
		_, corners, count := calcGardenDetails(grid, Coord{X: x, Y: y}, NewSet[Coord]())
		sum += corners * count
	})

	fmt.Println(sum)
}

func numCorners(grid Matrix[rune], start Coord) int {
	count := 0
	el := grid[start.Y][start.X]
	m := NewEdgeMap()
	//outer corners
	if start.Y == 0 || grid[start.Y-1][start.X] != el {
		m[up] = true
	}
	if start.Y == len(grid)-1 || grid[start.Y+1][start.X] != el {
		m[down] = true
	}
	if start.X == 0 || grid[start.Y][start.X-1] != el {
		m[left] = true
	}
	if start.X == len(grid[0])-1 || grid[start.Y][start.X+1] != el {
		m[right] = true
	}

	if m[up] && m[left] {
		count++
	}
	if m[up] && m[right] {
		count++
	}
	if m[down] && m[left] {
		count++
	}
	if m[down] && m[right] {
		count++
	}

	//inner corners
	//top left
	if start.X-1 >= 0 && start.Y-1 >= 0 &&
		grid[start.Y-1][start.X] == el && grid[start.Y][start.X-1] == el &&
		grid[start.Y-1][start.X-1] != el {
		count++
	}
	//top right
	if start.X+1 < len(grid[0]) && start.Y-1 >= 0 &&
		grid[start.Y-1][start.X] == el && grid[start.Y][start.X+1] == el &&
		grid[start.Y-1][start.X+1] != el {
		count++
	}
	//bottom left
	if start.X-1 >= 0 && start.Y+1 < len(grid) &&
		grid[start.Y+1][start.X] == el && grid[start.Y][start.X-1] == el &&
		grid[start.Y+1][start.X-1] != el {
		count++
	}
	//bottom right
	if start.X+1 < len(grid[0]) && start.Y+1 < len(grid) &&
		grid[start.Y+1][start.X] == el && grid[start.Y][start.X+1] == el &&
		grid[start.Y+1][start.X+1] != el {
		count++
	}
	return count
}

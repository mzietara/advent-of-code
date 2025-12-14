package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/mzietara/advent-of-code/util"
	. "github.com/mzietara/advent-of-code/util"
)

var grid Matrix[object]
var robotPos Coord
var moveStack []direction

type object rune

const (
	wall    object = '#'
	empty   object = '.'
	robot   object = '@'
	box     object = 'O'
	bigBoxL object = '['
	bigBoxR object = ']'
)

type direction rune

const (
	up    direction = '^'
	down  direction = 'v'
	left  direction = '<'
	right direction = '>'
)

func main() {
	defer Timer()()
	part1()
	part2()
}

func PrintGrid() {
	for y, row := range grid {
		for x, cell := range row {
			if robotPos.X == x && robotPos.Y == y {
				fmt.Print(string(robot))
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}

func part1() {
	processInput(Part1)
	for _, dir := range moveStack {
		moveRobot(dir)
	}
	PrintGrid()
	fmt.Println(sumGPSCoordinates())
}

func moveRobot(dir direction) {
	switch dir {
	case up:
		if robotPos.Y > 0 {
			nextPos := grid[robotPos.Y-1][robotPos.X]
			if nextPos == wall {
				return
			} else if nextPos == box {
				//i is the last box in the column to move
				i := robotPos.Y - 1
				for i >= 0 && grid[i][robotPos.X] == box {
					i--
				}
				if i == 0 || grid[i][robotPos.X] == wall {
					return
				} else {
					//move everything
					robotPos.Y--
					grid[robotPos.Y][robotPos.X] = empty
					grid[i][robotPos.X] = box
				}
			} else {
				robotPos.Y--
			}
		}
	case down:
		if robotPos.Y < len(grid)-1 {
			nextPos := grid[robotPos.Y+1][robotPos.X]
			if nextPos == wall {
				return
			} else if nextPos == box {
				//i is the last box in the column to move
				i := robotPos.Y + 1
				for i < len(grid) && grid[i][robotPos.X] == box {
					i++
				}
				if i == len(grid)-1 || grid[i][robotPos.X] == wall {
					return
				} else {
					//move everything
					robotPos.Y++
					grid[robotPos.Y][robotPos.X] = empty
					grid[i][robotPos.X] = box
				}
			} else {
				robotPos.Y++
			}
		}
	case left:
		if robotPos.X > 0 {
			nextPos := grid[robotPos.Y][robotPos.X-1]
			if nextPos == wall {
				return
			} else if nextPos == box {
				//i is the last box in the row to move
				i := robotPos.X - 1
				for i >= 0 && grid[robotPos.Y][i] == box {
					i--
				}
				if i == 0 || grid[robotPos.Y][i] == wall {
					return
				} else {
					//move everything
					robotPos.X--
					grid[robotPos.Y][robotPos.X] = empty
					grid[robotPos.Y][i] = box
				}
			} else {
				robotPos.X--
			}
		}
	case right:
		if robotPos.X < len(grid[0])-1 {
			nextPos := grid[robotPos.Y][robotPos.X+1]
			if nextPos == wall {
				return
			} else if nextPos == box {
				//i is the last box in the row to move
				i := robotPos.X + 1
				for i < len(grid[0]) && grid[robotPos.Y][i] == box {
					i++
				}
				if i == len(grid[0])-1 || grid[robotPos.Y][i] == wall {
					return
				} else {
					//move everything
					robotPos.X++
					grid[robotPos.Y][robotPos.X] = empty
					grid[robotPos.Y][i] = box
				}
			} else {
				robotPos.X++
			}
		}
	}
}

func sumGPSCoordinates() int {
	sum := 0
	grid.Iterate(func(y, x int, value object) {
		if value == box {
			sum += x + 100*y
		} else if value == bigBoxL {
			sum += int(math.Min(float64(x), float64(len(grid[0])-1-x))) +
				100*int(math.Min(float64(y), float64(len(grid)-y)))
		}
	})
	return sum
}

func part2() {
	processInput(Part2)
	PrintGrid()
}

func processInput(pp PuzzlePart) {
	robotPos = Coord{}
	moveStack = []direction{}
	grid = Matrix[object]{}
	onGrid := true
	rowNum := 0
	util.ProcessInputFile(func(i int, line string) {
		if strings.TrimSpace(line) == "" {
			onGrid = false
			return
		}
		if onGrid {
			row := []object(line)
			if pp == Part2 {
				newRow := make([]object, len(row)*2)
				for i, cell := range row {
					if cell == empty || cell == wall {
						newRow[i*2] = cell
						newRow[i*2+1] = cell
					} else if cell == robot {
						newRow[i*2] = cell
						newRow[i*2+1] = empty
					} else if cell == box {
						newRow[i*2] = bigBoxL
						newRow[i*2+1] = bigBoxR
					}
				}
				row = newRow
			}
			robotFound := false
			if !robotFound {
				for i, cell := range row {
					if cell == robot {
						robotPos = Coord{X: i, Y: rowNum}
						row[i] = empty
					}
				}
			}
			grid = append(grid, row)
			rowNum++
		} else {
			moveStack = append(moveStack, []direction(line)...)
		}
	})
}

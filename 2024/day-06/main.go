package main

import (
	"fmt"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

const obstacle = '#'
const newObstruction = 'O'
const guard = '^'
const visited = 'X'

var area = make([][]rune, 0)
var startingPosition = make([]int, 2)

type direction int // 0 - up, 1 - right, 2 - down, 3 - left
const (
	Up direction = iota
	Right
	Down
	Left
)

func main() {
	processInput()
	part1()
	part2()
}

func processInput() {
	i := 0
	util.ProcessFile("input.txt", func(_ int, line string) {
		if strings.ContainsRune(line, guard) {
			startingPosition[0] = i
			startingPosition[1] = strings.IndexRune(line, guard)
			line = strings.ReplaceAll(line, string(guard), "X")
		}
		area = append(area, []rune(line))
		i++
	})
}

func part1() {
	pos := make([]int, 2)
	copy(pos, startingPosition)
	dir := Up
	distinctSteps := 1

	for !isLeavingArea(dir, pos[0], pos[1]) {
		pos, dir = move(pos, dir)
		if area[pos[0]][pos[1]] != visited {
			distinctSteps++
			area[pos[0]][pos[1]] = visited
		}
	}
	fmt.Println(distinctSteps)

}

func move(pos []int, dir direction) ([]int, direction) {
	switch dir {
	case Up:
		if area[pos[0]-1][pos[1]] == obstacle {
			dir = nextDirection(dir)
			return pos, dir
		} else {
			pos[0]--
		}
	case Right:
		if area[pos[0]][pos[1]+1] == obstacle {
			dir = nextDirection(dir)
			return pos, dir
		} else {
			pos[1]++
		}
	case Down:
		if area[pos[0]+1][pos[1]] == obstacle {
			dir = nextDirection(dir)
			return pos, dir
		} else {
			pos[0]++
		}
	case Left:
		if area[pos[0]][pos[1]-1] == obstacle {
			dir = nextDirection(dir)
			return pos, dir
		} else {
			pos[1]--
		}
	}
	return pos, dir
}

func isLeavingArea(dir direction, row, col int) bool {
	if (dir == Up && row == 0) ||
		(dir == Right && col == len(area[0])-1) ||
		(dir == Down && row == len(area)-1) ||
		(dir == Left && col == 0) {
		return true
	}
	return false
}

func nextDirection(dir direction) direction {
	return (dir + 1) % 4
}

func part2() {
	sum := 0

	for i := range area {
		for j := range area[i] {
			if (startingPosition[0] == i && startingPosition[1] == j) ||
				area[i][j] == obstacle {
				continue
			}

			pos, dir, visitedObstruction, inLoop, upVisitsMap := make([]int, 2), Up, false, false, make(map[[2]int]int)
			oldValue := area[i][j]
			copy(pos, startingPosition)

			// place new obstruction
			area[i][j] = newObstruction

			//find something you turned from left to up on, and see the number of steps it took to visit again, and if it happened
			//three times, that should be good enough, and only do it after visiting the obstruction
			for !isLeavingArea(dir, pos[0], pos[1]) && !inLoop {
				pos, dir, visitedObstruction, upVisitsMap = moveWithNewObstruction(pos, dir, visitedObstruction, upVisitsMap)
				for _, visits := range upVisitsMap {
					if visits > 2 {
						sum++
						inLoop = true
						break
					}
				}
			}

			// remove new obstruction
			area[i][j] = oldValue
		}
	}
	fmt.Println(sum)
}

func moveWithNewObstruction(
	pos []int,
	dir direction,
	visitedObstruction bool,
	upVisitsMap map[[2]int]int) ([]int, direction, bool, map[[2]int]int) {
	switch dir {
	case Up:
		if area[pos[0]-1][pos[1]] == obstacle {
			if visitedObstruction {
				upVisitsMap[[2]int{pos[0] - 1, pos[1]}]++
			}
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else if area[pos[0]-1][pos[1]] == newObstruction {
			if visitedObstruction {
				upVisitsMap[[2]int{pos[0] - 1, pos[1]}]++
			}
			visitedObstruction = true
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else {
			pos[0]--
		}
	case Right:
		if area[pos[0]][pos[1]+1] == obstacle {
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else if area[pos[0]][pos[1]+1] == newObstruction {
			visitedObstruction = true
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else {
			pos[1]++
		}
	case Down:
		if area[pos[0]+1][pos[1]] == obstacle {
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else if area[pos[0]+1][pos[1]] == newObstruction {
			visitedObstruction = true
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else {
			pos[0]++
		}
	case Left:
		if area[pos[0]][pos[1]-1] == obstacle {
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else if area[pos[0]][pos[1]-1] == newObstruction {
			visitedObstruction = true
			dir = nextDirection(dir)
			return pos, dir, visitedObstruction, upVisitsMap
		} else {
			pos[1]--
		}
	}
	return pos, dir, visitedObstruction, upVisitsMap
}

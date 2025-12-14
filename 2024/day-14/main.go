package main

import (
	"fmt"
	"strings"

	. "github.com/mzietara/advent-of-code/util"
)

var GridSize = Coord{X: 101, Y: 103}

// var GridSize = Coord{X: 11, Y: 7}
var Robots []Robot

type Robot struct {
	P      Coord
	V      Coord
	CurPos *Coord
}

func NewRobot(p, v Coord) Robot {
	return Robot{P: p, V: v, CurPos: &p}
}

func (r Robot) String() string {
	return fmt.Sprintf("p: %v, v: %v", r.P, r.V)
}

func (r *Robot) Move(times int) {
	for i := 0; i < times; i++ {
		r.CurPos.X = (r.CurPos.X + r.V.X) % GridSize.X
		if r.CurPos.X < 0 {
			r.CurPos.X += GridSize.X
		}
		r.CurPos.Y = (r.CurPos.Y + r.V.Y) % GridSize.Y
		if r.CurPos.Y < 0 {
			r.CurPos.Y += GridSize.Y
		}
	}
}

func main() {
	defer Timer()()
	part1()
	part2()
}

type quadrant int

const (
	none quadrant = iota
	topLeft
	topRight
	bottomLeft
	bottomRight
)

func (r Robot) GetQuadrant() quadrant {
	if r.CurPos.X == GridSize.X/2 || r.CurPos.Y == GridSize.Y/2 {
		return none
	}
	if r.CurPos.X < GridSize.X/2 {
		if r.CurPos.Y < GridSize.Y/2 {
			return topLeft
		}
		return bottomLeft
	}
	if r.CurPos.Y < GridSize.Y/2 {
		return topRight
	}
	return bottomRight
}

func MoveAllRobots(times int) {
	for _, r := range Robots {
		r.Move(times)
	}
}

func part1() {
	quadrantCount := [5]int{}
	ProcessInput()
	for _, r := range Robots {
		r.Move(100)
		quadrantCount[r.GetQuadrant()]++
	}
	PrintCurrentState()

	sum := quadrantCount[topLeft] * quadrantCount[topRight] * quadrantCount[bottomLeft] * quadrantCount[bottomRight]
	fmt.Println(sum)
}

func part2() {
	ProcessInput()
	found := false
	count := 0
	for !found {
		MoveAllRobots(1)
		found = FindLineOfRobotsAndStump()
		count++
	}
	fmt.Println(count)
}

func FindLineOfRobotsAndStump() bool {
	grid := NewMatrix[int](GridSize.Y, GridSize.X)
	for _, r := range Robots {
		grid[r.CurPos.Y][r.CurPos.X] += 1
	}
	for _, row := range grid {
		lineCount := 0
		for i := 0; i < len(row)-1; i += 2 {
			if row[i] == 1 && row[i+1] == 1 {
				lineCount++
			} else {
				lineCount = 0
			}
			if lineCount == 10 {
				fmt.Println("Found line of robots")
				PrintCurrentState()
				return true
			}
		}
	}
	return false
}

func PrintCurrentState() {
	grid := NewMatrix[int](GridSize.Y, GridSize.X)
	for _, r := range Robots {
		grid[r.CurPos.Y][r.CurPos.X] += 1
	}
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func ProcessInput() {
	Robots = []Robot{}
	ProcessInputFile(func(_ int, line string) {
		s := strings.Split(strings.TrimPrefix(line, "p="), " v=")
		startStr := strings.Split(s[0], ",")
		start := Coord{X: StringToInt(startStr[0]), Y: StringToInt(startStr[1])}
		vStr := strings.Split(s[1], ",")
		v := Coord{X: StringToInt(vStr[0]), Y: StringToInt(vStr[1])}
		Robots = append(Robots, NewRobot(start, v))

	})

}

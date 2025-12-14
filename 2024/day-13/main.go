package main

import (
	"fmt"
	"strings"

	. "github.com/mzietara/advent-of-code/util"
)

type Game struct {
	Prize, ButtonA, ButtonB Coord
}

func (g Game) String() string {
	return fmt.Sprintf("Prize: %v, Button A: %v, Button B: %v", g.Prize, g.ButtonA, g.ButtonB)
}

func NewGame(prize, buttonA, buttonB Coord) Game {
	return Game{Prize: prize, ButtonA: buttonA, ButtonB: buttonB}
}

func main() {
	defer Timer()()
	part1()
	part2()
}

func part1() {
	sumSolvableGameTokens(Part1)
}

func part2() {
	sumSolvableGameTokens(Part2)
}

func sumSolvableGameTokens(pp PuzzlePart) {
	games := processInput(pp)
	totalTokens := 0
	for _, game := range games {
		bCount := ((game.ButtonA.X * game.Prize.Y) - (game.ButtonA.Y * game.Prize.X)) /
			((game.ButtonA.X * game.ButtonB.Y) - (game.ButtonA.Y * game.ButtonB.X))
		aCount := (game.Prize.X - (game.ButtonB.X * bCount)) / game.ButtonA.X

		if game.Prize.X != (game.ButtonA.X*aCount)+(game.ButtonB.X*bCount) ||
			game.Prize.Y != (game.ButtonA.Y*aCount)+(game.ButtonB.Y*bCount) {
			continue
		}
		totalTokens += int(aCount)*3 + int(bCount)
	}
	fmt.Println("Total tokens:", totalTokens)
}

func processInput(pp PuzzlePart) []Game {
	games := make([]Game, 0)
	gameLine := 0
	curGame := Game{}
	ProcessInputFile(func(_ int, line string) {
		if gameLine == 0 {
			split := strings.Split(strings.TrimPrefix(line, "Button A: X+"), ", Y+")
			curGame.ButtonA = Coord{X: StringToInt(split[0]), Y: StringToInt(split[1])}
		} else if gameLine == 1 {
			split := strings.Split(strings.TrimPrefix(line, "Button B: X+"), ", Y+")
			curGame.ButtonB = Coord{X: StringToInt(split[0]), Y: StringToInt(split[1])}
		} else if gameLine == 2 {
			split := strings.Split(strings.TrimPrefix(line, "Prize: X="), ", Y=")
			curGame.Prize = Coord{X: StringToInt(split[0]), Y: StringToInt(split[1])}

			if pp == Part2 {
				curGame.Prize.X += 10000000000000
				curGame.Prize.Y += 10000000000000
			}
			games = append(games, curGame)
		} else if gameLine == 3 {
			gameLine, curGame = 0, Game{}
			return
		}
		gameLine++
	})
	return games
}

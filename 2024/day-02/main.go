package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

func part1() {
	sum := 0

	util.ProcessFile("input.txt", func(_ int, line string) {
		levelsRaw := strings.Split(line, " ")
		lastLevel, _ := strconv.Atoi(levelsRaw[0])

		isIncreasing := 0

		curLevel, _ := strconv.Atoi(levelsRaw[1])

		if lastLevel < curLevel {
			isIncreasing = 1
		} else if lastLevel > curLevel {
			isIncreasing = -1
		}

		safe := true

		for _, levelRaw := range levelsRaw[1:] {
			curLevel, _ = strconv.Atoi(levelRaw)
			diff := math.Abs(float64(lastLevel - curLevel))
			if (isIncreasing == 1 && lastLevel < curLevel ||
				isIncreasing == -1 && lastLevel > curLevel) &&
				diff >= 1 && diff <= 3 {
				lastLevel = curLevel
			} else {
				safe = false
				break
			}
		}
		if safe {
			sum++
		}

	})
	fmt.Println(sum)
}

func isSafe(levels []int) bool {
	lastLevel := levels[0]
	curLevel := levels[1]
	isIncreasing := 0

	if lastLevel < curLevel {
		isIncreasing = 1
	} else if lastLevel > curLevel {
		isIncreasing = -1
	}

	for _, curLevel := range levels[1:] {
		diff := math.Abs(float64(lastLevel - curLevel))
		if (isIncreasing == 1 && lastLevel < curLevel ||
			isIncreasing == -1 && lastLevel > curLevel) &&
			diff >= 1 && diff <= 3 {
			lastLevel = curLevel
		} else {
			return false
		}
	}
	return true
}

func lineToLevels(line string) []int {
	levelsRaw := strings.Split(line, " ")
	levels := make([]int, len(levelsRaw))
	for i, levelRaw := range levelsRaw {
		levels[i], _ = strconv.Atoi(levelRaw)
	}
	return levels
}

func part2() {
	sum := 0

	util.ProcessFile("input.txt", func(_ int, line string) {
		levels := lineToLevels(line)
		if existsSafeLevelSubset(levels) {
			sum++
		}
	})
	fmt.Println(sum)
}

func existsSafeLevelSubset(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := range levels {
		subLevels := make([]int, 0, len(levels)-1)
		subLevels = append(subLevels, levels[:i]...)
		subLevels = append(subLevels, levels[i+1:]...)

		if isSafe(subLevels) {
			return true
		}
	}

	return false
}

func main() {
	part2()
}

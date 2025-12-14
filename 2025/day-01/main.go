package main

import (
	"fmt"

	"github.com/mzietara/advent-of-code/util"
)

func part1() {
	rotations := rotationsArray()
	currentNumber := 50
	count := 0

	for _, rotation := range rotations {
		currentNumber = (currentNumber + rotation.Clicks*int(rotation.Dir)) % 100
		if currentNumber < 0 {
			currentNumber += 100
		}
		fmt.Printf("rotation: %v, currentNumber: %d\n", rotation, currentNumber)
		if currentNumber == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func part2() {
	rotations := rotationsArray()
	currentNumber := 50
	count := 0

	for _, rotation := range rotations {
		// first see how many rotations above 100 there are and count those
		count += rotation.Clicks / 100

		if currentNumber != 0 {
			if rotation.Dir == right && (currentNumber+(rotation.Clicks%100)) >= 100 {
				// if going right and the current number plus the remainder of clicks crosses 100, count one more
				count++
			} else if rotation.Dir == left && (currentNumber-(rotation.Clicks%100)) <= 0 {
				// if going left and the current number minus the remainder of clicks crosses 0, count one more
				count++
			}
		}

		currentNumber = (currentNumber + rotation.Clicks*int(rotation.Dir)) % 100
		if currentNumber < 0 {
			currentNumber += 100
		}
		fmt.Printf("rotation: %v, currentNumber: %d\n", rotation, currentNumber)
	}
	fmt.Println(count)
}

func rotationsArray() []Rotation {
	rotations := []Rotation{}
	util.ProcessInputFile(func(_ int, line string) {
		var dir direction
		if line[0] == 'L' {
			dir = left
		} else if line[0] == 'R' {
			dir = right
		} else {
			panic("unknown direction")
		}

		clicks := 0
		fmt.Sscanf(line[1:], "%d", &clicks)

		rotations = append(rotations, Rotation{Dir: dir, Clicks: clicks})
	})

	return rotations
}

type Rotation struct {
	Dir    direction
	Clicks int
}

type direction int

const (
	left  direction = -1
	right direction = 1
)

func main() {
	defer util.Timer()()
	// part1()
	part2()
}

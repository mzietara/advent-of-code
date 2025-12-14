package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

var countMap = make(map[[2]int]int)

func main() {
	defer util.Timer()()
	part1()
	part2()
}

func part1() {
	countAllStones(25)
}
func part2() {
	countAllStones(75)
}

func countAllStones(depth int) {
	countMap = make(map[[2]int]int)
	stones := []int{}
	util.ProcessInputFile(func(_ int, line string) {
		for _, s := range strings.Split(line, " ") {
			stone, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			stones = append(stones, stone)
		}
	})

	sum := 0
	for _, stone := range stones {
		sum += countStones(stone, 0, depth)
	}
	fmt.Println(sum)
}

func countStones(face int, depth int, max int) int {
	stones := 0
	if depth == max {
		return 1
	}

	if countMap[[2]int{face, depth}] != 0 {

		return countMap[[2]int{face, depth}]
	}

	if face == 0 {
		stones = countStones(1, depth+1, max)
		countMap[[2]int{face, depth}] = stones
		return stones
	}
	str := strconv.Itoa(face)
	if len(str)%2 == 0 {
		mid := len(str) / 2
		left, _ := strconv.Atoi(str[0:mid])
		right, _ := strconv.Atoi(str[mid:])
		stones = countStones(left, depth+1, max) +
			countStones(right, depth+1, max)
		countMap[[2]int{face, depth}] = stones
		return stones
	}
	stones = countStones(face*2024, depth+1, max)
	countMap[[2]int{face, depth}] = stones
	return stones
}

// func countStones(numBlinks int) {
// 	stones := []int{}
// 	util.ProcessInputFile(func(line string) {
// 		for _, s := range strings.Split(line, " ") {
// 			stone, err := strconv.Atoi(s)
// 			if err != nil {
// 				panic(err)
// 			}
// 			stones = append(stones, stone)
// 		}
// 	})

//   sum := 0
// 	for i := 0; i < len(stones); i++ {
// 		if stones[i] == 0 {
// 			sum++
// 		} else if len(strconv.Itoa(stones[i]))%2 == 0 {
// 			stoneStr := strconv.Itoa(stones[i])
// 		} else {
// 			stones[i] = stones[i] * 2024
// 		}

// 	// for m := 0; m < numBlinks; m++ {
// 	// 	fmt.Println("blink", m)
// 	// 	for i := 0; i < len(stones); i++ {
// 	// 		if stones[i] == 0 {
// 	// 			stones[i] = 1
// 	// 		} else if len(strconv.Itoa(stones[i]))%2 == 0 {
// 	// 			stoneStr := strconv.Itoa(stones[i])
// 	// 			leftStone, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
// 	// 			rightStone, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
// 	// 			stones[i] = leftStone
// 	// 			stones = append(stones[:i+1], append([]int{rightStone}, stones[i+1:]...)...)
// 	// 			i++
// 	// 		} else {
// 	// 			stones[i] = stones[i] * 2024
// 	// 		}
// 	// 	}

// 	// }
// 	fmt.Println(len(stones))

// }

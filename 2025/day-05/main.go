package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

func part2() {
	freshIngRanges, _ := processInput()
	sum := int64(0)

	// check if any ranges overlap and merge them
	for i := range freshIngRanges {
		for j := i + 1; j < len(freshIngRanges); j++ {
			rng1 := freshIngRanges[i]
			rng2 := freshIngRanges[j]
			if (rng1[0] <= rng2[1] && rng1[1] >= rng2[0]) || (rng2[0] <= rng1[1] && rng2[1] >= rng1[0]) ||
				(rng1[0] >= rng2[0] && rng1[1] <= rng2[1]) || (rng2[0] >= rng1[0] && rng2[1] <= rng1[1]) {
				// merge ranges
				newRange := [2]int64{min(rng1[0], rng2[0]), max(rng1[1], rng2[1])}
				freshIngRanges[i] = newRange
				freshIngRanges = append(freshIngRanges[:j], freshIngRanges[j+1:]...)
				j = i
			}
		}
	}

	for _, rng := range freshIngRanges {
		sum += rng[1] - rng[0] + 1
	}
	fmt.Println(sum)
}

func part1() {
	freshIngRanges, ing := processInput()
	sum := 0

	for _, num := range ing {
		for _, rng := range freshIngRanges {
			if num >= rng[0] && num <= rng[1] {
				sum++
				break
			}
		}
	}
	fmt.Println(sum)
}

func processInput() (freshIngRanges [][2]int64, ing []int64) {
	freshIngRanges = [][2]int64{}
	ing = []int64{}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	processingRanges := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			processingRanges = false
			continue
		}
		if processingRanges {
			rngStr := strings.Split(line, "-")
			start, _ := strconv.Atoi(rngStr[0])
			end, _ := strconv.Atoi(rngStr[1])
			freshIngRanges = append(freshIngRanges, [2]int64{int64(start), int64(end)})
		} else {
			num, _ := strconv.ParseInt(line, 10, 64)
			ing = append(ing, num)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func main() {
	defer util.Timer()()
	part1()
	part2()
}

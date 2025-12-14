package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

// rules has the ordering rules where int keys need to be before all of the values
var rules = make(map[int]map[int]struct{})
var updates = make([][]int, 0)

func main() {
	processInput()
	//part1()
	part2()
}

func part1() {
	sum := 0
	for _, update := range updates {
		if isSuccessfulUpdate(update) {
			sum += update[len(update)/2]
		}
	}
	fmt.Println(sum)
}

func isSuccessfulUpdate(update []int) bool {
	correct := true
	for i, num := range update {
		for _, num2 := range update[i+1:] {
			if _, ok := rules[num][num2]; !ok {
				correct = false
				break
			}
		}
		if !correct {
			break
		}
	}
	return correct
}

func processInput() {
	rulesProcessed := false
	util.ProcessFile("input.txt", func(_ int, line string) {
		if strings.TrimSpace(line) == "" {
			rulesProcessed = true
		} else if !rulesProcessed {
			s := strings.Split(line, "|")
			first, _ := strconv.Atoi(s[0])
			second, _ := strconv.Atoi(s[1])

			if _, ok := rules[first]; !ok {
				rules[first] = make(map[int]struct{})
			}
			rules[first][second] = struct{}{}
		} else {
			s := strings.Split(line, ",")
			nums := make([]int, len(s))
			for i, num := range s {
				n, _ := strconv.Atoi(num)
				nums[i] = n
			}
			updates = append(updates, nums)
		}
	})
}

func part2() {
	sum := 0
	for z, update := range updates {

		if !isSuccessfulUpdate(update) {
			fmt.Printf("Processing update %d\n", z)
			size := len(update)
			correctOrder := make([]int, 0)

			for len(correctOrder) != size {
				//get the one that has everything in its map, remove it from update, and add it to correctOrder
				for i, num := range update {
					isLargest := true
					for _, num2 := range update {
						if num != num2 {
							if _, ok := rules[num][num2]; !ok {
								isLargest = false
								break
							}
						}
					}
					if isLargest {
						correctOrder = append(correctOrder, num)
						update = append(update[:i], update[i+1:]...)
						break
					}
				}
			}

			sum += correctOrder[len(correctOrder)/2]
			fmt.Printf("Correct order: %v\n", correctOrder)
		}
	}
	fmt.Println(sum)

}

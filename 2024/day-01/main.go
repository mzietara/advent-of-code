package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

func part1() {
	var list1, list2 []int

	util.ProcessFile("input.txt", func(_ int, line string) {
		tokens := strings.Split(line, "   ")
		num1, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatalf("Failed to convert %s to int: %v", tokens[0], err)
		}
		num2, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Failed to convert %s to int: %v", tokens[1], err)
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)

	})
	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for i := 0; i < len(list1); i++ {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println(sum)
}

func part2() {
	var list1 []int
	var numAppearances = make(map[int]int)

	util.ProcessFile("input.txt", func(_ int, line string) {
		tokens := strings.Split(line, "   ")
		num1, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatalf("Failed to convert %s to int: %v", tokens[0], err)
		}
		list1 = append(list1, num1)

		num2, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Failed to convert %s to int: %v", tokens[1], err)
		}
		numAppearances[num2]++

	})

	sum := 0

	for i := 0; i < len(list1); i++ {
		sum += list1[i] * numAppearances[list1[i]]
	}
	fmt.Println(sum)
}

func main() {
	//part1()
	part2()
}

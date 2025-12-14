package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

type equation struct {
	answer uint64
	nums   []uint64
}

type operator int

const (
	add operator = iota
	multiply
	concatenation
)

func main() {
	defer util.Timer()()
	part1()
	part2()
}

func part1() {
	findSumOfCorrectEquations([]operator{add, multiply})
}

func part2() {
	findSumOfCorrectEquations([]operator{add, multiply, concatenation})
}

func findSumOfCorrectEquations(ops []operator) {
	var equations = make([]equation, 0)
	util.ProcessInputFile(func(i int, line string) {
		equations = append(equations, lineToEquation(line))
	})

	sum := uint64(0)

	for _, eq := range equations {
		if operatorPermutationSolutions(eq.nums, eq.answer, ops, false) {
			sum += eq.answer
		}
	}
	fmt.Println(sum)
}

func lineToEquation(line string) equation {
	s := strings.Split(line, ": ")
	answer, _ := strconv.ParseUint(s[0], 10, 64)
	equation := equation{answer, make([]uint64, 0)}
	nums := strings.Split(s[1], " ")
	for _, num := range nums {
		n, _ := strconv.ParseUint(num, 10, 64)
		equation.nums = append(equation.nums, n)
	}
	return equation
}

func operatorPermutationSolutions(nums []uint64, answer uint64, ops []operator, found bool) bool {
	if found {
		return true
	}
	if len(nums) == 1 {
		if nums[0] == answer {
			return true
		}
	} else {
		for _, op := range ops {
			var newNums []uint64
			if op == multiply {
				newNums = append(newNums, nums[0]*nums[1])
			} else if op == add {
				newNums = append(newNums, nums[0]+nums[1])
			} else if op == concatenation {
				num, _ := strconv.ParseUint(fmt.Sprintf("%d%d", nums[0], nums[1]), 10, 64)
				newNums = append(newNums, num)
			}
			newNums = append(newNums, nums[2:]...)
			newFound := operatorPermutationSolutions(newNums, answer, ops, found)
			if newFound {
				return true
			}
		}
	}
	return false
}

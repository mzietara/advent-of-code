package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

func part2() {
	grid := processInputPart2()
	sum := int64(0)

	equationIndexStart := 0
	var curSum int64
	var curOperator byte
	for col := 0; col < len(grid[0]); col++ {
		endOfEquation := true
		if col == equationIndexStart {
			endOfEquation = false
			// get the curOperator from the last element in the column
			curOperator = grid[len(grid)-1][col]
			if curOperator == '+' {
				curSum = 0
			} else if curOperator == '*' {
				curSum = 1
			} else {
				panic("unknown operator")
			}
		} else {
			// check if we reached the end of the equation which will imply all rows in the column are spaces
			if grid[0][col] == ' ' {
				for row := 1; row < len(grid)-1; row++ {
					if grid[row][col] != ' ' {
						endOfEquation = false
						break
					}
				}
			} else {
				endOfEquation = false
			}
		}

		//build the number and add it to curSum depending on curOperator
		if !endOfEquation {
			numStr := ""
			for row := 0; row < len(grid)-1; row++ {
				if grid[row][col] != ' ' {
					numStr += string(grid[row][col])
				}
			}
			num, _ := strconv.ParseInt(numStr, 10, 64)
			if curOperator == '+' {
				curSum += num
			} else if curOperator == '*' {
				curSum *= num
			}
		}

		if col == len(grid[0])-1 {
			endOfEquation = true
		}
		if endOfEquation {
			sum += curSum
			equationIndexStart = col + 1
			continue
		}
	}
	fmt.Println(sum)
}

func processInputPart2() (input [][]byte) {
	input = [][]byte{}
	util.ProcessInputFile(func(_ int, line string) {
		input = append(input, []byte(line))
	})
	return
}

func part1() {
	equations, operations := processInputPart1()
	sum := 0

	for i := 0; i < len(equations[0]); i++ {
		curSum := 0
		if operations[i] == '*' {
			curSum = 1
		}
		for j := 0; j < len(equations); j++ {
			if operations[i] == '+' {
				curSum += int(equations[j][i])
			} else if operations[i] == '*' {
				curSum *= int(equations[j][i])
			}
		}
		sum += curSum
	}
	println(sum)
}

func processInputPart1() (equations [][]int64, operations []byte) {
	equations = [][]int64{}
	operations = []byte{}
	util.ProcessInputFile(func(_ int, line string) {
		if strings.HasPrefix(line, "+") || strings.HasPrefix(line, "*") {
			ops := strings.Fields(line)
			for _, op := range ops {
				operations = append(operations, op[0])
			}
			return
		}
		numStrs := strings.Fields(line)
		nums := make([]int64, len(numStrs))
		for i, ns := range numStrs {
			nums[i], _ = strconv.ParseInt(ns, 10, 64)
		}
		equations = append(equations, nums)

	})
	return
}

func main() {
	defer util.Timer()()
	part2()
}

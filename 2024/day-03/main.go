package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

const mulRegex = `mul\(\d+,\d+\)`

func main() {
	part2()
	part2()
}

// too low
func part2() {
	sum := 0
	allLines := ""
	util.ProcessFile("input.txt", func(_ int, line string) {
		allLines += line
	})
	// take out anything after `don't()` and resume after `do()`
	splitOnDont := strings.Split(allLines, "don't()")
	sum += sumAllMuls(splitOnDont[0])

	for _, line := range splitOnDont[1:] {
		if strings.Contains(line, "do()") {
			splitOnDo := strings.SplitN(line, "do()", 2)
			sum += sumAllMuls(splitOnDo[1])
		}
	}
	fmt.Println(sum)
}

func part1() {
	allLines := ""
	util.ProcessFile("input.txt", func(_ int, line string) {
		allLines += line
	})
	fmt.Println(sumAllMuls(allLines))
}

func sumAllMuls(s string) int {
	return sumRegexMatches(s)
}

func sumRegexMatches(line string) int {
	sum := 0
	re := regexp.MustCompile(mulRegex)
	matches := re.FindAllString(line, -1)
	for _, match := range matches {
		nums := getIntsFromMulRegexMatch(match)
		sum += nums[0] * nums[1]
	}
	return sum
}

func getIntsFromMulRegexMatch(match string) [2]int {
	split := strings.Split(match, ",")
	num1, _ := strconv.Atoi(strings.Trim(split[0], "mul("))
	num2, _ := strconv.Atoi(strings.Trim(split[1], ")"))
	return [2]int{num1, num2}
}

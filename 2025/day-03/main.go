package main

import (
	"fmt"
	"strconv"

	"github.com/mzietara/advent-of-code/util"
)

func solution(part int) {
	banks := banks()
	sum := int64(0)

	numSize := 2
	if part == 2 {
		numSize = 12
	}

	for _, bank := range banks {
		sum += highestKDigit(bank, numSize)
	}
	fmt.Println(sum)
}

func highestKDigit(s string, k int) int64 {
	if k <= 0 || k > len(s) {
		return int64(-1)
	}

	result := make([]byte, 0, k)
	startIndex := 0

	for i := 0; i < k; i++ {
		remaining := k - i - 1

		endIndex := len(s) - remaining

		maxDigit := byte('0') - 1
		maxIndex := startIndex

		for j := startIndex; j < endIndex; j++ {
			if s[j] > maxDigit {
				maxDigit = s[j]
				maxIndex = j
			}
		}

		result = append(result, s[maxIndex])

		startIndex = maxIndex + 1
	}

	num, err := strconv.ParseInt(string(result), 10, 64)
	if err != nil {
		return int64(-1)
	}

	return num
}

func banks() []string {
	banks := []string{}
	util.ProcessInputFile(func(_ int, line string) {
		banks = append(banks, line)
	})
	return banks
}

func main() {
	defer util.Timer()()
	solution(1)
	solution(2)
}

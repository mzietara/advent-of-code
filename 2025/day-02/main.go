package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

func solution(partNum int) {
	idsToCheck := processInputFile()
	// fmt.Println(idsToCheck)
	sum := 0

	for _, id := range idsToCheck {
		if id < 10 {
			continue
		} else {
			if partNum == 1 && len(strconv.Itoa(id))%2 == 1 {
				continue
			}

			// TODO continue the case for even number of digits
			// see how many digits there are and find the divisors of that number
			numDigits := len(strconv.Itoa(id))
			divisors := []int{}
			for i := 1; i <= numDigits/2; i++ {
				if numDigits%i == 0 {
					divisors = append(divisors, i)
				}
			}

			digits := strings.Split(strconv.Itoa(id), "")
			for _, divisor := range divisors {
				// split the number into parts of length divisor and see if all parts are the same
				if partNum == 1 && divisor != numDigits/2 {
					continue
				}
				if divisor > numDigits/2 {
					continue
				}

				firstPart := strings.Join(digits[0:divisor], "")

				repeating := true
				for i := divisor; i < numDigits; i += divisor {
					part := strings.Join(digits[i:i+divisor], "")
					if part != firstPart {
						repeating = false
						break
					}
				}
				if repeating {
					fmt.Printf("invalid id: %d\n", id)
					sum += id
					break
				}
			}
		}
	}
	fmt.Println(sum)

}

func processInputFile() []int {
	idsToCheck := []int{}
	util.ProcessInputFile(func(_ int, line string) {
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			bounds := strings.Split(r, "-")
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			if start > end {
				continue
			}
			for i := start; i <= end; i++ {
				idsToCheck = append(idsToCheck, i)
			}
		}

	})

	return idsToCheck
}

func main() {
	defer util.Timer()()
	// solution(1)
	solution(2)
}

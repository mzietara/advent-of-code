package main

import (
	"fmt"

	"github.com/mzietara/advent-of-code/util"
)

type wordSearch [][]rune

var WORD = []rune("XMAS")

func main() {
	//part1()
	part2()
}

func part1() {
	lines := wordSearchArray()
	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == WORD[0] {
				sum += countWordsFromFirstLetter(lines, i, j)
			}
		}
	}
	println(sum)
}

func wordSearchArray() wordSearch {
	var lines [][]rune
	util.ProcessFile("input.txt", func(_ int, line string) {
		lines = append(lines, []rune(line))
	})
	return lines
}

// countWordsFromFirstLetter counts the number of words that can be formed
// starting from the first letter of the word.
// i and j represent the index of the first letter.
func countWordsFromFirstLetter(ws wordSearch, i, j int) int {
	sum := 0
	if ws[i][j] != WORD[0] {
		fmt.Println("Error: first letter is not the first letter of the word")
		return 0
	}

	// check right
	if j+len(WORD) <= len(ws[i]) {
		found := true
		for k := range WORD {
			if ws[i][j+k] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check down
	if i+len(WORD) <= len(ws) {
		found := true
		for k := range WORD {
			if ws[i+k][j] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check down-right
	if i+len(WORD) <= len(ws) && j+len(WORD) <= len(ws[i]) {
		found := true
		for k := range WORD {
			if ws[i+k][j+k] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check down-left
	if i+len(WORD) <= len(ws) && j-len(WORD)+1 >= 0 {
		found := true
		for k := range WORD {
			if ws[i+k][j-k] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check up-right
	if i-len(WORD)+1 >= 0 && j+len(WORD) <= len(ws[i]) {
		found := true
		for k := range WORD {
			if ws[i-k][j+k] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check up-left
	if i-len(WORD)+1 >= 0 && j-len(WORD)+1 >= 0 {
		found := true
		for k := range WORD {
			if ws[i-k][j-k] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check up
	if i-len(WORD)+1 >= 0 {
		found := true
		for k := range WORD {
			if ws[i-k][j] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	// check left
	if j-len(WORD)+1 >= 0 {
		found := true
		for k := range WORD {
			if ws[i][j-k] != WORD[k] {
				found = false
				break
			}
		}
		if found {
			sum++
		}
	}

	return sum
}

func part2() {
	lines := wordSearchArray()
	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'A' {
				if isMatchFromMiddle(lines, i, j) {
					sum++
				}
			}
		}
	}
	println(sum)

}

func isMatchFromMiddle(ws wordSearch, i, j int) bool {
	if ws[i][j] != 'A' {
		fmt.Println("Error: middle letter is not the middle letter of the word")
		return false
	}

	// check if pattern can fit
	if !(i > 0 && i < len(ws)-1 && j > 0 && j < len(ws[i])-1) {
		return false
	}

	// see what top left corner is
	if ws[i-1][j-1] == 'M' {
		if (ws[i+1][j-1] == 'M' && ws[i+1][j+1] == 'S' && ws[i-1][j+1] == 'S') ||
			(ws[i+1][j-1] == 'S' && ws[i+1][j+1] == 'S' && ws[i-1][j+1] == 'M') {
			return true
		}

	} else if ws[i-1][j-1] == 'S' {
		if (ws[i+1][j-1] == 'M' && ws[i+1][j+1] == 'M' && ws[i-1][j+1] == 'S') ||
			(ws[i+1][j-1] == 'S' && ws[i+1][j+1] == 'M' && ws[i-1][j+1] == 'M') {
			return true
		}
	}

	return false
}

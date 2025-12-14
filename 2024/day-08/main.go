package main

import (
	"fmt"

	"github.com/mzietara/advent-of-code/util"
)

func main() {
	util.Timer()()
	part1()
	part2()
}

func part1() {
	antennas := map[rune][][2]int{}
	gridLength := 0
	gridHeight := 0
	i := 0
	util.ProcessInputFile(func(j int, line string) {
		if gridLength == 0 {
			gridLength = len(line)
		}
		for j, r := range line {
			if r != '.' {
				antennas[r] = append(antennas[r], [2]int{i, j})
			}
		}
		i++
	})
	gridHeight = i

	//gather antinodes
	antinodes := map[[2]int]struct{}{}
	for _, v := range antennas {
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				a := v[i][0] - v[j][0]
				b := v[i][1] - v[j][1]

				if a+v[i][0] >= 0 && a+v[i][0] < gridHeight && b+v[i][1] >= 0 && b+v[i][1] < gridLength {
					antinodes[[2]int{a + v[i][0], b + v[i][1]}] = struct{}{}
				}
				if -a+v[j][0] >= 0 && -a+v[j][0] < gridHeight && -b+v[j][1] >= 0 && -b+v[j][1] < gridLength {
					antinodes[[2]int{-a + v[j][0], -b + v[j][1]}] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(antinodes))

}

func part2() {
	antennas := map[rune][][2]int{}
	gridLength := 0
	gridHeight := 0
	i := 0
	util.ProcessInputFile(func(j int, line string) {
		if gridLength == 0 {
			gridLength = len(line)
		}
		for j, r := range line {
			if r != '.' {
				antennas[r] = append(antennas[r], [2]int{i, j})
			}
		}
		i++
	})
	gridHeight = i

	//gather antinodes
	antinodes := map[[2]int]struct{}{}
	for _, v := range antennas {
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				a := v[i][0] - v[j][0]
				b := v[i][1] - v[j][1]

				curPosition := [2]int{v[i][0], v[i][1]}
				for curPosition[0] >= 0 && curPosition[0] < gridHeight && curPosition[1] >= 0 && curPosition[1] < gridLength {
					antinodes[curPosition] = struct{}{}
					curPosition[0] += a
					curPosition[1] += b
				}

				curPosition = [2]int{v[j][0], v[j][1]}
				for curPosition[0] >= 0 && curPosition[0] < gridHeight && curPosition[1] >= 0 && curPosition[1] < gridLength {
					antinodes[curPosition] = struct{}{}
					curPosition[0] -= a
					curPosition[1] -= b
				}

			}
		}
	}
	fmt.Println(len(antinodes))

}

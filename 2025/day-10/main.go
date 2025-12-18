package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

type machine struct {
	lightsState lights
	buttons     [][]int
	voltages    []int
}

type lights []bool

type lightState struct {
	lights  lights
	presses int
}

func main() {
	defer util.Timer()()
	part1()
	part2()
}

func part2() {
	machines := processInput()
	sum := 0
	for _, m := range machines {
		sum += solveLinearSystem(m.buttons, m.voltages)
	}
	fmt.Println(sum)
}

func solveLinearSystem(buttons [][]int, target []int) int {
	n := len(target)
	m := len(buttons)

	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, m+1)
		aug[i][m] = float64(target[i])
	}

	for j, button := range buttons {
		for _, pos := range button {
			if pos < n {
				aug[pos][j] = 1
			}
		}
	}

	pivotRow := 0
	pivotCols := make([]int, 0, m)
	pivotRowForCol := make([]int, m)
	for i := range pivotRowForCol {
		pivotRowForCol[i] = -1
	}

	for col := 0; col < m && pivotRow < n; col++ {
		maxRow := pivotRow
		for row := pivotRow + 1; row < n; row++ {
			if abs64(aug[row][col]) > abs64(aug[maxRow][col]) {
				maxRow = row
			}
		}

		if abs64(aug[maxRow][col]) < 1e-10 {
			continue
		}

		aug[pivotRow], aug[maxRow] = aug[maxRow], aug[pivotRow]

		scale := aug[pivotRow][col]
		for k := col; k <= m; k++ {
			aug[pivotRow][k] /= scale
		}

		for row := 0; row < n; row++ {
			if row != pivotRow && abs64(aug[row][col]) > 1e-10 {
				factor := aug[row][col]
				for k := col; k <= m; k++ {
					aug[row][k] -= factor * aug[pivotRow][k]
				}
			}
		}

		pivotCols = append(pivotCols, col)
		pivotRowForCol[col] = pivotRow
		pivotRow++
	}

	for row := pivotRow; row < n; row++ {
		if abs64(aug[row][m]) > 1e-10 {
			return -1
		}
	}

	freeVars := []int{}
	for col := 0; col < m; col++ {
		if pivotRowForCol[col] == -1 {
			freeVars = append(freeVars, col)
		}
	}

	if len(freeVars) == 0 {
		solution := make([]int, m)
		for i, col := range pivotCols {
			val := aug[i][m]
			rounded := int(val + 0.5)
			if abs64(val-float64(rounded)) > 1e-9 || rounded < 0 {
				return -1
			}
			solution[col] = rounded
		}
		total := 0
		for _, v := range solution {
			total += v
		}
		return total
	}

	maxFreeVal := 0
	for _, t := range target {
		if t > maxFreeVal {
			maxFreeVal = t
		}
	}

	return searchFreeVariables(aug, m, pivotCols, freeVars, make([]int, len(freeVars)), 0, maxFreeVal)
}

func searchFreeVariables(aug [][]float64, m int, pivotCols []int, freeVars []int, freeVals []int, idx int, maxVal int) int {
	if idx == len(freeVars) {
		solution := make([]int, m)
		for i, col := range freeVars {
			solution[col] = freeVals[i]
		}

		for i, col := range pivotCols {
			val := aug[i][m]
			for _, freeCol := range freeVars {
				val -= aug[i][freeCol] * float64(solution[freeCol])
			}
			rounded := int(val + 0.5)
			if abs64(val-float64(rounded)) > 1e-9 || rounded < 0 {
				return -1
			}
			solution[col] = rounded
		}

		total := 0
		for _, v := range solution {
			total += v
		}
		return total
	}

	best := -1
	for v := 0; v <= maxVal; v++ {
		freeVals[idx] = v
		result := searchFreeVariables(aug, m, pivotCols, freeVars, freeVals, idx+1, maxVal)
		if result != -1 && (best == -1 || result < best) {
			best = result
		}
	}
	return best
}

func abs64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	machines := processInput()
	sum := 0
	for _, m := range machines {
		sum += minButtonPressesPart1(m)
	}
	fmt.Println(sum)
}

func minButtonPressesPart1(m machine) int {
	initialState := make(lights, len(m.lightsState))
	queue := []lightState{{lights: initialState, presses: 0}}
	visited := make(map[string]bool)
	visited[initialState.String()] = true

	for len(queue) > 0 {
		curState := queue[0]
		queue = queue[1:]

		for _, b := range m.buttons {
			newLights := pressLightsButton(curState.lights, b)
			newPresses := curState.presses + 1
			if newLights.Equals(m.lightsState) {
				return newPresses
			}
			if !visited[newLights.String()] {
				visited[newLights.String()] = true
				queue = append(queue, lightState{lights: newLights, presses: newPresses})
			}
		}
	}
	return -1
}

func pressLightsButton(cur lights, button []int) lights {
	newLights := make(lights, len(cur))
	copy(newLights, cur)
	for _, idx := range button {
		newLights[idx] = !newLights[idx]
	}
	return newLights
}

func (l lights) String() string {
	str := ""
	for _, light := range l {
		if light {
			str += "#"
		} else {
			str += "."
		}
	}
	return str
}

func (l lights) Equals(other lights) bool {
	if len(l) != len(other) {
		return false
	}
	for i := range l {
		if l[i] != other[i] {
			return false
		}
	}
	return true
}

func lightsFromString(s string) lights {
	l := make(lights, len(s))
	for i := 0; i < len(s); i++ {
		l[i] = s[i] == '#'
	}
	return l
}

func processInput() []machine {
	machines := []machine{}
	util.ProcessInputFile(func(i int, line string) {
		buttons := [][]int{}
		voltages := []int{}
		split := strings.Fields(line)

		lightsStr := strings.Trim(split[0], "[]")
		finalState := lightsFromString(lightsStr)

		for i := 1; i < len(split)-1; i++ {
			withoutParens := strings.Trim(split[i], "()")
			buttonValsStr := strings.Split(withoutParens, ",")
			button := []int{}
			for _, valStr := range buttonValsStr {
				val, _ := strconv.Atoi(valStr)
				button = append(button, val)
			}
			buttons = append(buttons, button)
		}

		voltagesRaw := strings.Trim(split[len(split)-1], "{}")
		voltageStrs := strings.Split(voltagesRaw, ",")
		for _, voltageStr := range voltageStrs {
			voltage, _ := strconv.Atoi(voltageStr)
			voltages = append(voltages, voltage)
		}

		machines = append(machines, machine{
			lightsState: finalState,
			buttons:     buttons,
			voltages:    voltages,
		})
	})
	return machines
}

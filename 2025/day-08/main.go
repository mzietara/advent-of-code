package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

type pair struct {
	a, b     point
	distance float64
}
type point [3]int

func solution(questionPart int, numTurns int) {
	points := processInput()
	pairs := make([]pair, 0)
	pointsSlice := make([]point, 0, len(points))
	for p := range points {
		pointsSlice = append(pointsSlice, p)
	}

	// gather list of all minimum distances
	for i := 0; i < len(pointsSlice); i++ {
		for j := i + 1; j < len(pointsSlice); j++ {
			d := distance(pointsSlice[i], pointsSlice[j])
			pairs = append(pairs, pair{a: pointsSlice[i], b: pointsSlice[j], distance: d})
		}
	}
	// sort pairs by distance
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})

	circuits := make([]map[point]struct{}, 0)

	for i := 0; i < numTurns || questionPart == 2; i++ {
		aInCircuit := points[pairs[i].a] != -1
		bInCircuit := points[pairs[i].b] != -1

		if !aInCircuit && !bInCircuit {
			// create new circuit
			points[pairs[i].a] = len(circuits)
			points[pairs[i].b] = len(circuits)
			circuits = append(circuits, map[point]struct{}{pairs[i].a: {}, pairs[i].b: {}})
		} else if aInCircuit && bInCircuit && points[pairs[i].a] != points[pairs[i].b] {
			// merge the circuits
			minCircuitIndex := min(points[pairs[i].a], points[pairs[i].b])
			maxCircuitIndex := max(points[pairs[i].a], points[pairs[i].b])
			for node := range circuits[maxCircuitIndex] {
				points[node] = minCircuitIndex
				circuits[minCircuitIndex][node] = struct{}{}
			}
			circuits = append(circuits[:maxCircuitIndex], circuits[maxCircuitIndex+1:]...)
			for node := range points {
				if points[node] > maxCircuitIndex {
					points[node]--
				}
			}
		} else if aInCircuit {
			// add b to a's circuit
			points[pairs[i].b] = points[pairs[i].a]
			circuits[points[pairs[i].a]][pairs[i].b] = struct{}{}
		} else if bInCircuit {
			// add a to b's circuit
			points[pairs[i].a] = points[pairs[i].b]
			circuits[points[pairs[i].b]][pairs[i].a] = struct{}{}
		}
		if questionPart == 2 {
			if len(circuits) == 1 && len(circuits[0]) == len(points) {
				println(pairs[i].a[0] * pairs[i].b[0])
				break
			}
		}
	}

	if questionPart == 2 {
		return
	}
	// sort circuits by size, largest first
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	sum := len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	println(sum)
}

func processInput() map[point]int {
	points := map[point]int{}

	util.ProcessInputFile(func(i int, line string) {
		lineSplit := strings.Split(line, ",")
		x, _ := strconv.Atoi(lineSplit[0])
		y, _ := strconv.Atoi(lineSplit[1])
		z, _ := strconv.Atoi(lineSplit[2])
		points[point{x, y, z}] = -1
	})
	return points
}

func distance(p1, p2 [3]int) float64 {
	dx := float64(p2[0] - p1[0])
	dy := float64(p2[1] - p1[1])
	dz := float64(p2[2] - p1[2])
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	defer util.Timer()()
	// solution(1, 1000)
	solution(2, -1)
}

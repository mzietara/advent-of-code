package main

import (
	"container/heap"
	"fmt"

	. "github.com/mzietara/advent-of-code/util"
)

var grid Matrix[object]
var data struct {
	Start       Coord
	End         Coord
	LowestScore map[dirCoord]int
	Visiting    map[Coord]struct{}
}

type dirCoord struct {
	pos Coord
	dir direction
}

const deadEnd = -1

type direction Coord

var (
	up    direction = direction(Coord{X: 0, Y: -1})
	down  direction = direction(Coord{X: 0, Y: 1})
	left  direction = direction(Coord{X: -1, Y: 0})
	right direction = direction(Coord{X: 1, Y: 0})
)

func (d direction) IsAdjacent(other direction) bool {
	if d == other {
		return false
	}
	return d == up && other != down || d == down && other != up || d == left && other != right || d == right && other != left
}

type object rune

const (
	wall  object = '#'
	empty object = '.'
	start object = 'S'
	end   object = 'E'
)

type PriorityQueue []*Node

type Node struct {
	value    dirCoord
	priority int
	index    int
	path     []Coord
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func main() {
	defer Timer()()
	part1()
	part2()
}

func part1() {
	processInput()
	fmt.Println(FindLowestScoreData(data.Start, right).lowestScore)
}

type pathInfo struct {
	lowestScore      int
	nodesOnBestPaths int
}

func FindLowestScoreData(start Coord, initialDir direction) (info pathInfo) {
	dist := make(map[dirCoord]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	startDirCoord := dirCoord{start, initialDir}
	dist[startDirCoord] = 0
	heap.Push(&pq, &Node{value: startDirCoord, priority: 0, path: []Coord{start}})
	sizeToNodes := make(map[int][]Coord)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		currentPos := current.value.pos
		currentDir := current.value.dir

		if currentPos == data.End {
			//fmt.Println("end", current.priority, info.lowestScore)
			if info.lowestScore == 0 || current.priority <= info.lowestScore {
				info.lowestScore = current.priority
				sizeToNodes[current.priority] = append(sizeToNodes[current.priority], current.path...)
			}
			continue
		}

		// Move in the current direction
		nextPos := currentPos.Add(Coord(currentDir))
		if nextPos.X >= 0 && nextPos.X < len(grid[0]) && nextPos.Y >= 0 && nextPos.Y < len(grid) {
			nextVal := grid.Get(nextPos)
			if nextVal != wall {
				nextDirCoord := dirCoord{nextPos, currentDir}
				alt := dist[dirCoord{currentPos, currentDir}] + 1
				if v, ok := dist[nextDirCoord]; !ok || alt <= v {
					dist[nextDirCoord] = alt
					nPath := make([]Coord, len(current.path))
					copy(nPath, current.path)
					heap.Push(&pq, &Node{value: nextDirCoord, priority: alt, path: append(nPath, current.value.pos)})
				}
			}
		}

		// Try all other directions
		for _, d := range []direction{up, down, left, right} {
			if currentDir.IsAdjacent(d) {
				nextPos := currentPos.Add(Coord(d))
				if nextPos.X < 0 || nextPos.X >= len(grid[0]) || nextPos.Y < 0 || nextPos.Y >= len(grid) || grid.Get(nextPos) == wall {
					continue
				}

				moveAmount := 1001
				if _, ok := data.Visiting[nextPos]; ok {
					continue
				}
				nextDirCoord := dirCoord{nextPos, d}
				alt := dist[dirCoord{currentPos, currentDir}] + moveAmount
				if v, ok := dist[nextDirCoord]; !ok || alt <= v {
					dist[nextDirCoord] = alt
					nPath := make([]Coord, len(current.path))
					copy(nPath, current.path)
					heap.Push(&pq, &Node{value: nextDirCoord, priority: alt, path: append(nPath, current.value.pos)})
				}
			}
		}
	}
	countMap := make(map[Coord]bool)
	for _, index := range sizeToNodes[info.lowestScore] {
		countMap[index] = true
	}
	info.nodesOnBestPaths = len(countMap) + 1 //add the end point

	//print grid
	for i, row := range grid {
		for j, val := range row {
			if countMap[Coord{X: j, Y: i}] {
				fmt.Print("O")
			} else {
				fmt.Print(string(val))
			}
		}
		fmt.Println()
	}

	return
}

func processInput() {
	resetData()
	ProcessInputFile(func(i int, line string) {
		objects := []object(line)
		for j, o := range objects {
			if o == start {
				data.Start = Coord{X: j, Y: i}
			} else if o == end {
				data.End = Coord{X: j, Y: i}
			}
		}
		grid = append(grid, []object(line))
	})
}

func resetData() {
	grid = Matrix[object]{}
	data.LowestScore = make(map[dirCoord]int)
	data.Visiting = make(map[Coord]struct{})
}

func part2() {
	processInput()
	allPaths := FindLowestScoreData(data.Start, right).nodesOnBestPaths

	fmt.Println(allPaths)
}

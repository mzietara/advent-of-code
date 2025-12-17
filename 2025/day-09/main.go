package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mzietara/advent-of-code/util"
)

func part2() {
	points := processInput()
	largestArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]

			if p1[0] == p2[0] || p1[1] == p2[1] {
				continue
			}

			width := abs(p2[0]-p1[0]) + 1
			height := abs(p2[1]-p1[1]) + 1
			area := width * height

			if area <= largestArea {
				continue
			}

			minX, maxX := min(p1[0], p2[0]), max(p1[0], p2[0])
			minY, maxY := min(p1[1], p2[1]), max(p1[1], p2[1])

			if isRectangleValid(minX, maxX, minY, maxY, points) {
				largestArea = area
			}
		}
	}

	fmt.Println(largestArea)
}

// Check if rectangle is entirely inside the polygon
func isRectangleValid(minX, maxX, minY, maxY int, points [][2]int) bool {
	// Check the 4 corners are inside or on the polygon
	corners := [][2]int{
		{minX, minY},
		{minX, maxY},
		{maxX, minY},
		{maxX, maxY},
	}

	for _, c := range corners {
		if !isPointInsideOrOnPolygon(c[0], c[1], points) {
			return false
		}
	}

	// For each edge of the rectangle, check it doesn't get crossed by a polygon edge
	// that would put part of the rectangle outside
	n := len(points)

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		px1, py1 := points[i][0], points[i][1]
		px2, py2 := points[j][0], points[j][1]

		// Check if this polygon edge crosses any of the 4 rectangle edges
		// Top edge: y = maxY, x from minX to maxX
		if segmentsCross(minX, maxY, maxX, maxY, px1, py1, px2, py2) {
			return false
		}
		// Bottom edge: y = minY, x from minX to maxX
		if segmentsCross(minX, minY, maxX, minY, px1, py1, px2, py2) {
			return false
		}
		// Left edge: x = minX, y from minY to maxY
		if segmentsCross(minX, minY, minX, maxY, px1, py1, px2, py2) {
			return false
		}
		// Right edge: x = maxX, y from minY to maxY
		if segmentsCross(maxX, minY, maxX, maxY, px1, py1, px2, py2) {
			return false
		}
	}

	return true
}

// Check if two axis-aligned segments cross in the interior
func segmentsCross(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) bool {
	if ay1 == ay2 { // Segment A is horizontal
		if by1 == by2 { // Segment B is also horizontal (parallel)
			return false
		}
		// Segment B is vertical
		bx := bx1
		bMinY, bMaxY := min(by1, by2), max(by1, by2)
		aMinX, aMaxX := min(ax1, ax2), max(ax1, ax2)

		return aMinX < bx && bx < aMaxX && bMinY < ay1 && ay1 < bMaxY
	} else { // Segment A is vertical
		if bx1 == bx2 { // Segment B is also vertical (parallel)
			return false
		}
		// Segment B is horizontal
		ax := ax1
		aMinY, aMaxY := min(ay1, ay2), max(ay1, ay2)
		bMinX, bMaxX := min(bx1, bx2), max(bx1, bx2)

		return bMinX < ax && ax < bMaxX && aMinY < by1 && by1 < aMaxY
	}
}

// Check if point is inside or on polygon using ray casting
func isPointInsideOrOnPolygon(x, y int, polygon [][2]int) bool {
	n := len(polygon)
	inside := false

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		xi, yi := polygon[i][0], polygon[i][1]
		xj, yj := polygon[j][0], polygon[j][1]

		// Check if point is on this edge
		if yi == yj && yi == y && min(xi, xj) <= x && x <= max(xi, xj) {
			return true
		}
		if xi == xj && xi == x && min(yi, yj) <= y && y <= max(yi, yj) {
			return true
		}

		if ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
	}

	return inside
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	points := processInput()
	largestArea := 0

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			xDiff := abs(points[i][0]-points[j][0]) + 1
			yDiff := abs(points[i][1]-points[j][1]) + 1
			area := xDiff * yDiff
			if area > largestArea {
				largestArea = area
			}
		}
	}
	println(largestArea)
}

func processInput() [][2]int {
	points := [][2]int{}
	util.ProcessInputFile(func(i int, line string) {
		lineSplit := strings.Split(line, ",")
		x, _ := strconv.Atoi(lineSplit[0])
		y, _ := strconv.Atoi(lineSplit[1])
		points = append(points, [2]int{x, y})
	})
	return points
}

func main() {
	defer util.Timer()()
	part2()
}

package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type PuzzlePart int

const (
	Part1 PuzzlePart = iota
	Part2
)

// Coord represents a coordinate with x and y values.
type Coord struct {
	X, Y int
}

func (c *Coord) Add(other Coord) Coord {
	return Coord{X: c.X + other.X, Y: c.Y + other.Y}
}

type Matrix[T any] [][]T

func (m Matrix[T]) Get(coord Coord) T {
	return m[coord.Y][coord.X]
}

func NewMatrix[T any](rows, cols int) Matrix[T] {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}
	return matrix
}

func ProcessInputMatrixInt(splitStr string) Matrix[int] {
	return ProcessInputMatrix(splitStr, func(s string) int {
		result, _ := strconv.Atoi(s)
		return result
	})
}

func (m Matrix[T]) Iterate(f func(y, x int, value T)) {
	for y, row := range m {
		for x, cell := range row {
			f(y, x, cell)
		}
	}
}

// ProcessInputMatrix processes the input file and converts each character to a value of type T using the provided converter function.
func ProcessInputMatrix[T any](splitStr string, fromString func(s string) T) Matrix[T] {
	matrix := make([][]T, 0)
	ProcessInputFile(func(i int, line string) {
		row := make([]T, 0)
		split := strings.Split(line, splitStr)
		for _, s := range split {
			row = append(row, fromString(s))
		}
		matrix = append(matrix, row)
	})
	return matrix
}

func PrintMatrix[T any](matrix [][]T) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

// ProcessInputFile processes each line of the input file using the provided processLine function.
func ProcessInputFile(processLine func(int, string)) {
	ProcessFile("input.txt", processLine)
}

// ProcessFile processes each line of the specified file using the provided processLine function.
func ProcessFile(filename string, processLine func(int, string)) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		processLine(i, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// Timer returns a function that prints the elapsed time since the Timer was created.
func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Executed in %dms\n", time.Since(start).Milliseconds())
	}
}

// Set represents a generic set of elements of type T.
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates a new Set.
func NewSet[T comparable]() *Set[T] {
	result := &Set[T]{elements: make(map[T]struct{})}
	return result
}

// Add adds one or more elements to the set.
func (s *Set[T]) Add(element ...T) {
	for _, e := range element {
		s.elements[e] = struct{}{}
	}
}

func (s *Set[T]) AddSet(set Set[T]) {
	for e := range set.elements {
		s.elements[e] = struct{}{}
	}
}

// Remove removes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Contains checks if the set contains an element.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s.elements)
}

// Elements returns a slice of all elements in the set.
func (s *Set[T]) Elements() []T {
	keys := make([]T, 0, len(s.elements))
	for k := range s.elements {
		keys = append(keys, k)
	}
	return keys
}

func StringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", s)
		panic(err)
	}
	return result
}

// RemoveElement removes the first occurrence of a specified element from a slice.
// It takes a slice of any comparable type and a value to be removed from the slice.
// If the value is found, it returns a new slice with the value removed.
// If the value is not found, it returns the original slice.
//
// T is a type parameter that must be comparable.
//
// Parameters:
//   - slice: The slice from which the element should be removed.
//   - value: The value to be removed from the slice.
//
// Returns:
//
//	A new slice with the first occurrence of the value removed, or the original slice if the value is not found.
func RemoveElement[T comparable](slice []T, value T) []T {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// RemoveElements removes all occurrences of specified elements from a slice.
// It takes a slice of any comparable type and one or more values to be removed from the slice.
// It returns a new slice with all occurrences of the values removed.
//
// T is a type parameter that must be comparable.
//
// Parameters:
//   - slice: The slice from which the elements should be removed.
//   - values: The values to be removed from the slice.
//
// Returns:
//
//	A new slice with all occurrences of the values removed.
func RemoveElements[T comparable](slice []T, values ...T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		found := false
		for _, value := range values {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			result = append(result, v)
		}
	}
	return result
}

// RemoveAllElements removes all occurrences of a specified element from a slice.
// It takes a slice of any comparable type and a value to be removed from the slice.
// It returns a new slice with all occurrences of the value removed.
//
// T is a type parameter that must be comparable.
//
// Parameters:
//   - slice: The slice from which the element should be removed.
//   - value: The value to be removed from the slice.
//
// Returns:
//
//	A new slice with all occurrences of the value removed.
func RemoveAllElements[T comparable](slice []T, value T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// RemoveElementAt removes the element at the specified index from a slice.
// It takes a slice of any type and an index to be removed from the slice.
// If the index is within the bounds of the slice, it returns a new slice with the element removed.
// If the index is out of bounds, it returns the original slice.
//
// T is a type parameter that can be any type.
//
// Parameters:
//   - slice: The slice from which the element should be removed.
//   - index: The index of the element to be removed from the slice.
//
// Returns:
//
//	A new slice with the element at the specified index removed, or the original slice if the index is out of bounds.
func RemoveElementAt[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// ContainsElement checks if a slice contains a specified element
// It takes a slice of any comparable type and a value to be checked.
// If the value is found in the slice, it returns true.
// If the value is not found in the slice, it returns false.
//
// T is a type parameter that must be comparable.
//
// Parameters:
//   - slice: The slice to be checked for the value.
//   - value: The value to be checked in the slice.
//
// Returns:
//
//	True if the value is found in the slice, false otherwise.
func ContainsElement[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

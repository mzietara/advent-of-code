package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	. "github.com/mzietara/advent-of-code/util"
)

type Computer struct {
	regA, regB, regC int
	program          []int
}

func main() {
	c := processInput()
	part1(c)
	part2(c.program)
}

func part2(program []int) {
	regA := 0
	for i := len(program) - 1; i >= 0; i-- {
		regA <<= 3
		output := runProgram(program, regA)
		for !slices.Equal(output, program[i:]) {
			regA++
			output = runProgram(program, regA)
		}
	}
	fmt.Println(regA)
}

func part1(c Computer) {
	fmt.Println(runProgram(c.program, c.regA))
}

func runProgram(program []int, regA int) (res []int) {
	var instructionPointer int
	c := Computer{
		regA:    regA,
		regB:    0,
		regC:    0,
		program: program,
	}

	for instructionPointer < len(program)-1 {
		operand := program[instructionPointer+1]

		switch operator := program[instructionPointer]; operator {
		case 0: // adv
			c.regA >>= getValue(operand, c)
		case 1: // bxl
			c.regB ^= operand
		case 2: // bst
			c.regB = getValue(operand, c) & 7
		case 3: // jnz
			if c.regA != 0 {
				instructionPointer = int(operand)
				continue
			}
		case 4: // bxc
			c.regB ^= c.regC
		case 5: // out
			val := getValue(operand, c) & 7
			res = append(res, val)
		case 6: // bdv
			c.regB = c.regA >> getValue(operand, c)
		case 7: // cdv
			c.regC = c.regA >> getValue(operand, c)
		}
		instructionPointer += 2
	}

	return
}

func getValue(operand int, c Computer) (value int) {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	}
	return -1
}

func processInput() Computer {
	c := Computer{}
	ProcessInputFile(func(_ int, line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		s := strings.Split(line, ": ")

		switch s[0] {
		case "Register A":
			c.regA, _ = strconv.Atoi(s[1])
		case "Register B":
			c.regB, _ = strconv.Atoi(s[1])
		case "Register C":
			c.regC, _ = strconv.Atoi(s[1])
		case "Program":
			program := strings.Split(s[1], ",")
			for _, instrRaw := range program {
				instr, _ := strconv.Atoi(instrRaw)
				c.program = append(c.program, instr)
			}
		}

	})
	return c
}

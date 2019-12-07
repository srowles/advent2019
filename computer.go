package advent2019

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// IntcodeComputer implements an AOC 2019 intcode computer
type IntcodeComputer struct {
	pointer  int
	program  []int
	input    []int
	inputIdx int
	output   int
}

// Incode Computer Instructions
var (
	Add    = 1
	Mul    = 2
	Input  = 3
	Output = 4
	JIT    = 5
	JIF    = 6
	LT     = 7
	EQ     = 8
	Halt   = 99
)

// Run runs the program until it halts
func (i *IntcodeComputer) Run() error {
	for {
		var codes []int
		codes = intToSlice(i.program[i.pointer], codes)
		for k := 0; k <= 3; k++ {
			codes = append(codes, 0)
		}

		opcode := codes[0]
		opcode += codes[1] * 10

		switch opcode {
		case Add:
			i.program[i.program[i.pointer+3]] = i.getVal(1, codes) + i.getVal(2, codes)
			i.pointer += 4
		case Mul:
			i.program[i.program[i.pointer+3]] = i.getVal(1, codes) * i.getVal(2, codes)
			i.pointer += 4
		case Input:
			if codes[2] == 0 {
				i.program[i.program[i.pointer+1]] = i.input[i.inputIdx]
			} else {
				i.program[i.pointer+1] = i.input[i.inputIdx]
			}
			i.inputIdx++
			i.pointer += 2
		case Output:
			if codes[2] == 0 {
				i.output = i.program[i.program[i.pointer+1]]
				// fmt.Println("Output:", i.program[i.program[i.pointer+1]])
			} else {
				i.output = i.program[i.pointer+1]
				// fmt.Println("Output:", i.program[i.pointer+1])
			}
			i.pointer += 2
		case JIT:
			first := i.getVal(1, codes)
			second := i.getVal(2, codes)

			if first != 0 {
				i.pointer = second
			} else {
				i.pointer += 3
			}
		case JIF:
			first := i.getVal(1, codes)
			second := i.getVal(2, codes)

			if first == 0 {
				i.pointer = second
			} else {
				i.pointer += 3
			}
		case LT:
			first := i.getVal(1, codes)
			second := i.getVal(2, codes)

			if first < second {
				i.program[i.program[i.pointer+3]] = 1
			} else {
				i.program[i.program[i.pointer+3]] = 0
			}
			i.pointer += 4
		case EQ:
			first := i.getVal(1, codes)
			second := i.getVal(2, codes)

			if first == second {
				i.program[i.program[i.pointer+3]] = 1
			} else {
				i.program[i.program[i.pointer+3]] = 0
			}
			i.pointer += 4
		case Halt:
			return nil
		default:
			return fmt.Errorf("Invalid opcode: %v found at pointer: %v", i.program[i.pointer], i.pointer)
		}
	}
}

func (i *IntcodeComputer) getVal(idx int, codes []int) int {
	if codes[idx+1] == 0 {
		return i.program[i.program[i.pointer+idx]]
	}
	return i.program[i.pointer+idx]
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append(sequence, i)
		return intToSlice(n/10, sequence)
	}
	return sequence
}

// Peek allows inspecting the program
func (i *IntcodeComputer) Peek(pos int) int {
	return i.program[pos]
}

// Poke allows inspecting the program
func (i *IntcodeComputer) Poke(pos, val int) {
	i.program[pos] = val
}

// Input sets the input to the program
func (i *IntcodeComputer) Input(val int) {
	i.input = append(i.input, val)
}

// Output returns the current output value
func (i *IntcodeComputer) Output() int {
	return i.output
}

// CreateIntcodeComputerFromFile creates an intcode computer from the supplied file
func CreateIntcodeComputerFromFile(fileName string) (*IntcodeComputer, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file '%s' with error: %v", fileName, err)
	}

	intStrs := strings.Split(strings.TrimSpace(string(data)), ",")

	var ints []int
	for _, i := range intStrs {
		val, err := strconv.Atoi(i)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse '%v' as an int: %v", i, err)
		}
		ints = append(ints, val)
	}

	return createIntcodeComputer(ints)
}

// CreateIntcodeComputer creates an intcode computer from the supplied file
func createIntcodeComputer(ints []int) (*IntcodeComputer, error) {
	computer := &IntcodeComputer{
		program: ints,
	}

	return computer, nil
}

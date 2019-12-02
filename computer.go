package advent2019

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// IntcodeComputer implements an AOC 2019 intcode computer
type IntcodeComputer struct {
	pointer int
	program []int
}

// Incode Computer Instructions
var (
	Add  = 1
	Mul  = 2
	Halt = 99
)

// Run runs the program until it halts
func (i *IntcodeComputer) Run() error {
	for {
		switch i.program[i.pointer] {
		case Add:
			i.program[i.program[i.pointer+3]] = i.program[i.program[i.pointer+1]] + i.program[i.program[i.pointer+2]]
			i.pointer += 4
		case Mul:
			i.program[i.program[i.pointer+3]] = i.program[i.program[i.pointer+1]] * i.program[i.program[i.pointer+2]]
			i.pointer += 4
		case Halt:
			return nil
		default:
			return fmt.Errorf("Invalid opcode: %v found at pointer: %v", i.program[i.pointer], i.pointer)
		}
	}
}

// Peek allows inspecting the program
func (i *IntcodeComputer) Peek(pos int) int {
	return i.program[pos]
}

// Poke allows inspecting the program
func (i *IntcodeComputer) Poke(pos, val int) {
	i.program[pos] = val
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

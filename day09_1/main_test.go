package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/srowles/advent2019"
)

func TestCopy(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("1.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}

	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		}
		fmt.Println(computer.Output())
	}
	t.Fail()
}

func Test16Digit(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("2.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}

	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		}
		fmt.Println(computer.Output())
	}
	t.Fail()
}

func TestLarge(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("3.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}

	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		}
		fmt.Println(computer.Output())
	}
	t.Fail()
}

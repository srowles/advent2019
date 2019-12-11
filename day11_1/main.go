package main

import (
	"fmt"
	"log"

	"github.com/srowles/advent2019"
)

type point struct {
	x, y int
}

func main() {
	var robotPosition point
	var robotDir int
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}

	painted := make(map[point]int)
	for {
		computer.Input(painted[robotPosition])
		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("halted after paint")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		// output painting - not required atm
		colour := computer.Output()
		painted[robotPosition] = colour
		// get move instruction
		err = computer.Run()
		move := computer.Output()
		if move == 1 {
			robotDir++
		} else {
			robotDir--
		}

		if robotDir < 0 {
			robotDir = 3
		} else if robotDir > 3 {
			robotDir = 0
		}

		if robotDir == 0 {
			robotPosition = point{x: robotPosition.x, y: robotPosition.y - 1}
		}
		if robotDir == 1 {
			robotPosition = point{x: robotPosition.x + 1, y: robotPosition.y}
		}
		if robotDir == 2 {
			robotPosition = point{x: robotPosition.x, y: robotPosition.y + 1}
		}
		if robotDir == 3 {
			robotPosition = point{x: robotPosition.x - 1, y: robotPosition.y}
		}

		if err == advent2019.ErrHalted {
			fmt.Println("Haleted after move")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
	}

	fmt.Println(painted)
	fmt.Println("UniqueCount:", len(painted))
}

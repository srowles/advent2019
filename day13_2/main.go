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
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}
	computer.Poke(0, 2)
	score := 0
	var paddlex int
	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("halted first")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		x := computer.Output()

		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("halted after x")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		y := computer.Output()

		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("halted after y")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		tileID := computer.Output()

		if x == -1 && y == 0 {
			score = tileID
			fmt.Println(score)
		}

		if tileID == 3 {
			paddlex = x
		}

		if tileID == 4 {
			// we know ball location, compare to paddle to move it
			if x > paddlex {
				computer.Input(1)
			} else if x < paddlex {
				computer.Input(-1)
			} else {
				computer.Input(0)
			}
		}

	}

	fmt.Println(score)
}

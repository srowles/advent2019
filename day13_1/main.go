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

	count := 0
	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("halted after paint")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		x := computer.Output()

		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("Haleted after move")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		y := computer.Output()

		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("Haleted after move")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		tileID := computer.Output()

		fmt.Println(x, y, tileID)
		if tileID == 2 {
			count++
		}
	}

	fmt.Println(count)
}

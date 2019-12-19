package main

import (
	"fmt"
	"log"

	"github.com/srowles/advent2019"
)

type coord struct {
	x, y int
}

func main() {
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}
	prog := computer.GetProgram()
	count := 0
	area := make(map[coord]int)
	startx := 0
	for y := 0; y < 5000; y++ {
		foundHashThisRow := false
		for x := startx; x < 5000; x++ {
			cpy := make([]int, len(prog))
			for i, v := range prog {
				cpy[i] = v
			}
			computer, _ := advent2019.CreateIntcodeComputer(cpy)
			computer.Input(x)
			computer.Input(y)
			computer.Run()
			out := computer.Output()
			if out == 0 {
				fmt.Print(".")
				if foundHashThisRow {
					break
				}
			} else {
				fmt.Print("#")
				if !foundHashThisRow {
					startx = x
				}
				foundHashThisRow = true
				area[coord{x: x, y: y}] = 1
				count++
			}
			check(area, x, y)
		}
		fmt.Println()
	}
	fmt.Println(count)
}

// check back up to see if we fit..
func check(area map[coord]int, x, y int) {
	if area[coord{x: x - 99, y: y - 99}] == 0 { // is top left in beam?
		return
	}
	if area[coord{x: x, y: y - 99}] == 0 { // is top right in beam?
		return
	}
	if area[coord{x: x - 99, y: y}] == 0 { // is bottom left in beam?
		return
	}
	// box fits, what is the location?
	log.Fatalf("Box top rigth is at: %#v", coord{x: x - 99, y: y - 99})

}

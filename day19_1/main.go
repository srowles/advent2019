package main

import (
	"fmt"
	"log"

	"github.com/srowles/advent2019"
)

func main() {
	count := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
			if err != nil {
				log.Fatalf("Failed to create comuter with error: %v", err)
			}
			computer.Input(x)
			computer.Input(y)
			computer.Run()
			out := computer.Output()
			if out == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
				count++
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
}

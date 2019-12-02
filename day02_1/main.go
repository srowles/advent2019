package main

import (
	"fmt"
	"log"

	"github.com/srowles/advent2019"
)

func main() {
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}
	computer.Poke(1, 12)
	computer.Poke(2, 2)
	computer.Run()
	fmt.Println(computer.Peek(0))
}

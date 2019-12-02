package main

import (
	"fmt"
	"log"
	"os"

	"github.com/srowles/advent2019"
)

func main() {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
			if err != nil {
				log.Fatalf("Failed to create comuter with error: %v", err)
			}
			computer.Poke(1, noun)
			computer.Poke(2, verb)
			computer.Run()
			fmt.Printf("noun=%d, verb=%d, answer = %d, output=%d\n", noun, verb, 100*noun+verb, computer.Peek(0))
			if computer.Peek(0) == 19690720 {
				os.Exit(0)
			}
		}
	}
}

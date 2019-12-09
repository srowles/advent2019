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

	computer.Input(1)
	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		fmt.Println(computer.Output())
	}
}

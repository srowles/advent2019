package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/srowles/advent2019"
)

func main() {
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}

	computer.Poke(0, 2)
	computer.Inputs(commandsToInts("A,B"))                 //main
	computer.Inputs(commandsToInts("R,6"))                 //A
	computer.Inputs(commandsToInts("L,6,L,6,4,L,6,2,L,6")) //B
	computer.Inputs(commandsToInts("2"))                   //C
	computer.Inputs(commandsToInts("y"))                   // video output
	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		fmt.Printf("%s", string(byte(computer.Output())))
	}
}

func commandsToInts(in string) []int {
	a := strings.Split(in, ",")
	var out []int
	for _, i := range a {
		out = append(out, int(i[0]))
		out = append(out, 44)
	}
	out[len(out)-1] = 10
	return out
}

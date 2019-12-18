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
	computer.Inputs(commandsToInts("A,B,A,B,C,A,B,C,A,C")) //main
	computer.Inputs(commandsToInts("R,6,L,6,L,10"))        //A
	computer.Inputs(commandsToInts("L,8,L,6,L,10,L,6"))    //B
	computer.Inputs(commandsToInts("R,6,L,8,L,10,R,6"))    //C
	computer.Inputs(commandsToInts("n"))                   // video output
	var lastoutput int
	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			fmt.Println("halted")
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		lastoutput = computer.Output()
		fmt.Printf("%s", string(byte(lastoutput)))
	}
	fmt.Println("\n", lastoutput)
}

func commandsToInts(in string) []int {
	a := strings.Split(in, ",")
	var out []int
	for _, i := range a {
		for _, c := range i {
			out = append(out, int(c))
		}
		out = append(out, 44)
	}
	out[len(out)-1] = 10
	return out
}

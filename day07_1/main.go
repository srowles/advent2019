package main

import (
	"fmt"
	"log"

	"github.com/srowles/advent2019"
)

type comp interface {
	Output() int
	Input(val int)
	Run() error
}

func main() {
	max := 0
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						if reused(a, b, c, d, e) {
							continue
						}
						o := run(a, b, c, d, e)
						if o > max {
							max = o
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}

func reused(a, b, c, d, e int) bool {
	if a == b || a == c || a == d || a == e {
		return true
	}
	if b == c || b == d || b == e {
		return true
	}
	if c == d || c == e {
		return true
	}
	if d == e {
		return true
	}
	return false
}

func run(a, b, c, d, e int) int {
	var amps []comp
	for i := 0; i < 5; i++ {
		computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
		if err != nil {
			log.Fatalf("Failed to create comuter with error: %v", err)
		}
		amps = append(amps, computer)
	}
	amps[0].Input(a)
	amps[1].Input(b)
	amps[2].Input(c)
	amps[3].Input(d)
	amps[4].Input(e)

	amps[0].Input(0)
	for i, amp := range amps {
		amp.Run()
		if i < 4 {
			amps[i+1].Input(amp.Output())
		}
	}
	return amps[4].Output()
}

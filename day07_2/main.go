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
	for a := 5; a <= 9; a++ {
		for b := 5; b <= 9; b++ {
			for c := 5; c <= 9; c++ {
				for d := 5; d <= 9; d++ {
					for e := 5; e <= 9; e++ {
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
	haltCount := 0
outer:
	for {
		for i, amp := range amps {
			err := amp.Run()
			if err == advent2019.ErrHalted {
				haltCount++
			}
			if i < 4 {
				amps[i+1].Input(amp.Output())
			} else {
				amps[0].Input(amp.Output())
			}
			if haltCount == 5 {
				break outer
			}
		}
	}
	return amps[4].Output()
}

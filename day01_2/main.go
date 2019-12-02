package main

import (
	"fmt"
	"log"

	"github.com/srowles/advent2019"
)

func main() {
	moduleMasses, err := advent2019.GetNumbers("input1.txt")
	if err != nil {
		log.Fatalf("Failed to read file with error: %v", err)
	}

	totalFuel := 0
	for _, mass := range moduleMasses {
		totalFuel += fuel(mass)
	}
	fmt.Println(totalFuel)
}

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module,
// take its mass, divide by three, round down, and subtract 2.
func fuel(mass int) int {
	fuelRequired := (mass / 3) - 2
	if fuelRequired > 0 {
		fuelRequired += fuel(fuelRequired)
	} else {
		return 0
	}
	return fuelRequired
}

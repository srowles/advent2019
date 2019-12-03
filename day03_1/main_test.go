package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDraw(t *testing.T) {
	drawWire("R8,U5,L5,D3", 1)
	fmt.Println(grid)
	t.Fail()
}

func TestDrawTwoSmallWires(t *testing.T) {
	drawWire("R8,U5,L5,D3", 1)
	drawWire("U7,R6,D4,L4", 2)
	overlaps := findOverlaps()
	shortest := math.MaxInt32
	for _, overlap := range overlaps {
		d := dist(overlap)
		if d < shortest && d > 0 {
			shortest = d
			fmt.Println("shortest:", shortest)
		}
	}
	fmt.Println("Shortest:", shortest)
	fmt.Println(grid)
	fmt.Println("overlap:", overlaps)
	t.Fail()
}

func TestDrawTwoMediumWires(t *testing.T) {
	drawWire("R75,D30,R83,U83,L12,D49,R71,U7,L72", 1)
	drawWire("U62,R66,U55,R34,D71,R55,D58,R83", 2)
	overlaps := findOverlaps()
	shortest := math.MaxInt32
	for _, overlap := range overlaps {
		d := dist(overlap)
		if d < shortest && d > 0 {
			shortest = d
			fmt.Println("shortest:", shortest)
		}
	}
	fmt.Println("Shortest:", shortest)
	fmt.Println(grid)
	fmt.Println("overlap:", overlaps)
	t.Fail()
}

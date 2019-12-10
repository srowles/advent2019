package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	ex1 := `.#..#
.....
#####
....#
...##`
	astroidField, err := readInput(ex1)
	if err != nil {
		log.Fatalf("Failed to read input with error: %v", err)
	}
	viewMap := make(map[point]int)
	for start := range astroidField {
		count := 0
		for other := range astroidField {
			if start == other {
				continue
			}
			// fmt.Print("Checking start=", start, "other=", other)
			uvec, length := unitvec(start, other)
			// fmt.Printf("start -> other = %#v -> %v\n", uvec, length)
			blocked := false
			for mayBlock := range astroidField {
				if mayBlock == start || mayBlock == other {
					continue
				}
				uv, ol := unitvec(start, mayBlock)
				// fmt.Printf("blocking? %v %#v -> %v\n", mayBlock, uv, ol)
				if uv == uvec && ol < length {
					// fmt.Println("blocked by", mayBlock, uv, ol)
					// closer with same vec, must be blocking
					blocked = true
					break
				}
			}
			if !blocked {
				count++
			}
		}
		viewMap[start] = count
		fmt.Printf("%#v - %d\n", start, count)
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			fmt.Print(viewMap[point{x: x, y: y}])
		}
		fmt.Println()
	}
}

func unitvec(start, other point) (vec, float64) {
	xv := float64(other.x - start.x)
	yv := float64(other.y - start.y)
	length := math.Sqrt(float64((other.x-start.x)*(other.x-start.x)) + float64((other.y-start.y)*(other.y-start.y)))

	return vec{x: xv / length, y: yv / length}, length
}

type point struct {
	x, y int
}

type vec struct {
	x, y float64
}

func readInput(data string) (map[point]bool, error) {
	lines := strings.Split(data, "\n")
	result := make(map[point]bool)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				result[point{x: x, y: y}] = true
			}
		}
	}

	return result, nil
}

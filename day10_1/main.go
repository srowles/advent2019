package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	// 	ex1 := `.#..#
	// .....
	// #####
	// ....#
	// ...##`

	in := `.###.###.###.#####.#
#####.##.###..###..#
.#...####.###.######
######.###.####.####
#####..###..########
#.##.###########.#.#
##.###.######..#.#.#
.#.##.###.#.####.###
##..#.#.##.#########
###.#######.###..##.
###.###.##.##..####.
.##.####.##########.
#######.##.###.#####
#####.##..####.#####
##.#.#####.##.#.#..#
###########.#######.
#.##..#####.#####..#
#####..#####.###.###
####.#.############.
####.#.#.##########.`
	astroidField, err := readInput(in)
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
				if eq(uv, uvec) && ol < length {
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
	}

	max := 0
	var maxp point
	for k, v := range viewMap {
		if v > max {
			max = v
			maxp = k
		}
	}

	fmt.Println(maxp, max)

	// for y := 0; y < 5; y++ {
	// 	for x := 0; x < 5; x++ {
	// 		fmt.Print(viewMap[point{x: x, y: y}])
	// 	}
	// 	fmt.Println()
	// }
}

func eq(a, b vec) bool {
	if a == b {
		return true
	}
	if close(a) == close(b) {
		return true
	}

	return false
}

func close(v vec) vec {
	return vec{x: toFixed(v.x, 5), y: toFixed(v.y, 5)}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func unitvec(start, other point) (vec, float64) {
	xv := float64(other.x - start.x)
	yv := float64(other.y - start.y)
	length := math.Sqrt((xv * xv) + (yv * yv))

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

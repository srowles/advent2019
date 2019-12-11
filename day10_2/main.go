package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

var pi180 = math.Pi / 180
var piby2 = math.Pi / 2

func main() {
	in := `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`

	// big ex
	// 	in := `.#..##.###...#######
	// ##.############..##.
	// .#.######.########.#
	// .###.#######.####.#.
	// #####.##.#.##.###.##
	// ..#####..#.#########
	// ####################
	// #.####....###.#.#.##
	// ##.#################
	// #####.##.###..####..
	// ..######..##.#######
	// ####.##.####...##..#
	// .#####..#.######.###
	// ##...#.##########...
	// #.##########.#######
	// .####.#.###.###.#.##
	// ....##.##.###..#####
	// .#.#.###########.###
	// #.#.#.#####.####.###
	// ###.##.####.##.#..##`

	// in := `.###.###.###.#####.#
	// #####.##.###..###..#
	// .#...####.###.######
	// ######.###.####.####
	// #####..###..########
	// #.##.###########.#.#
	// ##.###.######..#.#.#
	// .#.##.###.#.####.###
	// ##..#.#.##.#########
	// ###.#######.###..##.
	// ###.###.##.##..####.
	// .##.####.##########.
	// #######.##.###.#####
	// #####.##..####.#####
	// ##.#.#####.##.#.#..#
	// ###########.#######.
	// #.##..#####.#####..#
	// #####..#####.###.###
	// ####.#.############.
	// ####.#.#.##########.`
	astroidField, err := readInput(in)
	if err != nil {
		log.Fatalf("Failed to read input with error: %v", err)
	}
	var destroyed []point
	// start := point{x: 8, y: 16}
	// start := point{x: 11, y: 13} // big example
	start := point{x: 8, y: 3} // small example
	for len(astroidField) > 1 {
		for i := -piby2; i <= (3 * piby2); i = i + 0.01 {
			up := int(math.Sin(i) * 1000)
			right := int(math.Cos(i) * 1000)
			target := point{x: start.x + right, y: start.y + up}
			laserVec, _ := unitvec(start, target)
			impactPoints := make(map[point]float64)
			for other := range astroidField {
				if start == other {
					continue
				}
				uvec, length := unitvec(start, other)
				if eq(laserVec, uvec) {
					impactPoints[other] = length
				}
			}
			if len(impactPoints) > 0 {
				min := math.MaxFloat64
				var closest *point
				for p, l := range impactPoints {
					if l < min {
						cpy := p
						closest = &cpy
						min = l
					}
				}
				if closest != nil {
					destroyed = append(destroyed, *closest)
					fmt.Println(*closest)
					printField(astroidField)
					delete(astroidField, *closest)
				}
			}
		}
	}
	fmt.Println(destroyed)
}

func printField(field map[point]bool) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 15; x++ {
			p := point{x: x, y: y}
			if field[p] == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
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
	return vec{x: toFixed(v.x, 2), y: toFixed(v.y, 2)}
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

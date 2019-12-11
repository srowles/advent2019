package main

import (
	"fmt"
	"log"
	"math"
	"sort"
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

	type pointWithAngle struct {
		p point
		a float64
	}
	var astroids []pointWithAngle
	start := point{x: 8, y: 16}
	for other := range astroidField {
		if start == other {
			continue
		}

		uvec, length := unitvec(start, other)
		blocked := false
		for mayBlock := range astroidField {
			if mayBlock == start || mayBlock == other {
				continue
			}
			uv, ol := unitvec(start, mayBlock)
			if eq(uv, uvec) && ol < length {
				blocked = true
				break
			}
		}
		if !blocked {
			astroids = append(astroids, pointWithAngle{other, angle(other)})
		}
	}
	sort.Slice(astroids, func(i, j int) bool {
		return astroids[i].a < astroids[j].a
	})

	for i, p := range astroids {
		fmt.Println(i+1, p)
	}
}

func angle(p point) float64 {
	v, _ := unitvec(point{x: 8, y: 16}, p)
	angle := math.Atan2(v.y, v.x) - math.Atan2(-1, 0)
	if angle < 0 {
		angle += 2 * math.Pi
	}
	return angle
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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var grid = make(map[point]int)
var zerox = 0
var zeroy = 0

type point struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file with error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i, line := range lines {
		drawWire(line, i+1)
	}
	overlaps := findOverlaps()
	shortest := math.MaxInt32
	for _, overlap := range overlaps {
		d := dist(overlap)
		if d < shortest && d > 0 {
			shortest = d
		}
	}
	fmt.Println("Shortest:", shortest)
}

func drawWire(input string, val int) {
	px := zerox
	py := zeroy
	instructions := strings.Split(input, ",")
	for _, i := range instructions {
		dist, _ := strconv.Atoi(i[1:])
		switch i[0] {
		case 'U':
			for i := 0; i < dist; i++ {
				py++
				mark(px, py, val)
			}
		case 'D':
			for i := 0; i < dist; i++ {
				py--
				mark(px, py, val)
			}
		case 'L':
			for i := 0; i < dist; i++ {
				px--
				mark(px, py, val)
			}
		case 'R':
			for i := 0; i < dist; i++ {
				px++
				mark(px, py, val)
			}
		}
	}
}

func mark(x, y, val int) {
	p := point{x: x, y: y}
	if grid[p] == 0 {
		grid[p] = val
	} else if grid[p] != val {
		grid[p] = 9
	}
}

func dist(p point) int {
	a := zerox - p.x
	if a < 0 {
		a = -a
	}
	b := zeroy - p.y
	if b < 0 {
		b = -b
	}
	return a + b
}

func findOverlaps() []point {
	var overs []point
	for p, v := range grid {
		if v == 9 {
			overs = append(overs, p)
		}
	}

	return overs
}

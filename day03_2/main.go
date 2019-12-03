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
var lens = make(map[point]int)
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
		d := lens[overlap]
		if d < shortest && d > 0 {
			shortest = d
		}
	}
	fmt.Println("Shortest:", shortest)
}

func drawWire(input string, val int) {
	px := zerox
	py := zeroy
	l := 0
	instructions := strings.Split(input, ",")
	for _, i := range instructions {
		dist, _ := strconv.Atoi(i[1:])
		switch i[0] {
		case 'U':
			for i := 0; i < dist; i++ {
				py++
				l++
				mark(px, py, val, l)
			}
		case 'D':
			for i := 0; i < dist; i++ {
				py--
				l++
				mark(px, py, val, l)
			}
		case 'L':
			for i := 0; i < dist; i++ {
				px--
				l++
				mark(px, py, val, l)
			}
		case 'R':
			for i := 0; i < dist; i++ {
				px++
				l++
				mark(px, py, val, l)
			}
		}
	}
}

func mark(x, y, val, l int) {
	p := point{x: x, y: y}
	if grid[p] == 0 {
		grid[p] = val
		lens[p] = l
	} else if grid[p] != val {
		lens[p] += l
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

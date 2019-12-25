package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var grid = make(map[point]int)
var grids []map[point]int

type point struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	parseIntoGrid(string(data))
	printGrid(grid)
	for {
		newGrid := processGrid()
		printGrid(newGrid)
		if find(newGrid) {
			calcDiversity(newGrid)
			return
		}
		grids = append(grids, newGrid)
		grid = newGrid
	}
}

func calcDiversity(newGrid map[point]int) {
	bio := 0
	multiplier := 1
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			p := point{x: x, y: y}
			if newGrid[p] == 1 {
				bio += multiplier
			}
			multiplier *= 2
		}
	}
	fmt.Println("biodiversity:", bio)
}

func find(newGrid map[point]int) bool {
	for _, g := range grids {
		if eq(g, newGrid) {
			return true
		}
	}

	return false
}

func eq(a map[point]int, b map[point]int) bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			p := point{x: x, y: y}
			if b[p] != a[p] {
				return false
			}
		}
	}
	fmt.Println("MATCH!")
	return true
}

func processGrid() map[point]int {
	newGrid := make(map[point]int)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			p := point{x: x, y: y}
			c := adjacentBugs(p)
			newGrid[p] = grid[p]
			if grid[p] == 1 && c != 1 {
				newGrid[p] = 0
			}
			if grid[p] == 0 && (c == 1 || c == 2) {
				newGrid[p] = 1
			}
		}
	}
	return newGrid
}

func adjacentBugs(p point) int {
	up := point{x: p.x, y: p.y - 1}
	down := point{x: p.x, y: p.y + 1}
	left := point{x: p.x - 1, y: p.y}
	right := point{x: p.x + 1, y: p.y}
	count := grid[up]
	count += grid[down]
	count += grid[left]
	count += grid[right]
	return count
}

func parseIntoGrid(data string) {
	for y, line := range strings.Split(strings.TrimSpace(data), "\n") {
		for x, c := range line {
			p := point{x: x, y: y}
			if c == '#' {
				grid[p] = 1
			}
		}
	}
}

func printGrid(g map[point]int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			o := "."
			if g[point{x: x, y: y}] == 1 {
				o = "#"
			}
			fmt.Print(o)
		}
		fmt.Println()
	}
	fmt.Println()
}

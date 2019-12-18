package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var oxygenStart coord
var maxx, maxy int

func main() {
	data, err := ioutil.ReadFile("map.txt")
	if err != nil {
		log.Fatalf("read err: %v", err)
	}
	area := parseMap(string(data))
	fmt.Println(countFreeSpace(area))
	time := 1
	for countFreeSpace(area) > 0 {
		var toFlowFrom []coord
		for y := 0; y < maxy; y++ {
			for x := 0; x < maxx; x++ {
				loc := coord{x: x, y: y}
				if area[loc] == 2 {
					toFlowFrom = append(toFlowFrom, loc)
				}
			}
		}
		for _, loc := range toFlowFrom {
			l := coord{x: loc.x + 1, y: loc.y}
			if area[l] == 0 {
				area[l] = 2
			}
			l = coord{x: loc.x, y: loc.y + 1}
			if area[l] == 0 {
				area[l] = 2
			}
			l = coord{x: loc.x - 1, y: loc.y}
			if area[l] == 0 {
				area[l] = 2
			}
			l = coord{x: loc.x, y: loc.y - 1}
			if area[l] == 0 {
				area[l] = 2
			}
		}
		render(area)
		time++
	}
	fmt.Println("Time taken:", time)
}

func render(area map[coord]int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			fmt.Print(charAt(area, x, y))
		}
		fmt.Println()
	}
}

func charAt(area map[coord]int, x, y int) string {
	if area[coord{x: x, y: y}] == 0 {
		return " "
	}
	if area[coord{x: x, y: y}] == 1 {
		return "#"
	}
	if area[coord{x: x, y: y}] == 2 {
		return "O"
	}
	return "."
}

type coord struct {
	x, y int
}

func parseMap(input string) map[coord]int {
	area := make(map[coord]int)
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if y > maxy {
			maxy = y
		}
		for x, c := range line {
			if x > maxx {
				maxx = x
			}
			loc := coord{x: x, y: y}
			if c == '#' {
				area[loc] = 1
			}
			if c == 'O' {
				oxygenStart = loc
				area[loc] = 2
			}
			if c == ' ' {
				area[loc] = 0
			}
		}
	}

	return area
}

func countFreeSpace(area map[coord]int) int {
	count := 0
	for _, i := range area {
		if i == 0 {
			count++
		}
	}
	return count
}

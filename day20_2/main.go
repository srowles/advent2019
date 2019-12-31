package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var doors map[coord]*door
var rawMaze map[coord]rune
var maze map[coord]*tile
var maxx, maxy int
var entrance *tile
var exit *tile

type door struct {
	pos    coord
	name   string
	inpos  coord
	outpos coord
}

type coord struct {
	x, y int
}

type tile struct {
	loc        coord
	symbol     rune
	n, s, e, w *tile
}

func (t tile) String() string {
	return fmt.Sprintf("[%d,%d] - %c - n:%p s:%p e:%p w:%p", t.loc.x, t.loc.y, t.symbol, t.n, t.s, t.e, t.w)
}

func main() {
	rawMaze = make(map[coord]rune)
	maze = make(map[coord]*tile)
	doors = make(map[coord]*door)
	parseInput(fromFile("1.txt"))
	postProcessMaze()
	buildGraph()
	printMaze()
	search(entrance, nil, make([]*tile, 0))
}

func search(t *tile, p *tile, tiles []*tile) bool {
	if t == exit {
		fmt.Println("Found Exit!", len(tiles))
		return true
	}
	tiles = append(tiles, t)
	if okToWalk(t.n, tiles) {
		search(t.n, t, tiles)
	}
	if okToWalk(t.s, tiles) {
		search(t.s, t, tiles)
	}
	if okToWalk(t.e, tiles) {
		search(t.e, t, tiles)
	}
	if okToWalk(t.w, tiles) {
		search(t.w, t, tiles)
	}
	return false
}

func okToWalk(t *tile, tiles []*tile) bool {
	if t == nil {
		return false
	}

	for _, tt := range tiles {
		if t.loc == tt.loc {
			return false
		}
	}

	return true
}

func buildGraph() {
	var in, out coord
	for p, d := range doors {
		if d.name == "AA" {
			in = d.inpos
			continue
		}
		if d.name == "ZZ" {
			out = d.inpos
			continue
		}
		for op, od := range doors {
			if op == p {
				// skip ourselves
				continue
			}
			if d.name == od.name {
				// found other end of door, attach
				d.outpos = od.inpos
			}
		}
	}

	for c, t := range maze {
		if c == in {
			entrance = t
			fmt.Println("found entrace:", entrance)
		}
		if c == out {
			exit = t
			fmt.Println("found exit:", exit)
		}

		up := coord{x: c.x, y: c.y - 1}
		down := coord{x: c.x, y: c.y + 1}
		left := coord{x: c.x - 1, y: c.y}
		right := coord{x: c.x + 1, y: c.y}

		if maze[up] != nil {
			t.n = maze[up]
		}
		if maze[down] != nil {
			t.s = maze[down]
		}
		if maze[left] != nil {
			t.w = maze[left]
		}
		if maze[right] != nil {
			t.e = maze[right]
		}

		if d, found := doors[c]; found {
			if d.name == "AA" || d.name == "ZZ" {
				continue
			}
			if d.pos == up {
				fmt.Println("connected north", t.loc, "to", d.outpos)
				t.n = maze[d.outpos]
			}
			if d.pos == down {
				fmt.Println("connected south", t.loc, "to", d.outpos)
				t.s = maze[d.outpos]
			}
			if d.pos == left {
				fmt.Println("connected west", t.loc, "to", d.outpos)
				t.w = maze[d.outpos]
			}
			if d.pos == right {
				fmt.Println("connected east", t.loc, "to", d.outpos)
				t.e = maze[d.outpos]
			}
		}
	}
}

func postProcessMaze() {
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			p := coord{x: x, y: y}
			c := rawMaze[p]
			if c == '#' || c == ' ' || c == 0 {
				continue
			}
			if c == '.' {
				addTile(p)
				continue
			}
		}
	}
}

func storeDoor(p coord, attach coord) {
	name := fmt.Sprintf("%c", rawMaze[p])

	up := coord{x: p.x, y: p.y - 1}
	down := coord{x: p.x, y: p.y + 1}
	left := coord{x: p.x - 1, y: p.y}
	right := coord{x: p.x + 1, y: p.y}
	if rawMaze[up] != '.' && rawMaze[up] != '#' && rawMaze[up] != '\x00' && rawMaze[up] != ' ' {
		name = fmt.Sprintf("%c%s", rawMaze[up], name)
		doors[attach] = &door{pos: p, name: name, inpos: attach}
		return
	}
	if rawMaze[down] != '.' && rawMaze[down] != '#' && rawMaze[down] != '\x00' && rawMaze[down] != ' ' {
		name = fmt.Sprintf("%s%c", name, rawMaze[down])
		doors[attach] = &door{pos: p, name: name, inpos: attach}
		return
	}
	if rawMaze[left] != '.' && rawMaze[left] != '#' && rawMaze[left] != '\x00' && rawMaze[left] != ' ' {
		name = fmt.Sprintf("%c%s", rawMaze[left], name)
		doors[attach] = &door{pos: p, name: name, inpos: attach}
		return
	}
	if rawMaze[right] != '.' && rawMaze[right] != '#' && rawMaze[right] != '\x00' && rawMaze[right] != ' ' {
		name = fmt.Sprintf("%s%c", name, rawMaze[right])
		doors[attach] = &door{pos: p, name: name, inpos: attach}
		return
	}
	if len(name) != 2 {
		return
	}

}

func addTile(p coord) {
	maze[p] = &tile{symbol: '.', loc: p}

	up := coord{x: p.x, y: p.y - 1}
	down := coord{x: p.x, y: p.y + 1}
	left := coord{x: p.x - 1, y: p.y}
	right := coord{x: p.x + 1, y: p.y}
	if rawMaze[up] != '.' && rawMaze[up] != '#' && rawMaze[up] != '\x00' && rawMaze[up] != ' ' {
		storeDoor(up, p)
	}
	if rawMaze[down] != '.' && rawMaze[down] != '#' && rawMaze[down] != '\x00' && rawMaze[down] != ' ' {
		storeDoor(down, p)
	}
	if rawMaze[left] != '.' && rawMaze[left] != '#' && rawMaze[left] != '\x00' && rawMaze[left] != ' ' {
		storeDoor(left, p)
	}
	if rawMaze[right] != '.' && rawMaze[right] != '#' && rawMaze[right] != '\x00' && rawMaze[right] != ' ' {
		storeDoor(right, p)
	}
}

func parseInput(data string) {
	for y, row := range strings.Split(data, "\n") {
		for x, c := range row {
			p := coord{x: x, y: y}
			rawMaze[p] = c
			if x > maxx {
				maxx = x
			}
		}
		if y > maxy {
			maxy = y
		}
	}
}

func printMaze() {
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			p := coord{x: x, y: y}
			if _, ok := maze[p]; ok {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func fromFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("failed to read file" + err.Error())
	}
	return string(data)
}

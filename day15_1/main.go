package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/srowles/advent2019"
)

var (
	area = make(map[coord]int)
	pos  = coord{x: 24, y: 24}
	dir  = 1
)

type coord struct {
	x, y int
}

func main() {
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}

	iterations := 0
	for {
		computer.Input(dir)
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error: %v", err)
			}
		}
		out := computer.Output()
		if out == 0 {
			markWall()
		}
		if out == 1 {
			move()
		}
		if out == 2 {
			// move()
			markOxygen()
			break
		}

		if iterations%1000 == 0 {
			render()
		}
		// fmt.Println(iterations)
		iterations++
		dir = rand.Intn(4) + 1
		// r := bufio.NewReader(os.Stdin)
		// c, err := r.ReadByte()
		// if err != nil {
		// 	panic("wtf" + err.Error())
		// }
		// switch c {
		// case 'w':
		// 	dir = 1
		// case 's':
		// 	dir = 2
		// case 'a':
		// 	dir = 3
		// case 'd':
		// 	dir = 4
		// }
	}
	render()
}

func markOxygen() {
	area[pos] = 2
}

func move() {
	switch dir {
	case 1:
		pos = coord{x: pos.x, y: pos.y - 1}
	case 2:
		pos = coord{x: pos.x, y: pos.y + 1}
	case 3:
		pos = coord{x: pos.x - 1, y: pos.y}
	case 4:
		pos = coord{x: pos.x + 1, y: pos.y}
	}
}

func markWall() {
	var loc coord
	switch dir {
	case 1:
		loc = coord{x: pos.x, y: pos.y - 1}
	case 2:
		loc = coord{x: pos.x, y: pos.y + 1}
	case 3:
		loc = coord{x: pos.x - 1, y: pos.y}
	case 4:
		loc = coord{x: pos.x + 1, y: pos.y}
	}
	area[loc] = 1
}

func render() {
	fmt.Printf("minx:%d,miny:%d,maxx:%d,maxy:%d\n", minx(), miny(), maxx(), maxy())
	for y := miny() - 5; y < maxy()+5; y++ {
		for x := minx() - 5; x < maxx()+5; x++ {
			fmt.Print(charAt(x, y))
		}
		fmt.Println()
	}
}

func miny() int {
	m := math.MaxInt64
	for l := range area {
		if l.y < m {
			m = l.y
		}
	}

	if pos.y < m {
		m = pos.y
	}
	return m
}

func minx() int {
	m := math.MaxInt64
	for l := range area {
		if l.x < m {
			m = l.x
		}
	}
	if pos.x < m {
		m = pos.x
	}
	return m
}

func maxy() int {
	m := 0
	for l := range area {
		if l.y > m {
			m = l.y
		}
	}
	if pos.y > m {
		m = pos.y
	}
	return m
}

func maxx() int {
	m := 0
	for l := range area {
		if l.x > m {
			m = l.x
		}
	}
	if pos.x > m {
		m = pos.x
	}
	return m
}

func charAt(x, y int) string {
	if area[coord{x: x, y: y}] == 2 {
		return "O"
	}

	if x == pos.x && y == pos.y {
		return "D"
	}
	if area[coord{x: x, y: y}] == 1 {
		return "#"
	}

	return "."
}

package main

import (
	"fmt"
)

type moon struct {
	x, y, z          int
	velx, vely, velz int
}

var (
	// moons = []*moon{
	// 	&moon{x: -1, y: 0, z: 2},
	// 	&moon{x: 2, y: -10, z: -7},
	// 	&moon{x: 4, y: -8, z: 8},
	// 	&moon{x: 3, y: 5, z: -1},
	// }
	// moons = []*moon{
	// 	&moon{x: -8, y: -10, z: 0},
	// 	&moon{x: 5, y: 5, z: 10},
	// 	&moon{x: 2, y: -7, z: 3},
	// 	&moon{x: 9, y: -8, z: -3},
	// }
	start = []*moon{
		&moon{x: -13, y: -13, z: -13},
		&moon{x: 5, y: -8, z: 3},
		&moon{x: -6, y: -10, z: -3},
		&moon{x: 0, y: 5, z: -5},
	}
	moons = []*moon{
		&moon{x: -13, y: -13, z: -13},
		&moon{x: 5, y: -8, z: 3},
		&moon{x: -6, y: -10, z: -3},
		&moon{x: 0, y: 5, z: -5},
	}
)

func main() {
	for i := 0; ; i++ {
		calcNewVelocity(moons[0], moons[1])
		calcNewVelocity(moons[0], moons[2])
		calcNewVelocity(moons[0], moons[3])
		calcNewVelocity(moons[1], moons[2])
		calcNewVelocity(moons[1], moons[3])
		calcNewVelocity(moons[2], moons[3])

		for _, m := range moons {
			applyVelocity(m)
		}

		matchCount := 0
		for n, m := range moons {
			if m.x == start[n].x && m.velx == start[n].velx {
				matchCount++
			}
		}
		if matchCount == 4 {
			fmt.Println("x matched after", i, "iterations")
			return
		}
	}

	// xperiod = 268295 + 1 = 268296
	// yperiod = 231613 + 1 = 231614
	// zperiod = 23325 + 1 = 23326

	// Calculated externally
	// LCM = 362375881472136
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func printMoons(moons []*moon) {
	for _, m := range moons {
		fmt.Printf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>\n", m.x, m.y, m.z, m.velx, m.vely, m.velz)
	}
	fmt.Println()
}

func calcNewVelocity(moon1 *moon, moon2 *moon) {
	if moon1.x > moon2.x {
		moon1.velx--
		moon2.velx++
	} else if moon1.x < moon2.x {
		moon1.velx++
		moon2.velx--
	}
	if moon1.y > moon2.y {
		moon1.vely--
		moon2.vely++
	} else if moon1.y < moon2.y {
		moon1.vely++
		moon2.vely--
	}
	if moon1.z > moon2.z {
		moon1.velz--
		moon2.velz++
	} else if moon1.z < moon2.z {
		moon1.velz++
		moon2.velz--
	}
}

func applyVelocity(moon *moon) {
	moon.x += moon.velx
	moon.y += moon.vely
	moon.z += moon.velz
}

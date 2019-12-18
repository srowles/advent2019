package main

import (
	"fmt"
	"strconv"
)

func main() {
	signal := parse("59790677903322930697358770979456996712973859451709720515074487141246507419590039598329735611909754526681279087091321241889537569965210074382210124927546962637736867742660227796566466871680580005288100192670887174084077574258206307557682549836795598410624042549261801689113559881008629752048213862796556156681802163843211546443228186862314896620419832148583664829023116082772951046466358463667825025457939806789469683866009241229487708732435909544650428069263180522263909211986231581228330456441451927777125388590197170653962842083186914721611560451459928418815254443773460832555717155899456905676980728095392900218760297612453568324542692109397431554")
	initia := signal
	for i := 0; i < 10000; i++ {
		signal = append(signal, initia...)
	}
	for phase := 0; phase < 100; phase++ {
		var next []int
		for pos := 0; pos < len(signal); pos++ {
			next = append(next, getOutput(signal, pos))
		}
		signal = next
	}

	fmt.Println(signal[5979067 : 5979067+8])
}

func getOutput(signal []int, pos int) int {
	pattern := getPattern(pos, len(signal))
	// fmt.Println(pattern)
	out := 0
	for i, s := range signal {
		m := pattern[i]
		val := s * m
		// fmt.Println(s, "*", m, "=", val)
		out = out + val
	}
	// fmt.Println(out)
	return abs(out % 10)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

var basePattern = []int{0, 1, 0, -1}

// 0,1,0,-1
func getPattern(pos int, length int) []int {
	pos++
	var out []int
	for _, i := range basePattern {
		for r := 0; r < pos; r++ {
			out = append(out, i)
		}
	}

	for len(out) <= length {
		out = append(out, out...)
	}
	return out[1:]
}

func parse(in string) []int {
	var out []int
	for _, c := range in {
		n, _ := strconv.Atoi(string(c))
		out = append(out, n)
	}

	return out
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	parts := strings.Split("178416-676461", "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	count := 0
	for i := start; i <= end; i++ {
		if followsRules(i) {
			count++
		}
	}
	fmt.Println("matches:", count)
}

func followsRules(i int) bool {
	var nums []int
	nums = intToSlice(i, nums)
	prev := -1
	doubleFound := false
	for _, n := range nums {
		if n == prev {
			doubleFound = true
		}
		if n < prev {
			return false
		}
		prev = n
	}

	return doubleFound
}

func intToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return intToSlice(n/10, sequence)
	}
	return sequence
}

package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAngles(t *testing.T) {
	for i := -90; i < 270; i++ {
		up := math.Sin(float64(i) * pi180)
		right := math.Cos(float64(i) * pi180)
		fmt.Println(i, "a=", up, "b=", right)
	}
	t.Fail()
}

func TestRadians(t *testing.T) {
	for i := -piby2; i <= (3 * piby2); i = i + 0.01 {
		up := math.Sin(i)
		right := math.Cos(i)
		fmt.Println(i, "up=", up, "right=", right)
	}
	t.Fail()
}

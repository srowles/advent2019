package main

import (
	"fmt"
	"testing"

	"github.com/srowles/advent2019"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("t1.txt")
	require.NoError(t, err)
	computer.Run()
	fmt.Println(computer.Peek(4))
	t.Fail()
}

func TestSimple2(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("t2.txt")
	require.NoError(t, err)
	computer.Run()
	fmt.Println(computer.Peek(4))
	t.Fail()
}

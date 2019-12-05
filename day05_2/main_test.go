package main

import (
	"testing"

	"github.com/srowles/advent2019"
	"github.com/stretchr/testify/require"
)

func TestSimple1(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("1.txt")
	require.NoError(t, err)
	computer.Run()
	t.Fail()
}

func TestLarger(t *testing.T) {
	computer, err := advent2019.CreateIntcodeComputerFromFile("2.txt")
	require.NoError(t, err)
	computer.Run()
	t.Fail()
}

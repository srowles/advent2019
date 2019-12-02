package advent2019

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompExample(t *testing.T) {
	require := require.New(t)
	c, err := createIntcodeComputer([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50})
	require.NoError(err)
	err = c.Run()
	require.NoError(err)
	require.Equal(3500, c.Peek(0))
}

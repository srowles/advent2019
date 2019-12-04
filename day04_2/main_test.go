package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	require.True(t, followsRules(112233))
	require.False(t, followsRules(123444))
	require.True(t, followsRules(111122))
}

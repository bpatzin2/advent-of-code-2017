// test each function in main.go
package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1pt1(t *testing.T) {
	require.Equal(t, 1141, day1pt1())
}

func TestDay1pt2(t *testing.T) {
	require.Equal(t, 950, day1pt2())
}

func TestDay7pt1(t *testing.T) {
	require.Equal(t, "vmpywg", day7pt1())
}

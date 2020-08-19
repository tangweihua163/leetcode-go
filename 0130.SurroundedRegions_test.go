package leetcode_go

import (
	"fmt"
	"testing"
)

func TestSolveSurroundedRegions(t *testing.T) {

	board := [][]byte{
		{X, X, X, X},
		{X, O, O, X},
		{X, X, O, X},
		{X, O, X, X},
	}

	solveSurroundedRegions(board)

	fmt.Println(board)
}

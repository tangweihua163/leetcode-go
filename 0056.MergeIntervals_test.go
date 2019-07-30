package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	//fmt.Println(mergeIntervals([][]int{{1,3},{2,6},{8,10},{15,18}}))
	//fmt.Println(mergeIntervals([][]int{{1,4},{5,6}}))

	fmt.Println(mergeIntervals([][]int{
		{1, 3},
		{0, 2},
		{2, 3},
		{4, 6},
		{4, 5},
		{5, 5},
		{0, 2},
		{3, 3},
	}))

}

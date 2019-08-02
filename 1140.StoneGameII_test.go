package leetcode_go

import (
	"fmt"
	"testing"
)

func TestStoneGameII(t *testing.T) {
	fmt.Println(stoneGameII([]int{7468, 6245, 9261, 3958, 1986, 1074,
		5677, 9386, 1408, 1384, 8811, 3885, 9678, 8470, 8893,
		7514, 4941, 2148, 5217, 5425, 5307, 747, 1253, 3518, 5238,
		5834, 9133, 8391, 6100, 3362, 7807, 2581, 6121, 7684, 8744,
		9584, 4068, 7204, 4285, 8635}))
	fmt.Println(stoneGameII([]int{2, 7, 9, 4, 4}))
}

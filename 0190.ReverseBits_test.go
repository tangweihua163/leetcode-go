package leetcode_go

import (
	"fmt"
	"math"
	"testing"
)

func TestReverseBit(t *testing.T) {
	fmt.Println(reverseBits(uint32(1)) - math.MaxInt32)
}

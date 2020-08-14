package leetcode_go

import "math"

func reverseBits(num uint32) uint32 {
	var result uint32 = math.MaxUint32
	for i := 0; i < 32; i++ {
		if ((1 << i) & num) == 0 {
			result = result &^ (1 << (31 - i))
		}
	}
	return result
}

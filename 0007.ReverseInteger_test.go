package leetcode_go

import (
	"fmt"
	"testing"
)

func TestReverseInteger(t *testing.T) {
	fmt.Println(
		reverse(1),
		reverse(-1),
		reverse(0),
		reverse(2147447412),
		reverse(2147447413),

		reverse(-2147447412),
		reverse(-2147447413),
	)
}

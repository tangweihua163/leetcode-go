package leetcode_go

func multiply(num1 string, num2 string) string {

	result, _ := helper(num1, 0, num2, 0)

	var i int
	for i = 0; i < len(result)-1; i++ {
		if result[i] != '0' {
			break
		}
	}

	return result[i:]
}

func helper(num1 string, base1 int, num2 string, base2 int) (r string, base int) {

	if len(num1) == 1 {
		t := mul(num2, int(num1[0]-'0'))
		bt := base1 + base2
		if t == "0" {
			bt = 0
		}
		return t, bt
	}
	if len(num2) == 1 {
		t := mul(num1, int(num2[0]-'0'))
		bt := base1 + base2
		if t == "0" {
			bt = 0
		}
		return t, bt
	}

	mid1 := len(num1) / 2
	mid2 := len(num2) / 2

	a, basea := num1[:mid1], base1+len(num1)-mid1
	b, baseb := num1[mid1:], base1
	c, basec := num2[:mid2], base2+len(num2)-mid2
	d, based := num2[mid2:], base2

	ac, bac := helper(a, basea, c, basec)
	ad, bad := helper(a, basea, d, based)
	bc, bbc := helper(b, baseb, c, basec)
	bd, bbd := helper(b, baseb, d, based)

	t, baset := "0", 0

	t, baset = add(t, baset, ac, bac)
	t, baset = add(t, baset, ad, bad)
	t, baset = add(t, baset, bc, bbc)
	t, baset = add(t, baset, bd, bbd)

	return t, baset
}

func add(num1 string, base1 int, num2 string, base2 int) (r string, b int) {

	rl := max(len(num1)+base1, len(num2)+base2)
	result := make([]byte, rl+1)

	op1 := make([]byte, rl)
	var i = 0
	for ; i < rl-(len(num1)+base1); i++ {
		op1[i] = '0'
	}
	copy(op1[i:i+len(num1)], num1)
	for i = rl - base1; i < len(op1); i++ {
		op1[i] = '0'
	}

	op2 := make([]byte, rl)
	var j = 0
	for ; j < rl-(len(num2)+base2); j++ {
		op2[j] = '0'
	}
	copy(op2[j:j+len(num2)], num2)
	for j = rl - base2; j < len(op2); j++ {
		op2[j] = '0'
	}

	var c = 0
	var ri = len(result) - 1
	for k := rl - 1; k >= 0; k-- {
		sum := int(op1[k]-'0') + int(op2[k]-'0') + c
		c = sum / 10
		result[ri] = byte(sum%10 + '0')
		ri--
	}

	if c != 0 {
		result[ri] = byte(c + '0')
		return string(result), 0
	} else {
		return string(result[1:]), 0
	}
}

func mul(num1 string, m int) (r string) {

	if m == 0 {
		return "0"
	}
	if m == 1 {
		return num1
	}

	var c = 0
	var result = make([]byte, len(num1)+1)
	var j = len(result) - 1

	for i := len(num1) - 1; i >= 0; i-- {
		cal := m*int(num1[i]-'0') + c
		c = cal / 10
		result[j] = byte(cal%10) + '0'
		j--
	}
	if c != 0 {
		result[j] = byte(c) + '0'
		r = string(result)
		return
	} else {
		r = string(result[1:])
		return
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

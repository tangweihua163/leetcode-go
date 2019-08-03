package leetcode_go

func calculate(s string) int {

	number := make([]int, 0, len(s))
	operator := make([]byte, 0, len(s))

	number = append(number, 0)
	operator = append(operator, '+')

	var n = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			n = n*10 + int(s[i]-'0')
		case '+', '-', '*', '/':
			ln := len(number)
			lo := len(operator)
			top := operator[lo-1]
			if top == '*' {
				number[ln-1] = number[ln-1] * n
				operator[lo-1] = s[i]
			} else if top == '/' {
				number[ln-1] = number[ln-1] / n
				operator[lo-1] = s[i]
			} else {
				number = append(number, n)
				operator = append(operator, s[i])
			}
			n = 0
		}
	}
	number = append(number, n)

	ln := len(number)
	lo := len(operator)
	top := operator[lo-1]
	if top == '*' {
		number[ln-2] = number[ln-2] * number[ln-1]
		number = number[:ln-1]
		operator = operator[:lo-1]
	} else if top == '/' {
		number[ln-2] = number[ln-2] / number[ln-1]
		number = number[:ln-1]
		operator = operator[:lo-1]
	}

	i := 0
	j := 0
	for i < len(number)-1 && j < len(operator) {
		if operator[j] == '+' {
			number[i+1] = number[i] + number[i+1]
		} else if operator[j] == '-' {
			number[i+1] = number[i] - number[i+1]
		}
		i++
		j++
	}

	return number[len(number)-1]
}

func _calculate(s string) int {
	stack, num, op, res := []int{0}, 0, '+', 0
	for i, c := range s {
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		}

		// i == len(s)-1 时，最后一个字符
		if c == '+' || c == '-' || c == '*' || c == '/' || i == len(s)-1 {
			if op == '+' {
				stack = append(stack, num)
			}
			if op == '-' {
				stack = append(stack, -num)
			}
			if op == '*' {
				stack[len(stack)-1] *= num
			}
			if op == '/' {
				stack[len(stack)-1] /= num
			}
			op = c
			num = 0
		}
	}
	for _, v := range stack {
		res += v
	}
	return res
}

package libs

func DigitSet(num int) map[int]struct{} {
	var s = make(map[int]struct{})
	exists := struct{}{}

	for num > 0 {
		d := num % 10
		num /= 10
		s[d] = exists

	}

	return s
}

func SameSet(s1, s2 map[int]struct{}) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, _ := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}

	return true
}

func IsNumPandigitToN(num int, n int) bool {
	return IsNumPandigitAtoB(num, 1, n)
}

func IsNumPandigitAtoB(num int, a int, b int) bool {
	// assume a < b
	digits := DigitSet(num)
	min, max := 10, -1
	for k, _ := range digits {
		if k > max {
			max = k
		}
		if k < min {
			min = k
		}
	}

	return (len(digits) == (b-a+1) && min == a && max == b)
}

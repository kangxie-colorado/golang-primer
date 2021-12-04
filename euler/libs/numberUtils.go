package libs

func productOfArray(vals []int) int {
	product := 1
	for i := range vals {
		product *= vals[i]
	}

	return product
}

func HasDuplicatedDigit(num int) bool {
	var s = make(map[int]struct{})
	var exists = struct{}{}

	for num > 0 {
		d := num % 10
		if _, ok := s[d]; ok {
			return true
		}
		s[d] = exists
		num /= 10
	}

	return false
}

func HasZeroDigit(num int) bool {
	for num > 0 {
		if num%10 == 0 {
			return true
		}
		num /= 10
	}

	return false
}

func divideAndMod(n int, d int) (int, int) {
	return n / d, n % d
}

func UnitFrac(d int, maxLen int) []int {
	// calculate the fractional part until it appears to be repeating or it can divided evenly(e.g. 1/2=0.5)
	// now use the elementary dividing calculation
	fracs := []int{}
	numerator := 1
	for numerator%d != 0 {
		if numerator < d {
			fracs = append(fracs, 0)
			numerator *= 10
		} else {
			num, change := divideAndMod(numerator, d)
			fracs = append(fracs, num)
			numerator = change * 10
		}

		// always started with a 0, so the first it > maxLen it is just maxLen fractional digits
		if len(fracs) > maxLen {
			break
		}
	}

	// divided evenly
	if numerator%d == 0 {
		fracs = append(fracs, numerator/d)
	}

	return fracs
}

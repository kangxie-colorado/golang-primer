package libs

import "math"

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := int(2); i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func PrimeFactors(given int) []int {
	var pf []int
	if IsPrime(given) {
		pf = append(pf, given)
	} else {
		for i := 2; i < given; i++ {
			if IsPrime(i) && given%i == 0 {
				pf = append(pf, i)
				pf = append(pf, PrimeFactors(given/i)...)

				if productOfArray(pf) == given {
					break
				}
			}
		}
	}

	return pf
}

func turnOff(s []int, basis int) []int {
	for i := basis * 2; i < len(s); i += basis {
		s[i] = -1
	}

	return s
}

// trick, index is same as the number value
// use it with same sematical meaning
func PrimesUnderN(N int) []int {
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}

	a[0] = -1
	a[1] = -1

	for p := 0; p < N; p++ {
		if a[p] == -1 {
			continue
		}

		if IsPrime(p) {
			turnOff(a, p)
		}
	}

	primes := []int{}
	for _, p := range a {
		if p != -1 {
			primes = append(primes, p)
		}
	}

	return primes
}

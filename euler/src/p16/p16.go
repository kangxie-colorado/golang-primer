package main

func int_as_array_multiply_by_2(arr []int) []int {
	carry := 0
	for idx := range arr {
		n := arr[idx]*2 + carry
		carry = n / 10
		arr[idx] = n % 10
	}

	if carry > 0 {
		arr = append(arr, carry)
	}

	return arr
}

func sum_pow2_digit(n int) int {
	pows_digits := []int{1}
	for i := 1; i <= n; i++ {
		pows_digits = int_as_array_multiply_by_2(pows_digits)
	}

	sum := 0
	for _, n := range pows_digits {
		sum += n
	}

	return sum
}

func main() {
	print(sum_pow2_digit(1000))
}

package libs

func numAndCarryByAdd(num1 int, num2 int, carry int) (int, int) {
	num3 := (num1 + num2 + carry) % 10
	carry = (num1 + num2 + carry) / 10
	/**
	can you spot why doing it in this order is wrong?
	carry = (num1 + num2 + carry) / 10
	num3 := (num1 + num2 + carry) % 10
	**/

	return num3, carry
}

func numAndCarryByMul(num1 int, num2 int, carry int) (int, int) {
	num3 := (num1*num2 + carry) % 10
	carry = (num1*num2 + carry) / 10
	/**
	can you spot why doing it in this order is wrong?
	carry = (num1 + num2 + carry) / 10
	num3 := (num1 + num2 + carry) % 10
	**/

	return num3, carry
}

func AddTwoIntSlices(slice1 []int, slice2 []int) []int {
	ret := []int{}

	if len(slice1) > len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	carry := 0
	idx := 0
	num3 := 0
	for ; idx < len(slice1); idx++ {
		num3, carry = numAndCarryByAdd(slice1[idx], slice2[idx], carry)
		ret = append(ret, num3)
	}

	for ; idx < len(slice2); idx++ {
		num3, carry = numAndCarryByAdd(0, slice2[idx], carry)
		ret = append(ret, num3)
	}

	if carry > 0 {
		ret = append(ret, carry)
	}

	return ret
}

func MutliplySliceByInt(slice []int, m int) []int {
	ret := []int{}

	carry := 0
	digit := 0
	for _, num := range slice {
		digit, carry = numAndCarryByMul(num, m, carry)
		ret = append(ret, digit)
	}

	if carry > 0 {
		ret = append(ret, carry)
	}

	return ret
}

func MultiplySliceBySlice(slice1 []int, slice2 []int) []int {
	res := []int{}

	if len(slice1) > len(slice2) {
		slice1, slice2 = slice2, slice1
	}

	for idx, num := range slice1 {
		prefix := []int{}
		for num_to_prepend := idx; num_to_prepend > 0; num_to_prepend-- {
			prefix = append(prefix, 0)
		}

		to_mul_slice := append(prefix, slice2...)
		res = AddTwoIntSlices(res, MutliplySliceByInt(to_mul_slice, num))
	}

	return res
}

func IntToSlice(n int) []int {
	return IntToSliceBase(n, 10)
}

func IntToSliceBase(num int, base int) []int {
	ret := []int{}
	for num > 0 {
		ret = append(ret, (num % base))
		num /= base
	}

	return ret
}

func SliceToIntBase(nums []int, base int) int {
	// high digits to the end
	ret := 0
	for idx := len(nums) - 1; idx >= 0; idx-- {
		ret = ret*base + nums[idx]
	}

	return ret
}

func XPowY(x int, y int) []int {

	// x*x*x...
	pow := IntToSlice(x)
	for mul_times := 1; mul_times < y; mul_times++ {
		pow = MultiplySliceBySlice(pow, IntToSlice(x))
	}

	return pow
}

func TestIntSlicesEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareTwoNumSlices(slice1 []int, slice2 []int) int {
	// assume len(slice1) == len(slice2)
	for idx := range slice1 {
		if slice1[idx] < slice2[idx] {
			return -1
		}
		if slice1[idx] > slice2[idx] {
			return 1
		}

	}
	return 0
}

func _isNumInSlice(num int, left int, right int, slice []int) bool {
	if left >= right {
		return false
	}

	mid := left + (right-left)/2

	if slice[mid] == num {
		return true
	} else if slice[mid] < num {
		left = mid + 1
	} else {
		right = mid

	}

	return _isNumInSlice(num, left, right, slice)
}

func IsNumInSlice(num int, slice []int) bool {
	return _isNumInSlice(num, 0, len(slice), slice)
}

func UniqNumbersInSlice(nums []int) (int, []int) {
	var s = make(map[int]struct{})
	var exists = struct{}{}

	for _, n := range nums {
		s[n] = exists
	}

	uniqNums := []int{}
	for k, _ := range s {
		uniqNums = append(uniqNums, k)
	}

	return len(uniqNums), uniqNums
}

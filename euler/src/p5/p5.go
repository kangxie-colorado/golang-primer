// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

func getGCD(num1 int, num2 int) int {
	for num1%num2 != 0 {
		num1, num2 = num2, (num1 % num2)
		if num1 == 1 || num2 == 1 {
			return 1
		}
	}

	return num2
}

func getLCM(num1 int, num2 int) int {
	return num1 * num2 / getGCD(num1, num2)
}

func getLCMOfArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	lcm := 1
	for _, num := range nums {
		lcm = getLCM(lcm, num)
	}

	return lcm
}

func main() {
	println(getGCD(2, 3))
}

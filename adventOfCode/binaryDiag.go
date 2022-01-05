package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func powerConsumption(data []string) int {
	bitLen := len(data[0])

	nums := []int{}
	for _, d := range data {
		num, _ := strconv.ParseInt(d, 2, 0)
		nums = append(nums, int(num))
	}

	gamma, epsilon := getPowerRates(bitLen, nums)
	return gamma * epsilon
}

func getPowerRates(bitLen int, nums []int) (int, int) {
	var gammaRate, epsilonrate int = 0, 0

	for shift := 0; shift < bitLen; shift++ {
		oddNum := 0
		evenNum := 0
		for i, n := range nums {
			if n%2 == 0 {
				evenNum += 1
			} else {
				oddNum += 1
			}

			nums[i] = n >> 1
		}

		if oddNum > evenNum {
			gammaRate += int(math.Pow(2, float64(shift)))
		} else {
			epsilonrate += int(math.Pow(2, float64(shift)))
		}

	}

	return gammaRate, epsilonrate
}

func powerConsumptionStdin() int {
	scanner := bufio.NewScanner(os.Stdin)

	nums := []int{}
	bitLen := 0
	for scanner.Scan() {
		if bitLen == 0 {
			bitLen = len(scanner.Text())
		}
		num, _ := strconv.ParseInt(scanner.Text(), 2, 0)
		nums = append(nums, int(num))

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	gamma, epsilon := getPowerRates(bitLen, nums)
	return gamma * epsilon
}

func getO2Rate(bitLen int, nums []int) int {
	res := 0
	candidates := nums

	for shift := 0; shift < bitLen; shift++ {
		bigs := []int{}
		smalls := []int{}

		threshold := int(math.Pow(2, float64(bitLen-shift-1)))

		for _, n := range candidates {
			if n >= threshold {
				// bigs means prefix 1xxxx, no need to keep the leading 1
				// so n%threshold will be the new numbers to run thru
				bigs = append(bigs, n%threshold)
			} else {
				smalls = append(smalls, n%threshold)
			}
		}

		if len(bigs) >= len(smalls) {
			candidates = bigs
			res = res*2 + 1
		} else {
			candidates = smalls
			res = res * 2
		}

		if len(candidates) == 1 {
			// in case the one candidate was determined before all the bits are checked
			// bitLen-1-shift should be the left-shift power for pre-determined bits
			// e.g. [00100, 11110, 10110] by 2nd iteration, it will determine 110 the remainder of 11110 to be sole remaining candidate
			// and prefix is 11, so 11<<3+110 = 30
			res = res<<(bitLen-1-shift) + candidates[0]
			break
		}
	}

	return res
}

func getCO2Rate(bitLen int, nums []int) int {
	res := 0
	candidates := nums

	for shift := 0; shift < bitLen; shift++ {
		bigs := []int{}
		smalls := []int{}

		threshold := int(math.Pow(2, float64(bitLen-shift-1)))

		for _, n := range candidates {
			if n >= threshold {
				// bigs means prefix 1xxxx, no need to keep the leading 1
				// so n%threshold will be the new numbers to run thru
				bigs = append(bigs, n%threshold)
			} else {
				smalls = append(smalls, n%threshold)
			}
		}

		if len(bigs) < len(smalls) {
			candidates = bigs
			res = res*2 + 1
		} else {
			candidates = smalls
			res = res * 2
		}

		if len(candidates) == 1 {
			// in case the one candidate was determined before all the bits are checked
			// bitLen-1-shift should be the left-shift power for pre-determined bits
			// e.g. [00100, 11110, 10110] by 2nd iteration, it will determine 110 the remainder of 11110 to be sole remaining candidate
			// and prefix is 11, so 11<<3+110 = 30
			res = res<<(bitLen-1-shift) + candidates[0]
			break
		}
	}

	return res
}

func getLifeSupportRates(bitLen int, nums []int) (int, int) {
	o2Rate := getO2Rate(bitLen, nums)
	co2Rate := getCO2Rate(bitLen, nums)

	return o2Rate, co2Rate
}

func getLifeSupportRatesStdin() int {
	scanner := bufio.NewScanner(os.Stdin)

	nums := []int{}
	bitLen := 0
	for scanner.Scan() {
		if bitLen == 0 {
			bitLen = len(scanner.Text())
		}
		num, _ := strconv.ParseInt(scanner.Text(), 2, 0)
		nums = append(nums, int(num))

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	o2, co2 := getLifeSupportRates(bitLen, nums)
	return o2 * co2
}

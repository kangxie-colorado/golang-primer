package main

import (
	"fmt"
	"math"
)

func main() {

	mySin(207)
	var d float64
	for d = -180; d <= 3660; d += 0.1 {
		myResult := mySin(float64(d))
		mathLibResult := math.Sin(float64(d) * math.Pi / 180)
		diff := myResult - mathLibResult
		fmt.Println(d, ":", myResult, "vs", mathLibResult, "diff is", diff)
		if math.Abs(diff) > 0.000001 {
			panic("Wrong!")
		}
	}

}

func myFact(n float64) float64 {
	var prod float64 = 1
	for n > 0 {
		prod *= n
		n--
	}

	return prod
}

func mySin(x float64) float64 {
	var sum float64 = 0
	var item float64 = math.MaxFloat64
	xIntPart := int64(x)
	xFracPart := x - float64(xIntPart)

	xIntPart %= 360
	x = float64(xIntPart) + xFracPart

	arcX := x * math.Pi / 180

	for i := 0; item < -0.0000001 || item > 0.0000001; i++ {
		var sign float64 = 1
		if int(i)%2 == 1 {
			sign = -1
		}

		itemPow := math.Pow(arcX, float64(2*i+1))
		itemFac := float64(myFact(float64(2*i + 1)))

		item = sign * (itemPow / itemFac)
		sum += item
	}

	return sum
}

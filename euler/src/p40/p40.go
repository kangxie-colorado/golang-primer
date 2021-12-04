/***
An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

***/

package main

import (
	"fmt"
	"strconv"
)

// "0" : 48

func main() {

	strRepr := "0"
	for num := 1; len(strRepr) < 1000001; num++ {
		strRepr += strconv.Itoa(num)
	}

	fmt.Println(strRepr[1]-48, strRepr[10]-48, strRepr[100]-48, strRepr[1000]-48, strRepr[10000]-48, strRepr[100000]-48, strRepr[1000000]-48)

}

package main

var num_to_str = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

func number_to_words(num int) string {
	str := ""

	if num >= 1000 {
		n := num / 1000
		str += num_to_str[n] + num_to_str[1000]
		num = num % 1000
	}

	if num >= 100 {
		n := num / 100
		str += num_to_str[n] + num_to_str[100]
		num = num % 100
		if num != 0 {
			str += "and"
		}
	}

	if num >= 20 {
		n := num / 10
		str += num_to_str[n*10]
		num = num % 10

	}

	if num != 0 {
		str += num_to_str[num]
	}

	return str
}

func main() {
	str := ""
	for n := 1; n <= 1000; n++ {
		str += number_to_words(n)
	}
	println(str)
	print(len(str))
}

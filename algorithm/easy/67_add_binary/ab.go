package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	var result string
	var sum, adv int
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 {
		sum = adv
		if i >= 0 {
			sum += int(a[i]) - int('0')
			i--
		}
		if j >= 0 {
			sum += int(b[j]) - int('0')
			j--
		}
		result = strconv.Itoa(sum%2) + result
		adv = sum / 2
	}
	if adv == 1 {
		result = "1" + result
	}
	return result
}

func addBinarySuccsee(a string, b string) string {
	var max, min int
	var adv bool
	var tmp, result string
	if len(a) > len(b) {
		max, min = len(a), len(b)
		tmp = a
	} else {
		max, min = len(b), len(a)
		tmp = b
	}

	for i := 0; i < min; i++ {
		if (string(a[len(a)-1-i]) == "1") && (string(b[len(b)-1-i]) == "1") { //11
			if adv {
				result = "1" + result
			} else {
				result = "0" + result
			}
			adv = true
		} else if (string(a[len(a)-1-i]) == "0") && (string(b[len(b)-1-i]) == "0") { //00
			if adv {
				result = "1" + result
			} else {
				result = "0" + result
			}
			adv = false
		} else { //10, 01
			if adv {
				result = "0" + result
				adv = true
			} else {
				result = "1" + result
				adv = false
			}
		}
	}

	for i := max - min - 1; i >= 0; i-- {
		//1
		if string(tmp[i]) == "1" {
			if adv {
				result = "0" + result
			} else {
				result = "1" + result
				adv = false
			}
		} else {
			if adv {
				result = "1" + result
			} else {
				result = "0" + result
			}
			adv = false
		}
	}

	if adv {
		result = "1" + result
	}

	return result
}

// this limit is int64's limit. problem accept over the limit number...
func addBinaryFail(a string, b string) string {
	ca, _ := strconv.ParseInt(a, 2, 64)
	cb, _ := strconv.ParseInt(b, 2, 64)
	fmt.Println("ca:", ca)
	fmt.Println("cb:", cb)

	result := fmt.Sprintf("%b", ca+cb)

	return result
}

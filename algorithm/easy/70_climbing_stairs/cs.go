package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

// use combination
func climbStairs(n int) int {
	var x, y, min, total, result int //x:1step,y:2step
	//var tt, tm uint
	var tmp uint
	if n == 1 {
		return 1
	}
	if n%2 == 0 { //even
		x = 0
	} else { // odd
		x = 1
	}
	for ; x <= n; x += 2 {
		//tt, tm = 1, 1
		y = (n - x) / 2
		total = x + y

		if x < y && x != 0 {
			min = x
		} else {
			min = y
		}

		tmp = 1
		for i := 1; i <= min; i++ {
			tmp *= uint(total)
			total--
			tmp /= uint(i)
		}

		result += int(tmp)
	}
	return result
}

// use permutation
func climbStairsFailed(n int) int {
	var x, y, xy, result int //x:1step,y:2step
	var flim, frem int = 1, 1
	var lim, rem int
	if n == 1 {
		return 1
	}
	if n%2 == 0 { //even
		x = 0
	} else { // odd
		x = 1
	}
	for ; x <= n; x += 2 {
		flim, frem = 1, 1
		y = (n - x) / 2

		xy = x + y

		if x > y {
			lim, rem = x, y
		} else {
			lim, rem = y, x
		}

		for i := xy; lim < i; i-- {
			flim *= int(i)
		}
		for i := 1; i <= rem; i++ {
			frem *= int(i)
		}
		result += int(flim / frem)
		fmt.Println("x:", x, ",y:", y, ",xy:", xy, ",flim:", flim)
	}
	fmt.Println("x:", x, ",y:", y, ",xy:", xy, ",result:", result)

	return result
}

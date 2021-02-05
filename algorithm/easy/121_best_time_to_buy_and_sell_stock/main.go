package main

func maxProfit(prices []int) int {
	var max, min int = 0, prices[0]

	for _, p := range prices {
		if p < min {
			min = p
			continue
		}
		tmp := p - min
		if tmp > max {
			max = tmp
		}
	}
	if max > 0 {
		return max
	}
	return 0
}

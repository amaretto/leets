package main

func maxArea(height []int) int {
	var il, jl, max, tmp int
	for i := 0; i < len(height)-1; i++ {
		il = height[i]
		if il*(len(height)-1-i) < max {
			continue
		}

		for j := i + 1; j < len(height); j++ {
			jl = height[j]
			if il > jl {
				tmp = jl * (j - i)
			} else {
				tmp = il * (j - i)
			}
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

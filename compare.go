package main

func MaxInt(nums ...int) int {
	max := -Max * 100
	for _, n := range(nums) {
		if n > max {
			max = n
		}
	}
	return max
}

func MinInt(nums ...int) int {
	min := Max * 100
	for _, n := range(nums) {
		if n < min {
			min = n
		}
	}
	return min
}
package main

func Mix(nums []int) []int {
	n := len(nums) / 2
	x := nums[:n]
	y := nums[n:]
	result := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		result = append(result, x[i])
		result = append(result, y[i])
	}
	return result
}

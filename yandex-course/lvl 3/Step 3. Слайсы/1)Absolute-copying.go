package main

func SliceCopy(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	copy(result, nums)
	return result
}

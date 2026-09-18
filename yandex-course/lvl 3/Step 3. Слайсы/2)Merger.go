package main

func Join(nums1, nums2 []int) []int {
	result := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))
	copy(result, nums1)
	copy(result[len(nums1):], nums2)
	return result
}

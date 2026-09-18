package main

import (
	"errors"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	result := []int{}
	if n <= 0 {
		err := errors.New("n = 0")
		return nil, err
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < limit {
			result = append(result, nums[i])
			if len(result) >= n {
				return result, nil
			}
		}
	}
	return result, nil
}

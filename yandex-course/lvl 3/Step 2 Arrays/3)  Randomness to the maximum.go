package main

func FindMaxMinInArray(array [10]int) (int, int) {
	max := array[0]
	min := array[0]

	for _, number := range array {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	return max, min
}

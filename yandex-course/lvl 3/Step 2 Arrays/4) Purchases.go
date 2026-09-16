package main

func SumOfArray(array [6]int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

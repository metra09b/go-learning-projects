package main

func FiveSteps(array [5]int) [5]int {
	var reversed [5]int
	reversed[0] = array[4]
	reversed[1] = array[3]
	reversed[2] = array[2]
	reversed[3] = array[1]
	reversed[4] = array[0]
	return reversed
}

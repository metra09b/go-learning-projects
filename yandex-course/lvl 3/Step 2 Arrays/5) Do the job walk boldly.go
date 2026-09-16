package main

import (
	"fmt"
)

func PrettyArrayOutput(array [9]string) {
	i := 1
	for _, number := range array {
		if i < 8 {
			fmt.Printf("%d я уже сделал: %s\n", i, number)
			i++
		} else {
			fmt.Printf("%d не успел сделать: %s\n", i, number)
			i++
		}
	}
}

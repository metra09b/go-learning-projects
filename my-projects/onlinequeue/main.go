package main

import (
	"fmt"
)

func main() {
	p1 := "-"
	p2 := "-"
	p3 := "-"
	p4 := "-"
	p5 := "-"

	for {

		name := ""
		place := 0

		n, _ := fmt.Scanln(&name, &place)
		if n == 0 {
			continue
		}
		if name == "очередь" || name == "количество" || name == "конец" {
			if name == "конец" {
				fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", p1, p2, p3, p4, p5)
				break
			} else if name == "очередь" {
				fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", p1, p2, p3, p4, p5)

			} else if name == "количество" {
				freePlace := 0
				noFreePlace := 0
				if p1 == "-" {
					freePlace++
				} else {
					noFreePlace++
				}
				if p2 == "-" {
					freePlace++
				} else {
					noFreePlace++
				}
				if p3 == "-" {
					freePlace++
				} else {
					noFreePlace++
				}
				if p4 == "-" {
					freePlace++
				} else {
					noFreePlace++
				}
				if p5 == "-" {
					freePlace++
				} else {
					noFreePlace++
				}
				fmt.Printf("Осталось свободных мест: %d\n", freePlace)
				fmt.Printf("Всего человек в очереди: %d\n", noFreePlace)
			}
		} else { //заходим если не команда
			if n < 2 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
				continue
			} else if place < 1 || place > 5 {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", place)
				continue
			} else if p1 != "-" && p2 != "-" && p3 != "-" && p4 != "-" && p5 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", place)
				continue
			} else {
				switch place {
				case 1:
					if p1 != "-" {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					} else {
						p1 = name
					}
				case 2:
					if p2 != "-" {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					} else {
						p2 = name
					}
				case 3:
					if p3 != "-" {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					} else {
						p3 = name
					}
				case 4:
					if p4 != "-" {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					} else {
						p4 = name
					}
				case 5:
					if p5 != "-" {
						fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", place)
					} else {
						p5 = name
					}
				}
			}
		}
	}
}

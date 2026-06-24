package main

import "fmt"

func main() {
	var age int
	var name string
	var passportSeriesAndNumber string

	_, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)
	if err != nil {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	if age < 14 || age > 150 {
		fmt.Println("Ошибка: невалидный возраст")
		return
	}

	if len(name) < 2 {
		fmt.Println("Ошибка: невалидное имя")
		return
	}

	if len(passportSeriesAndNumber) != 10 {
		fmt.Println("Ошибка: невалидная серия и номер паспорта")
		return
	}

	fmt.Println(fmt.Sprintf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s", name, age, passportSeriesAndNumber))
}

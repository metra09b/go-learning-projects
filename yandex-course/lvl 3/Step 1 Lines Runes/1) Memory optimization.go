package main

import (
	"fmt"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	name := fmt.Sprintf("%s%s", first, second)
	a := len(name)
	b := utf8.RuneCountInString(name)
	return fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", name, a, b)
}

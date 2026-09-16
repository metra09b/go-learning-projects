package main

func NumbersToLetters(input string) string {
	result := ""

	for _, ch := range input {
		switch ch {
		case '0':
			result += "ноль"
		case '1':
			result += "один"
		case '2':
			result += "два"
		case '3':
			result += "три"
		case '4':
			result += "четыре"
		case '5':
			result += "пять"
		case '6':
			result += "шесть"
		case '7':
			result += "семь"
		case '8':
			result += "восемь"
		case '9':
			result += "девять"
		case '+':
			result += "плюс"
		case '-':
			result += "минус"
		case '*':
			result += "умножить на"
		case '/':
			result += "разделить на"
		default:
			result += string(ch)
		}
	}

	return result
}

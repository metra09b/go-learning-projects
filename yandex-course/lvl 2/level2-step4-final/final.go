package main

import (
	"errors"
	"strings"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {
	numberDay := int(TimeNow().Weekday())
	switch numberDay {
	case 0:
		return "Воскресенье"
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	}
	return ""
}

func dayOrNight() string {
	t := TimeNow().Hour()
	if t >= 10 && t <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	numberDay := int(TimeNow().Weekday())
	if numberDay <= 5 {
		return 5 - numberDay
	}
	return 6
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	dayNow := currentDayOfTheWeek()
	answer = strings.ToLower(answer)
	dayNow = strings.ToLower(dayNow)
	if answer == dayNow {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	day := dayOrNight()
	day = strings.ToLower(day)
	answer = strings.ToLower(answer)
	if day == answer {
		return true, nil
	}
	return false, nil
}

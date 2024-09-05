package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	var a, b, c int
	isRomans := false

	line, _ := in.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	words := strings.Split(line, " ")

	// Проверим что у нас задана правильная входящая строка
	if len(words) != 3 {
		panic("строка не является математической операцией")
	}

	// Проверим что у нас задана нужная математическая операция
	var IsCorrectOperation = regexp.MustCompile(`[+-/*]`).MatchString

	if !IsCorrectOperation(words[1]) {
		panic("математическая операция некоректная")
	}

	if IsRomanNumerals(words[0]) && IsRomanNumerals(words[2]) {
		isRomans = true
		i := RtoA(words[0])
		a = int(i)
		i = RtoA(words[2])
		b = int(i)

	} else {
		// Проверим допустимые значения чисел
		err := errors.New("")
		a, err = strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}

		b, err = strconv.Atoi(words[2])
		if err != nil {
			panic(err)
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Недопустимое значение чисел")
	}

	switch words[1] {
	case "+":
		c = a + b
	case "-":
		c = a - b
	case "*":
		c = a * b
	case "/":
		c = a / b
	}

	if isRomans {
		if c <= 0 {
			panic("Не может быть отрицательные числа и ноль")
		}
		fmt.Println("Ответ: ", AtoR(c))

	} else {
		fmt.Println("Ответ: ", c)
	}
}

// Процедуры работы с римскими цыфрами

var romans = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func IsRomanNumerals(in string) bool {
	_, found := romans[in]

	return found
}

func AtoR(in int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for in >= conversion.value {
			roman.WriteString(conversion.digit)
			in -= conversion.value
		}
	}

	return roman.String()
}

func RtoA(in string) int {
	rez, _ := romans[in]
	return rez
}

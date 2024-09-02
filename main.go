package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	romans "github.com/summed/goromans"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	var a, b, c int
	isRomans := false

	line, _ := in.ReadString('\n')
	line = strings.TrimRight(line, "\n")

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

	if romans.IsRomanNumerals(words[0]) && romans.IsRomanNumerals(words[2]) {
		isRomans = true
		i, _ := romans.RtoA(words[0])
		a = int(i)
		i, _ = romans.RtoA(words[2])
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
		fmt.Println("Ответ: ", romans.AtoR(uint(c)))

	} else {
		fmt.Println("Ответ: ", c)
	}
}

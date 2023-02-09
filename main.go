package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	value   int
	isRoman bool
}

type payload struct {
	num1, num2 number
	op         rune
}

var romanDict = map[string]int{
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

func roman(number int) string {
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

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}
func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func isOperator(r rune) bool {
	switch r {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}
func parseNumber(str string) (*number, error) {

	if val, err := strconv.Atoi(str); err != nil {
		if val := romanDict[str]; val != 0 {
			return &number{val, true}, nil
		} else {
			return nil, errors.New("Неверное число")
		}
	} else {
		return &number{val, false}, nil
	}
}
func parseString(inputString string) (*payload, error) {

	unSpaceString := spaceMap(inputString)
	operatorIndex := strings.IndexFunc(unSpaceString, isOperator)
	if operatorIndex == -1 {
		return nil, errors.New("Строка не содержит оператор")
	}
	leftSubString := unSpaceString[:operatorIndex]
	rightSubString := unSpaceString[operatorIndex+1:]
	op := rune(unSpaceString[operatorIndex])
	num1, err := parseNumber(leftSubString)
	if err != nil {
		return nil, fmt.Errorf("Не удалось сконвертировать первое число %s ошибка %s", leftSubString, err)
	}
	num2, err := parseNumber(rightSubString)
	if err != nil {
		return nil, fmt.Errorf("Не удалось сконвертировать второе число %s ошибка %s", rightSubString, err)
	}
	return &payload{*num1, *num2, op}, nil
}
func validatePayLoad(pl payload) error {
	if pl.num1.isRoman != pl.num2.isRoman {
		return errors.New("Системы счистления не соответствуют друг другу")
	}
	if (pl.num1.value > 10) || (pl.num1.value < 1) {
		return fmt.Errorf("Первое число %d не входит в границы от 1 до 10", pl.num1.value)
	}
	if (pl.num2.value > 10) || (pl.num2.value < 1) {
		return fmt.Errorf("Второе число %d не входит в границы от 1 до 10", pl.num2.value)
	}
	return nil
}
func estimatePayLoad(pl payload) (int, error) {
	switch pl.op {
	case '+':
		return pl.num1.value + pl.num2.value, nil
	case '-':
		return pl.num1.value - pl.num2.value, nil
	case '/':
		return pl.num1.value / pl.num2.value, nil
	case '*':
		return pl.num1.value * pl.num2.value, nil
	default:
		return 0, errors.New("Неверный оператор")
	}
}
func report(answer int, isRoman bool) error {
	if isRoman {
		if answer < 1 {
			return errors.New("Римские цифры не могут быть меньше еденицы")
		}
		_, err := fmt.Print(roman(answer))
		if err != nil {
			return err
		}
	} else {
		_, err := fmt.Print(answer)
		if err != nil {
			return err
		}
	}
	return nil
}
func main() {
	_, err := fmt.Print("Введите выражение:\n")
	if err != nil {
		panic(err)
	}
	var inputString string
	reader := bufio.NewReader(os.Stdin)
	inputString, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	payload, err := parseString(inputString)
	if err != nil {
		panic(err)
	}
	if err := validatePayLoad(*payload); err != nil {
		panic(err)
	}
	answer, err := estimatePayLoad(*payload)
	if err != nil {
		panic(err)
	}
	err = report(answer, payload.num1.isRoman)
	if err != nil {
		panic(err)
	}
}

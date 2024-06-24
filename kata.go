package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rim_to_arab = map[string]int{
	"I":    1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func main() {
	fmt.Println("Задание для вступления в Kata")
	fmt.Println("Введите выражение в формате x + y или exit для выхода:")

	for {
		getting_an_expression := bufio.NewReader(os.Stdin)
		resulting_expression, _ := getting_an_expression.ReadString('\n')
		resulting_expression = strings.TrimSpace(resulting_expression)

		if resulting_expression == "exit" {
			fmt.Println("Вы покинули проект")
			break
		}

		result, err := calculation(resulting_expression)
		if err != nil {
			panic("Ошибка")
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func calculation(expression string) (string, error) {
	parts_of_expression := strings.Fields(expression)
	if len(parts_of_expression) != 3 {
		panic("Вы ввели выражение неверно")
	}

	first_num, operator, second_num := parts_of_expression[0], parts_of_expression[1], parts_of_expression[2]

	rim_first_num := check_rim_numeral(first_num)
	rim_second_num := check_rim_numeral(second_num)

	if rim_first_num != rim_second_num {
		panic("Не смешивайте две системы исчисления")
	}

	var a, b int
	var err error
	if rim_first_num {
		a, b, err = parseRomanOperands(first_num, second_num)
	} else {
		a, b, err = parseArabicOperands(first_num, second_num)
	}
	if err != nil {
		return "", err
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль невозможно")
		}
		result = a / b
	default:
		panic("Неизвестный оператор")
	}

	if rim_first_num {
		if result <= 0 {
			panic("Римские числа должны быть только положительными")
		}
		return in_rim_answer(result), nil
	}

	return strconv.Itoa(result), nil
}

func parseArabicOperands(first_num, second_num string) (int, int, error) {
	a, err := strconv.Atoi(first_num)
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(second_num)
	if err != nil {
		return 0, 0, err
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Не соответствует диапазону [1:10]")
	}

	return a, b, nil
}

func parseRomanOperands(first_num, second_num string) (int, int, error) {
	a, check := rim_to_arab[first_num]
	if !check {
		panic("Неправильно набрано римское число")
	}
	b, check := rim_to_arab[second_num]
	if !check {
		panic("Неправильно набрано римское число")
	}

	return a, b, nil
}

func in_rim_answer(num int) string {
	if num <= 0 || num > 3999 {
		return ""
	}

	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var rim strings.Builder

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			rim.WriteString(syb[i])
		}
	}

	return rim.String()
}

func check_rim_numeral(expression string) bool {
	_, check := rim_to_arab[expression]
	return check
}

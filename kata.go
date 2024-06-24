package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rim_to_arab = map[string]int {
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
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
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func calculation(expression string) (string, error) {
	parts_of_expression := strings.Fields(expression)
	if len(parts_of_expression) != 3 {
		return "", fmt.Errorf("Вы ввели выражение неверно")
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
		a, b, err = parseOperands(first_num, second_num)
	}
	if err != nil {
		return "", err
	}


	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", fmt.Errorf("Деление на ноль невозможно")
		}
		result = num1 / num2
	default:
		panic("Неизвестный оператор: ")
	}

	return fmt.Sprintf("%d", result), nil
}

func check_rim_numeral (expression string) bool {
	_, check := rim_to_arab[expression]
	return check
}
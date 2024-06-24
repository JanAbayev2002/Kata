package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	num1, err1 := strconv.Atoi(first_num)
	num2, err2 := strconv.Atoi(second_num)

	if err1 != nil || err2 != nil {
		return "", fmt.Errorf("Не удалось преобразовать операнды в числа")
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
		return "", fmt.Errorf("Неизвестный оператор: %s", operator)
	}

	return fmt.Sprintf("%d", result), nil
}

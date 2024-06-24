package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Задание для втупление в Kata\n")
	fmt.Println("Введите выражение в формате x + y или exit для выхода:")

	getting_an_expression := bufio.NewReader(os.Stdin)
	resulting_expression, _ := getting_an_expression.ReadString('\n')
	resulting_expression = strings.TrimSpace (resulting_expression)

	if resulting_expression == "exit" {
		fmt.Println ("Вы покинули проект")
	}

func calculation (expression string)(string) {

	parts_of_expression := strings.Fields (expression)
	if len(parts_of_expression) != 3 {
		panic("Вы не вверно ввели выражение")
	}	

	first_num, operator, second_num := parts_of_expression[0], parts_of_expression[1], parts_of_expression[2]

	
}


}

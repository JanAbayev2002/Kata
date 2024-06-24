package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Задание для втупление в Kata\n")
	fmt.Println("Введите выражение в формате x + y:")
	getting_an_expression := bufio.NewReader(os.Stdin)
	resulting_expression, _ := getting_an_expression.ReadString('\n')
	fmt.Println(resulting_expression)
}

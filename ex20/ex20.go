/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	words := strings.Split(str, " ")

	// Case 1
	fmt.Println(words)
	var res []string
	for i := len(words) - 1; i >= 0; i-- {
		res = append(res, words[i])
	}
	length := len(words) - 1
	for i := 0; i < length/2; i++ {
		words[i], words[length-i] = words[length-i], words[i]
	}
	fmt.Printf("Строка с перевернутыми словами №1: %v\n", res)
	fmt.Printf("Строка с перевернутыми словами №2: %v", words)
}

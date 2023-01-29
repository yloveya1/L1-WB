/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse1(str string) string {
	var res string
	for _, val := range str {
		res = string(val) + res
	}
	return res
}

func reverse2(str string) string {
	length := len(str) - 1
	buffer := []rune(str)
	for i := 0; i < length/2; i++ {
		buffer[i], buffer[length-i] = buffer[length-i], buffer[i]
	}
	return string(buffer)
}

func main() {
	fmt.Println("Введите строку:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Исходная строка = %v\n", str)
	fmt.Printf("Перевернутая строка №1 = %v\n", reverse1(str))
	fmt.Printf("Перевернутая строка №2 = %v\n", reverse1(str))

}

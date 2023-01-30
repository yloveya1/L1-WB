/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse1(str string) string {
	var res string
	for _, val := range str {
		res = string(val) + res // реализуем добавление символов строки str в начало результирующей строки
	}
	return res
}

func reverse2(str string) string {
	length := len(str) - 1
	buffer := []rune(str) // Преобразуем строку в массив рун
	for i := 0; i < length/2; i++ {
		buffer[i], buffer[length-i] = buffer[length-i], buffer[i] // Проходим по массиву с начала и конца и меняем местами символыф
	}
	return string(buffer) // конвертиурем обратно в строку
}

func main() {
	fmt.Println("Введите строку:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') // Читаем строку со стандортного потока ввода до символа переноса строки
	str = strings.Trim(str, "\r\n")                      // Удаляем символ возрата коретки и переноса строки
	fmt.Printf("Исходная строка = %v\n", str)
	fmt.Printf("Перевернутая строка №1 = %s\n", reverse1(str))
	fmt.Printf("Перевернутая строка №2 = %s", reverse1(str))

}

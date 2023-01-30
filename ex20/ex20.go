/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') // Читаем строку со стандортного потока ввода до символа переноса строки
	words := strings.Fields(str)                         // Fields разбивает str вокруг каждого экземпляра одного или нескольких последовательных символов пробела

	// Case 1

	fmt.Println(words)
	var res []string
	for i := len(words) - 1; i >= 0; i-- {
		res = append(res, words[i]) // идем с конца массива и добавляем в результирующий массив слова в обратном порядке
	}

	// Case 2

	length := len(words) - 1
	for i := 0; i < length/2; i++ {
		words[i], words[length-i] = words[length-i], words[i]
	}

	// конвертируем массив строк в строку

	fmt.Printf("Строка с перевернутыми словами №1: %v\n", strings.Join(res, " "))
	fmt.Printf("Строка с перевернутыми словами №2: %v", strings.Join(words, " "))
}

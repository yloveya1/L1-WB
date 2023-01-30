/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}
*/
package main

import (
	"fmt"
	"strings"
)

// Неоправданное использоавние глобальной переменной
//var justString string

func main() {
	fmt.Println(someFunc())
	fmt.Println("================================")
	fmt.Println(editedFunc())
}

func createHugeString(size int) (result string) {
	var str strings.Builder
	for i := 0; i < size; i++ {
		str.WriteString("д")
	}
	return str.String()
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100] // unicode символы занимают 2 байта, можем получить не тот результат, который ждали
}

func editedFunc() string {
	v := createHugeString(1 << 10)
	tmp := []rune(v)         // создаем массив рун для корректной реализации среза
	return string(tmp[:100]) // конвертируем массив обратно в строку
}

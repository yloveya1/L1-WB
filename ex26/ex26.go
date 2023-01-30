// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Printf("%v - %v\n", s1, CheckUnique(s1))
	fmt.Printf("%v - %v\n", s2, CheckUnique(s2))
	fmt.Printf("%v - %v\n", s3, CheckUnique(s3))
}

func CheckUnique(str string) bool {
	str = strings.ToLower(str) // переводим строку в нижний регистр для обеспечения регистронезависимости
	m := make(map[rune]struct{})
	for _, s := range str {
		_, ok := m[s] // проверка на наличие элемента с таким ключом в мапе
		if ok {
			return false // если ключ с таким значением уже есть в мапе, значит символ в строке не уникальный
		}
		m[s] = struct{}{}
	}
	return true
}

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})
	var result []string

	for _, val := range str {
		// итерируемся по массиву и записываем значения в качестве ключей мапы, повторяющиеся элементы будут перезаписаны
		m[val] = struct{}{} // пустая структура, чтобы не занимать место
	}

	for key := range m {
		result = append(result, key) // добавляем ключи в результирующий массив
	}

	fmt.Println(result)
}

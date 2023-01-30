// Удалить i-ый элемент из слайса.
package main

import (
	"fmt"
)

func main() {
	mass := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var i int
	for {
		fmt.Println("Введите номер элемента для удаления:")
		_, err := fmt.Scan(&i)
		if err != nil || (i < 0 || i > len(mass)-1) { // Проверка на некорректное число и выход за границы массива
			fmt.Println("Проверьте корректность введнных данных и повторите попытку")
			continue
		}
		break
	}
	fmt.Printf("Исходный массив: %v\n", mass)
	mass = append(mass[:i], mass[i+1:]...) // объединение двух слайсов без i-го элемента
	fmt.Println(mass)
	fmt.Println("===============================")
	mass = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(mass[i:], mass[i+1:]) // делаем копирование, тем самым затирая i-ый элемент
	mass = mass[:len(mass)-1]  // срезаем последний элемент

	fmt.Println(mass)
}

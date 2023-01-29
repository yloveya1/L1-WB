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
		if err != nil || (i < 0 || i > len(mass)-1) {
			fmt.Println("Проверьте корректность введнных данных и повторите попытку")
			continue
		}
		break
	}
	mass = append(mass[:i], mass[i+1:]...)
	fmt.Println(mass)
	fmt.Println("===============================")
	mass = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(mass[i:], mass[i+1:])
	mass = mass[:len(mass)-1]

	fmt.Println(mass)
}

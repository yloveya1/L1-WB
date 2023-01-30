/*Реализовать пересечение двух неупорядоченных множеств*/
package main

import "fmt"

func Intersect(mass1 []int, mass2 []int) (c []int) {
	res := make(map[int]struct{})
	for _, val := range mass1 { // делаем значения 1-го множества ключами мапы
		res[val] = struct{}{}
	}
	for _, val := range mass2 {
		if _, ok := res[val]; ok != false { // если элементы 2-го множества являются ключами мапы, добавляем их в результирующий массив
			c = append(c, val)
		}
	}
	return c
}

func main() {
	mass1 := []int{1, 2, 3, 4, 5, 888}
	mass2 := []int{5, 4, 3, 1, 9, 888}
	fmt.Println(Intersect(mass1, mass2))
}

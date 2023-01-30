/*Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.*/
package main

import (
	"fmt"
	"math/big"
)

// Пакет math/big позволяет выполнять математические операции, которые включают в себя очень большие
// целочисленные вычисления, выходящие за пределы всех доступных примитивных типов данных.

func main() {
	var s1, s2 string
	a, b := new(big.Float), new(big.Float)
	for {
		fmt.Println("Введите первое число:")
		fmt.Scan(&s1)
		a, _ = a.SetString(s1)
		if a == nil {
			a = new(big.Float) // Так как в случае ошибки a возвращает nil
			fmt.Println("Проверьте корректность введнных данных и повторите попытку")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите второе число:")
		fmt.Scan(&s2)
		b, _ = b.SetString(s2)
		if b == nil {
			b = new(big.Float)
			fmt.Println("Проверьте корректность введнных данных и повторите попытку")
			continue
		}
		break
	}

	res := new(big.Float)
	fmt.Println(res.Add(a, b)) // Сложение
	fmt.Println(res.Sub(a, b)) // Вычитание
	fmt.Println(res.Mul(a, b)) // Умножение
	fmt.Println(res.Quo(a, b)) // Деление
}

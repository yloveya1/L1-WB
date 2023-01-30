/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/
package main

import (
	"fmt"
	"strconv"
)

func SetBit(temp *int64, index int, bit int) {
	if bit == 1 { // устанавливаем бит
		*temp = (*temp) | (1 << index)
	} else if bit == 0 { // сбрасываем бит
		*temp = (*temp) & ^(1 << index)
	}
}

func main() {
	var temp int64
	var bit, index int
	for {
		fmt.Println("Введите число:")
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Проверьте корректность введенных данных и повторите попытку!")
			continue
		}
		break
	}
	for {
		fmt.Println("Введите значение бита (0 или 1):")
		_, err := fmt.Scan(&bit)
		if err != nil || (bit != 0 && bit != 1) {
			fmt.Println("Проверьте корректность введенных данных и повторите попытку!")
			continue
		}
		break
	}
	for {
		fmt.Println("Введите номер бита:")
		_, err := fmt.Scan(&index)
		if err != nil {
			fmt.Println("Проверьте корректность введенных данных и повторите попытку!")
			continue
		} else if index > len(strconv.FormatInt(int64(temp), 2))-1 || index < 0 {
			fmt.Println("Номер бита выходит за границы числа в двоичном представлении")
			continue
		}
		break
	}
	fmt.Printf("Число в двоичной системе счисления до установки бита: %b\n", temp)
	SetBit(&temp, index, bit)
	fmt.Printf("Число в двоичной системе счисления после установки бита: %b\n", temp)
}

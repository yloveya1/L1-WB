/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).*/

package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) printer() {
	fmt.Println("Name = ", h.name)
}

// При передачи структуры Human в струкстуру Action, в Action встраиваются поля и методы от Human

type Action struct {
	age int
	Human
}

func (a *Action) printer() {
	fmt.Println("age = ", a.age, "name = ", a.Human.name)
}
func main() {
	a := Action{
		age:   18,
		Human: Human{"Yuliya"},
	}
	fmt.Println(a.name)
	a.printer()
	a.Human.printer()

}

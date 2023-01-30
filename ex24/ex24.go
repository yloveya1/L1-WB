/*
Разработать программу нахождения расстояния между двумя точками,которые представлены
в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point { // функция создания новой точки
	return &Point{
		x: x,
		y: y,
	}
}

// Геттер и сеттер для записи и чтения координат точки

func (p *Point) GetPoint() (float64, float64) {
	return p.x, p.y
}

func (p *Point) SetPoint(x, y float64) {
	p.x = x
	p.y = y
}

func Distance(p1 Point, p2 Point) float64 { // Функция для определения рассточния между двумя точками
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	a := NewPoint(1, 2)
	b := NewPoint(3, 4)
	fmt.Println(Distance(*a, *b))
}

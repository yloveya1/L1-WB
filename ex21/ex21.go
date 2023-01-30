//Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

type transport interface {
	drive() // интерфейс с которым мы хотим работать
}

type car struct {
}

func (c *car) drive() {
	fmt.Println("Let's go!!")
}

// объект plane - адаптируемый объект

type plane struct {
}

func (p *plane) fly() { // метод, к которому мы получим доступ с помощью адаптера
	fmt.Println("Let's fly!!")
}

// Необходиомсть адаптировать самолет, чтобы реализовать интерфейс, требуемый клиентом

type planeAdapter struct {
	plAdapt *plane
}

func (p *planeAdapter) drive() {
	p.plAdapt.fly() // реализация с помощью адаптера интерфейса, которого придерживается клиент
}

type client struct {
}

// Клиент ожидает реализацию метода drive

func (c *client) start(t transport) {
	t.drive()
}

func main() {
	client := &client{}
	c := &car{}
	p := &plane{}
	pAd := &planeAdapter{plAdapt: p}
	client.start(c)
	client.start(pAd)
}

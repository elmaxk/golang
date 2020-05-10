package main

/*
Prototype package
*/

import "fmt"

func pointers() {
	var count = int(100)
	ptr := &count
	fmt.Println("\nThe pointer is: ", *ptr)
	*ptr = 110
	fmt.Println("\nNew pointer at: ", count)
}

type Car struct {
	Model  string
	Engine int
}

func (c *Car) StartEngine() {
	fmt.Println("Driving: ", c.Model)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello", p.Name)
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woooooof!!!")
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	var coupe = new(Car)
	coupe.Model = "Porsche"
	coupe.Engine = 8
	coupe.StartEngine()

	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)

	var dog = new(Dog)
	Greet(dog)

	pointers()

}

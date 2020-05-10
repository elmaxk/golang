package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

func conditional(x interface{}) {
	if x == 1 {
		fmt.Println("x is equal to 1.")
	} else {
		fmt.Println("X is not equal to 1.")
	}
}

func switching(y interface{}) {
	switch y {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!", v)
	case string:
		fmt.Println("I'm a string!", v)
	default:
		fmt.Println("Unknown type!", v)
	}
}

func phoLoop(i interface{}) {
	for i := 0; i < 25; i++ {
		if i%2 == 0 {
			fmt.Println("Fizz")
			if i%3 == 0 {
				fmt.Println("Banging")
			} else {
				fmt.Println(i)
			}
		}
	}
}

func varLoop() {
	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}

func f() {
	fmt.Println("f function")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

type Foo struct {
	Key   string
	Value string
}

type XML struct {
	Key   string `xml:"id,attr"`
	Value string `xml:"parent>child"`
}

func main() {

	conditional(5)
	switching("bar")
	typeSwitch("yee")
	phoLoop(1)
	varLoop()
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function")

	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	f := XML{"Max K", "Hello Shabado"}
	b, _ := xml.Marshal(f)
	fmt.Println(string(b))
	json.Unmarshal(b, &f)
}

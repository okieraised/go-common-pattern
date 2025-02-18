package main

import "fmt"

func main() {
	B()
}

func B() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	A()
}

func A() {
	defer fmt.Println(3)
	defer fmt.Println(4)
}

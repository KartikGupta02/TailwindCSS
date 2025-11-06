package main

import (
	"Hello/myutil"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	myutil.PrintMessage("This is a message from myutil package.")

	// var name string = "Kartik"
	// var version = "1.0.0"
	// fmt.Println(name)
	// fmt.Println(version)

	// var money int = 67000
	// var currency = 91.48
	// fmt.Println(money)
	// fmt.Println("Currency:", currency)

	// var dimension float64 = 5.6
	// fmt.Println(dimension)

	// var decided bool = false
	// decided = true
	// fmt.Println(decided)

	// var person = 23
	// fmt.Println(person)

	// const pi = 3.14
	// fmt.Println("Value of pi:", pi)

	person := 123
	fmt.Println(person)

	var Public = "I am Public"
	var private = "I am Private"
	fmt.Println(Public)
	fmt.Println(private)

}

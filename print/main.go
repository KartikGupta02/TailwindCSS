package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.754

	// fmt.Println("Age: ", age, "Height: ", height, "Name: ", name)
	// fmt.Println("Hello World")

	// fmt.Printf("Age is %d\n", age)
	// fmt.Printf("Height is %.2f\n", height)
	// fmt.Printf("Name is %s\n", name)

	// fmt.Printf("Type of name is %T\n", name)
	// fmt.Printf("Type of age is %T\n", age)
	// fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}

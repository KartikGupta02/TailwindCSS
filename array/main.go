package main

import "fmt"

func main() {
	fmt.Println("We are learning arrays in Golang")

	// var name [5]string
	// name[0] = "Kartik"
	// name[1] = "Ankit"
	// name[2] = "Anshul"
	// name[3] = "Anurag"
	// name[4] = "Aman"

	// fmt.Println("Name of Persons are:", name)

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Numbers are:", numbers)

	// fmt.Println("Length of numbers array is:", len(numbers))
	// fmt.Println("Value of name at 2 index is:", name[2])

	var name [5]string
	name[0] = "Kartik"
	name[2] = "Ankit"
	fmt.Println("Names are:", name)
	fmt.Printf("Name array is %q\n", name)

}

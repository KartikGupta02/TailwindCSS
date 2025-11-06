package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("We are learning functions in Golang")
	simpleFunction()

	ans := add(5, 6)
	fmt.Println("The sum is:", ans)

	ans2 := mult(5, 6)
	fmt.Println("The multiplication is:", ans2)
}

// This show error in Golang
// You have to start curly braces in the same line as function declaration otherwise it will show error
// func main()
// {

// }

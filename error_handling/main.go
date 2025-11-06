package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// func divide(a, b float64) (float64, string) {
// 	if b == 0 {
// 		return 0, "division by zero is not allowed"
// 	}
// 	return a / b, "nil"
// }

func main() {
	fmt.Println("Started error handing in functions")
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error Handling")
	// }
	// fmt.Println("The answer is:", ans)

	ans, _ := divide(10, 0)
	fmt.Println("The answer is:", ans)
}

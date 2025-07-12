package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a int, b int) (result int) {
	result = a + b
	return
}

// func add1(a int, b) (result int) {         // This function has a syntax error because the second parameter does not have a type.
// 	result = a + b
// 	return
// }

func sub(a int, b int) int {
	result := a - b
	return result
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learning function in Golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("add of two number is :", ans)

	subans := sub(3, 4)
	fmt.Println("sub of two number is :", subans)

	data := multiply(3, 4)
	fmt.Println("multiplication of two numbers is :", data)
}

package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 20      // Modify the value at the address pointed to by num
	fmt.Println("Value modified by reference: ", *num)
}

func main() {
	// var num int
	num := "Laado"

	// var ptr *int
	// ptr = &num

	ptr := &num

	// fmt.Println("Num has value: ", num)
	fmt.Println("pointer contains: ", ptr)
	fmt.Println("Data contains through Pointer: ", *ptr)   // Dereferencing the pointer to get the value

	var pointer *int // Pointer is not assigned yet so it is nil
	if pointer == nil {           
		fmt.Println("Pointer is not assigned")
	}

	value := 5
	modifyValueByReference(&value)   // Passing the address of value to modifyValueByReference function
	fmt.Println("Value contains : ", value)

}

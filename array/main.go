package main

import "fmt"

func main() {
	fmt.Println("we are learning Array in Golang")

	var name1[5]string

	//  0 1 2 3 4
	name1[0] = "prince"
	name1[2] = "Aakash"

	fmt.Println("Names of Person is :", name1)

	var numbers = [8]int{1, 2, 3, 4, 5}
	fmt.Println("Number is :", numbers)

	fmt.Println("Length of Numbers array is :", len(numbers))

	fmt.Println("value of name at 2 index is :", name1[2])

	// ""
	var name[5]string
	name[2] = "Prince"
	name[0] = "aaksh"

	fmt.Println("name is :", name)
	fmt.Printf("name Array is  %q\n", name)
}
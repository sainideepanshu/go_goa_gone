package main

import (
	"fmt"
)

func main() {

	numbers1 := []int{1, 2}
	numbers1 = append(numbers1, 3, 10, 12, 13, 14, 15, 16, 17, 18, 19)
	fmt.Println("numbers1 : ", numbers1)
	fmt.Printf("numbers1 has data type : %T\n", numbers1)
	fmt.Println("Length : ", len(numbers1))

	name := []string{}
	fmt.Println("name : ", name)

	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)
	numbers = append(numbers, 6)

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))


	stock := make([]int, 0)
	fmt.Println("Slice:", stock)
	fmt.Println("Length:", len(stock))
	fmt.Println("Capacity:", cap(stock))
	
}

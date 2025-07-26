package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is ", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	data = data + 1.23
	fmt.Println("Data is ", data)
	fmt.Printf("Type of Data is %T\n", data)

	num = 123
	str := strconv.Itoa(num)  // this converts int to string meaning "integer to ASCII"
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)  // this converts string to int meaning "ASCII to integer" , it return two values, the second one is an error,first one is the converted value
	fmt.Println("number_int is ", number_int)
	number_int = number_int + 262892
	fmt.Println("number_int is ", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)   // this converts string to float64, it also returns two values, the second one is an error, first one is the converted value
	fmt.Println("number_float is ", number_float)
	fmt.Printf("Type of number_float is %T\n", number_float)

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple.orange.banana.laado"
	parts := strings.Split(data, ".")   // Split the string by the delimiter "." and return a slice of substrings
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")  // Count the occurrences of the substring "two" in the string and return the count
	fmt.Println("count: ", count)

	str = "                       Hello,              Go!   "
	trimmed := strings.TrimSpace(str)     // Remove leading and trailing whitespace from the string
	fmt.Println("trimmed: ", trimmed)

	str1 := "Laado"
	str2 := "Ji"
	result := strings.Join([]string{str1, "cutie", str2}, " ")  // Join the strings str1, "cutie", and str2 with a space as the delimiter and return the resulting string
	fmt.Println("result: ", result)
}

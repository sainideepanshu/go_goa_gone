package main

import "fmt"

func main() {

	// name <-> grade
	studentGrades := make(map[string]int)  // we cannot initialize a map with a value, we need to use make function
	//studentGrades := map[string]int{} // this is also valid

	studentGrades["Prince"] = 100
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	fmt.Println("Marks of Bob : ", studentGrades["Bob"])
	studentGrades["Bob"] = 100
	fmt.Println("new Marks of Bob : ", studentGrades["Bob"])

	delete(studentGrades, "Bob")      // delete a key-value pair from the map
	fmt.Println("Marks of Bob : ", studentGrades["Bob"]) // this will print 0 because Bob is deleted , but it is confusing since we are not checking if the key exists or not

	// if something is not present in the map, it will return the zero value of the value type,so we should check if the key exists or not to be safe

	// Checking if a key exists
	grades, exists := studentGrades["David"]          // when we try to access a key like this, it will return the value and a boolean indicating if the key exists or not
	// if the key exists, it will return the value and true, otherwise it will return the zero value of the value type and false
	fmt.Println("Grades of David : ", grades)
	fmt.Println("Davis exists : ", exists)

	// fmt.Println("Marks of David : ", studentGrades["David"])

	// Checking if a key exists
	Grades, Exists := studentGrades["Prince"]
	fmt.Println("Grades of Prince : ", Grades)
	fmt.Println("Prince exists : ", Exists)

	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}

	for index, value := range person {         // here range will return the key and value of the map
		fmt.Printf("---------Key is %s and marks is %d\n", index, value)
	}
}

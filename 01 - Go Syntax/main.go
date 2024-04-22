package main

import "fmt"

// Ex 1: Function signature template.
//
//	func FunctionName(param1 type1, param2 type2) returnType {
//		// Do something
//		return 0
//	}
func printHello() {
	fmt.Printf("Hello World")
}

func main() {

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 10
	// myArray = [20, 30, 40] // Raise error. Unexpected comma.
	fmt.Println(myArray)

	var my2ndArray = [3]int{11, 22, 33}
	fmt.Println(my2ndArray)

	// Slice
	mySlice := []int{10, 20, 30}
	fmt.Println(mySlice)

	// Append item to Slice
	mySlice = append(mySlice, 222)
	fmt.Println(mySlice)

	// Map
	// Declare Map
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	fmt.Println("Apple", myMap["apple"])
	fmt.Println("Fruits", myMap)
}

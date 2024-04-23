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

func temp() {
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

func pointer() {
	// Declare an integer variable
	x := 10

	// Declare a pointer to an integer and assign it the address of x
	var p *int = &x

	// Print the value of x and the value at the pointer p
	fmt.Println("Value of x:", x)                          // Output: Value of x: 10
	fmt.Println("Value at p (point to address):", p)       // Output: Value at address p:
	fmt.Println("Value at *p (actual pointed value):", *p) // Output: Value at p: 10
	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Type of *p: %T\n", *p)
	// Modify the value at the pointer p
	*p = 20

	// x is modified since p points to x
	fmt.Println("New value of x:", x) // Output: New value of x: 20
}

type ListNode struct {
	Value int
	Next  *ListNode
}

// Function to add a node to the front of the list
func prepend(head **ListNode, value int) {
	newNode := ListNode{Value: value, Next: *head}
	*head = &newNode
}
func pointer_link_list() {
	var head *ListNode

	prepend(&head, 10)
	prepend(&head, 20)

	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

type Employee struct {
	Name   string
	Salary int
}

// Function to give a raise to an employee
func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

func pointer_struct() {
	employee := Employee{"Pete", 333000}
	fmt.Printf("employee ssalary: %d\n", employee.Salary)

	giveRaise(&employee, 500)
	fmt.Printf("employee ssalary: %d\n", employee.Salary)

}

func main() {

	// pointer()
	// pointer_struct()
	pointer_link_list()
}

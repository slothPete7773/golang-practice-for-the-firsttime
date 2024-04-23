# Data Structure

## Data Type

**Primitive Data Type**

1. Boolean
   - `True` and `False` value
2. Numeric
   - Integer
     1. `int`
     2. `int8`
     3. `int16`
     4. `int32`
     5. `int64`
     6. `uint` Unsigned Integer, Only positive integer, Not including negative integer.
  - Float
    1. `float32`
    2. `float64`
3. `string`
4. `byte`
5. `rune`
   1. alias for int32.
   2. represents a Unicode code point
6. Complex
   1. `complex64`
   2. `complex128`

**Data Wrapper Structure**

A data structure that keep a collection/set of primitive data.

Useful when keeping a complex data model.

1. Array: 
   - Fixed-size sequence of item.
   - All items need to be of same type.
2. Slice
   - Similar to Array but dynamic in size.
3. Map
   - A set of key-value dictionary of items.
   - Fast look-up. 
4. Struct
   - Same as Struct in C. 
   - A template of group of variables.
   - Can be represent as the Object in Go. But **Go Do Not Have OOP Concept!!**
5. Function
   - Function are 1st-order data type in Go
   - It is treated as Data Type as well. Can be assigned to a variable.
6. Interface
   ```go
   type Shape interface {
        Area() float64
        Perimeter() float64
    }
   ```
   - Can be thought of the data type for abstract function in the Interface.
   - Where the Interface is a pre-defined set of method signature that variable(s) need to implement. If they will extend the interface.
   - Usually used with Struct.
7. Pointer
   - Data type for storing the memory address. An address is where the actual value of a variable is stored in the memory.
8. Channel
   - Data type for communciating with Go-routine.
   - Required for working in Concurrent thread.


## Array and Slice

**Key Note**
- Both store the sequence of items. 
- Need to be declared type.
- When passing Array and/or Slice to a function argument.
  - Array: Will pass-by-value, pass the value to function argument.
  - Slice: Will pass-by-reference, pass the reference to the value in the function argument.

**Slice need to be used with caution.**
- Each time a Slice is declared, more memory will be reserved, could lead to more memory consumption.
- Pass-by-reference of Slice can also unintentional change the value while processing.
 
### Array

- Array need to declare size. While Slice does not require.
- Cannot assign a list of number to the Array variable. As show in `Snippet-2`. Must assign to each index only.
- The size cannot change after declared.

Good things about Array.
- Fixed size of items. **Deterministic**
- Avoid Memory Leak. Because with fixed size, it won't eat more memory. 


Snippet-1
```go
// Example
var myArray [{Size}]{type}

// Declare V1
var myArrayNumber [10]int

myArray[0] = 10    // Assign values
myArray[1] = 20
myArray[2] = 30
fmt.Println(myArray) // Output: [10 20 30]

// Declare V2: With default value.
var my2ndArray = [3]int{11, 22, 33}
// Declare V3: Shorthand, With default value.
my3rdArray := [3]int{11, 22, 33}

// Looping through the array
for i := 0; i < len(myArray); i++ {
  fmt.Println(myArray[i]) // Output: 10, 20, 30
}
```

Snippet-2
```go
var myArray [3]int // An array of 3 integers
myArray[0] = 10    // Assign values
myArray[1] = 20
myArray[2] = 30

myArray = [10, 20, 30] // Raise Error. Unexpected Comma
```

### Slice 

- No size declaration required.
- Use `append()` to add or remove data.
- Declare as Array but without size, still need bracket([]) though.
- Array can be converted to Slice.

```go
// Example

// Declaration
mySlice := []int{10, 20, 30, 40, 50}

fmt.Println(mySlice)          // Output: [10 20 30 40 50]
fmt.Println(len(mySlice))     // Length of the slice: 5
fmt.Println(cap(mySlice))     // Capacity of the slice: 5

// Add item to Slice
mySlice = append()


// Convert Array to Slice
myArray := [3]int{10, 20, 30}
mySlice := myArray[:] // Convert Array to Slice by retreive all items.
mySlice = append(mySlice, 12345) // Then add a new item to Slice.
```

## Map

- Same as Map in other language, or Dictionary in Python.
- Key Value object.
- Map only allow Unique Key, using the same Key will overwrite the previous value.
- Need to declare type.
- Not fixed size.
- No order in map, no Sort.
- Pass-by-reference when passed to a function argument.
- when accessing the Map item, it actaully returns 2 items.
  - `myMap["item"] -> (value, isOk)`
  - Where the isOk indicate if the map has that value or not, as a Boolean value.
### Example 
```go
// Declare Map
myMap := make(map[string]int)

// Add key-value pairs to the map
myMap["apple"] = 5
myMap["banana"] = 10
myMap["orange"] = 8

// Access item in map
fmt.Println("Apple", myMap["apple"])

// Update value
myMap["banana"] = 12

// Delete a key-value pair
delete(myMap, "orange")

// Iterate over the map
for key, value := range myMap {
  fmt.Printf("%s -> %d\n", key, value)
}

// Checking if a key exists
val, ok := myMap["pear"]
if ok {
  fmt.Println("Pear's value:", val)
} else {
  fmt.Println("Pear not found inmap")
}
```

From the snippet above.

1. Declare Map:
   1. `make` is a command for allocating memory when create this Map. Ensuring that this Map variable will have memory reserved for it.
   2. the `[string]` is the data type for `Key`.
   3. and the end `int` is the data type for `Value`.
2. Delete
   1. Use `delete({Map}, {Key})` to delete item from Map
3. Iterate over Map
   1. Use for loop, with `range {Map}` to iterate over items in map, not guarantee order.
4. Checking existence
   1. Accessing Map item return 2 things, `value, isExists` kind of thing, the value will have the declared type, the isExists will have boolean.


## Struct

Purpose of Struct is to keep the data that is more complex than Primitive Data Type can handle.

Note 
- Group the related variables in the same group, same Struct.
- Struct is free, apart from type.
- It pass value when passed into function argument.
- When assign Struct to a variable, it also pass new value, not reference.
- Can be treated as a primitive type.
  - Can be put into Array, Slice, and Map.
- Struct can have nested Struct, for more complex data.

```go
package main

import "fmt"

// Define a struct type
type Student struct {
	Name    string
	Weight  int
    Height  int
    Grade   string
}

func main() {
	// Create an instance of the Student struct
	var student1 Student
	student1.Name = "Mikelopster"
	student1.Weight = 60
    student1.Height = 180
    student1.Grade = "F"

	// Print struct values
	fmt.Println(student1)
}
```
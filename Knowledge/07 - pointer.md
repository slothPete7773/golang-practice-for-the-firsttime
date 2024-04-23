# Pointer 

What's the point of Pointer? For Pointing an actual address of the variable in the memory.

It is a kind of data type, specifically for the point of address in memory.

Main purpose of Pointer:

1. Mutability
   - To be able to mutate, change the value of a variable using a reference.
   - Sometimes it is required more complex use case for changing a value, like, changing a value in the Array or Slice. Or deeply nested Struct.
2. Efficient passing very-large struct
   - Passing Struct into a function is passed as a value, if it is too large, then it is a waste of memory. 
   - But passing Struct as reference, or a pointer, can save the memory that about to be allocated.

```go
// Declare an integer variable
x := 10

// Declare a pointer to an integer and assign it the address of x
var p *int = &x

// Print the value of x and the value at the pointer p
fmt.Println("Value of x:", x)  // Output: Value of x: 10
fmt.Println("Value at p:", *p) // Output: Value at p: 10

// Modify the value at the pointer p
*p = 20

// x is modified since p points to x
fmt.Println("New value of x:", x) // Output: New value of x: 20
```

From the snippet above:
1. Declare a pointer variable at `var p *int = &x`
   1. Where `*int` means that this is a **pointer**.
   2. And `&x` is the address of `x`, the `&` is the operator for getting an address of provided variable.
   3. Now, `p` is a variable
2. Call just `p` is calling for the pointed address value.
3. But calling to `*p` means calling to the value of pointed address.

Item 2, 3 is the essential to understand of Pointer. To clearly understand the distinction between **`p` pointed address** and **`p*` value at pointed address**


# Pass by Value VS Pass by Reference

Pass-by-value and pass-by-reference is a type of approach for passing value to a function or for mutating value of variables.

Sometimes, such approach is natively built within the Programming Language itself. For instance, in Java, specifically for Non-primitive type; Objects, are accessed by object reference. While the primitive types; Int, Boolean, String, are accessed by value.

In Go, *almost* all are passed by value. Pass by reference in Go are for variables such as; Pointer, Slice, Array, Map, Channel.

By theory, Go does not have pass-by-reference? But we can pass the address for Go to mutate the value at address.
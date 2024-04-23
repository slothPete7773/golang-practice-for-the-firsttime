# Function

Purpose of function. Also true in any programming language, and/or paradigm.

1. Increase Modularity
2. Increase Reusability
3. Take parameters to get a returned value.

## Function as data type in Go

In Go, function is treated as first-class. Means that function is also treated as a data type. 

It can be assign to a variable as well. Making Go a functional paradigm compatible.

This is note-worthy mentioning because some language treat Function as a special thing, that takes things to return somthing else.

## Function Signature

A function signature means how to write a function. It includes the function name, function parameters and their types, and the return type.

The signature is a set of these. 

**Function code**

Function Signature Template.
```go
func FunctionName(param1 type1, param2 type2) returnType {
	// Do something
	return 0
}
```

Function: Print something

```go
func printHello() {
    fmt.Printf("Hello World")
}
```

Function: Add 2 numbers

```go
func addition(numA int, numB int) int {
    return numA + numB
}
```




### Public and Private function

Go has a concept of Package and Module, and it control publicity of function access from outside of package with `Public` and `Private` function.

1. Public Function:
   - Function name starts with capital letter.
   - `func PrintHelloWorld() {...}`
2. Private Function
   - Function name starts with small letter.
   - `func printHelloWorld() {...}`

### Return Type of Function

- Go function declares return type after the parentheses of parameters, if omitted, the function does not return anything.
- Function can have any number of return item.
- `Naked Return`: A function can have named return and return nothing. The return will infer the variable in the function with the same name in the return declaration.
  - Should use with short function
```go
// V1: Without Return Type. Function not return anything.
func PrintName() {
    fmt.Printf("SlothPete")
}

// V2: With Return Type. Function return something with specific type.
func GetName() string {
    return "slothPete"
}

// V3: With multiple return
func SwapParams(x string, y string) (string string) {
    return y, x
}

// V4: Naked Return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

### Receiver Method

อย่างที่บอกไปว่า GO เป็นสิ่งที่ไม่มี OOP ดังนั้น Struct ที่เราใช้แทนโมเดลของก้อนสิ่งหนึ่ง ถ้าต้องการให้ Struct Student มี method ที่ใช้ร่วมกันได้ โดยใช้ Struct attributes ของแต่ละตัวเอง เพื่อ PrintFullName() ได้ทั้งคู่ และผลลัพธ์แตกต่างไปตาม attributes ของก้อน struct เอง

In Go, a **method**, is a function with **scope**. Where a scope is an accessibility to attributes of its own object, or Struct.

A receiver-argument is needed to add to a function to make it become a Method. Fullname is a Receiver Method.

Method characteristics:

1. When a method is declared, a function is declared with receiver-argument, that method is tied to a type. 
   1. Type in this case, let's say it is a type of a Struct.
2. Receiver Argument will refer to the type of Struct, to bind this method to that Struct instance.
3. The method can access to the property of the receiver instance.

Now, the Receiver-argument can be thought of as the Type of the Struct that this method will bind with.

Whenever this method is called, it can only associated with the bound Receiver-argument (Struct Type).

```go
// example
type Student struct {
    Name string
    Lastname string
}

// 2 Students
student_1 := Student{"Pete", "Prasertpol"}
student_2 := Student{"Veera", "Sloth"}

// Method with a receiver of type Student
// This method returns the full name of the student
func ({Parameter StructType}) {FunctionName}() {ReturnType} {
    ...
}
func (s Student) GetFullName() string {
	return s.Firstname + " " + s.Lastname
}

student := Student{
		Firstname: "Mike",
		Lastname:  "Lopster",
	}

// Call the FullName method on the Student instance
fullName := student.GetFullName()
fmt.Println("Full Name of the student:", fullName)

```

### Interface

อินเตอร์เฟส เป็นอีกสิ่งหนึ่ง ที่ตึขนาบคู่มาพร้อมกับ method ของ Go 
อินเตอร์เฟส ในที่นี้ คือการระบุ set of method signatures ที่ Interface นั้นจะต้องมี เพื่อที่จะสร้างฟังก์ชั่นที่รับก้อน Interface ใน parameter โดยตอนที่ส่งก้อนของ object ไปในฟังก์ชั่น สามารถส่ง object ใดก็ได้ แต่ต้องเป็น object ที่อิมพลีเม้นฟังก์ชั่นใน interface นั้น ๆ

ตัวอย่างคำอธิบายด้านล่างจะอธิบายได้คลอบคลุมกว่า [จาก](https://docs.mikelopster.dev/c/goapi-essential/chapter-2/function#receiver-method-%E0%B8%99%E0%B8%B4%E0%B8%A2%E0%B8%B2%E0%B8%A1-method-%E0%B9%83%E0%B8%99-go)

```
- interface คือการระบุ set ของ method ที่ต้องมีภายใต้ interface นั้นๆ
- interface เป็นเหมือนต้นแบบของ method ว่า method นั้นควรจะต้องมีคุณสมับบัติอะไรบ้าง (เป็นเหมือนการกำหนดต้นแบบเอาไว้)
    - เช่น สมมุติเรามี class สำหรับสร้าง สี่เหลี่ยม, สามเหลี่ยม, วงกลม
    - ทั้ง 3 class นี้ก็มีการเก็บขนาดที่แตกต่างกัน แต่เราอยากให้ทั้ง 3 class นี้ ควรมี method ที่มีคุณสมบัติในการหาพื้นที่เหมือนกัน
    - ทั้ง 3 class เลยต้องสร้าง method สำหรับการคำนวนพื้นที่ออกมา ที่ไม่เหมือนกันออกมา
    - ซึ่งการกำหนดว่าทั้ง 3 class ควรต้องมี method ในการหาพื้นที่ ออกมา เราเรียกสิ่งนี้ว่า "interface" = ระบุแค่สิ่งที่ต้องมี แต่สิ่งที่ต้องมีนั้นทำอะไรก็จะเป็นการ implement ผ่าน method อีกที
- เราเรียกคุณสมบัตินี้ว่า "Polymorphism" ใน Go คือความสามารถในการจัดการ object ให้มีหลากหลายรูปแบบจากการสืบทอดผ่าน interface ได้
- ใครเรียนรู้ OOP มา มันคือ concept เดียวกันกับการใช้ implements ในพวกภาษา OOP
```

**Rule of Interface**

- An interface may have many method signatures.
- All method signatures must be implemented in the extended Object.
- Good practice in Go. Interface should end with `-er` phrase.

```go
package main

import "fmt"

// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p Person) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}

	makeSound(dog)
	makeSound(person)
}
```
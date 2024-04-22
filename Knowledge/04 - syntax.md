# Go Syntax and basic programming

- Variable และ Basic type (integer, float, string, boolean)
- Control structure (if, else, switch, loop)
- Data structure (Array, map, struct)
- Function (parameter, method, interface)
- Pointer
- Error handling (throw error)


# Variables

4 ways to declare variables

1. Full variable
   - `var <name> <type>`: `var fullname string`
2. Full variable with default value
   - `var <name> <type> = <value>`
3. Shorthand variable with default value
   - `var <name> = <value>`
   - Type is inferred from actual type of value.
4. *Very popular* Shorthand variable with default value without `var`
   - `<name> := <value>`
   - Type is inferred from actual type of value.

**Note**

1. Variable can be assigned a new value of the same type anytime.
2. Variable can not be assigned a new value of different type.


# Operators

Regular arithmetic operators: `+ - * /`
Relational operators: `== != >= <= > <`
Logical operators: `&& || !`

# Control Structure

**Note**

Regular control structure in programming

1. Sequential: Line-by-line
2. Selection: If-else, Switch-case
3. Iteration: Iteration over a block of code, for while, do-while

## Selection

### If-Else

```go
// V1
if <condition> {
    ... // Operate on True
} else {
    ... // Operate on False
}

// V2
if <condition_1> {
    ... // Operate on True with condition 1
} else if <condition_2> {
    ... // Operate on True with condition 2
} else {
    ... // Operate on False
}
```

### Switch

```go
// V1
switch <variable> {
    case <value_1>
        ... // Operate if variable == value_1
    case <value_2>
        ... // Operate if variable == value_1
    case <value_3>
        ... // Operate if variable == value_1
    default
        ... // Operate if variable not equal to any case
}
```

### Mixed (If-Else + Switch)

```go
var score int = 62
var grade string

switch {
  case score >= 80:
    grade = "A"
  case score >= 70:
    grade = "B"
  case score >= 60:
    grade = "C"
  default:
    grade = "F"
  }
```

### Pre-processed If-Else

เพื่อย่อโค้ดให้ compact มากขึ้นเฉย ๆ โดยการรวบ assignment และ equility comparison ไว้ในบรรทัดเดียวกัน และแบ่งด้วย semi-colon (;)

```go

num1 := 5;
num2 := 10;

// V1: Without Preprocessed
sumNum := num1 + num2;
if sumNum >= 10 {
   fmt.Println("sumNum more than 10")
}

// V2: With Preprocessed
if sumNum := num1 + num2; sumNum >= 10 {
   fmt.Println("sumNum more than 10")
}
```

## Iteration

Golang only has `for` keyword, no `while` or `do while`.

For while and do-while loop, need to modify `for` to function it as `while` and `do-while`

### For loop

```go
for i := 1; i < 10; i++ {
  fmt.Printf("number: %d", i)
}
```

### Do while

```go
i := 1
for {
    fmt.Printf("number: %d\n", i)
    i++
    if i>100 {
        break
    } 
}
```

### While 

```go
for i<100 {
    fmt.Printf("Number: %d \n", i)
    i++
}
```
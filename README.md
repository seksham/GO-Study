# GO-Study

## Table of Contents

1. [Introduction to Go](#1-introduction-to-go)
2. [Go Basics](#2-go-basics)
3. [Functions and Methods](#3-functions-and-methods)
4. [Data Structures](#4-data-structures)
5. [Pointers](#5-pointers)
6. [Structs](#6-structs)
7. [Interfaces](#7-interfaces)
8. [Concurrency](#8-concurrency)
9. [Error Handling](#9-error-handling)
10. [Advanced Topics](#10-advanced-topics)
11. [Memory Management](#11-memory-management)
12. [Web Development](#12-web-development)
13. [Database Integration](#13-database-integration)
14. [Testing](#14-testing)
15. [Tools and Commands](#15-tools-and-commands)
16. [Best Practices and Patterns](#16-best-practices-and-patterns)
17. [File Operations](#17-file-operations)

## 1. Introduction to Go

Go, also known as Golang, is a statically typed, compiled programming language designed at Google. It provides excellent support for concurrent programming and emphasizes simplicity and efficiency.

## 2. Go Basics

### 2.1 Naming Conventions

Go has specific naming conventions that contribute to code readability:

- Package names: lowercase, single-word names (e.g., `time`, `http`)
- Variable names: camelCase for local, PascalCase for exported
- Constant names: PascalCase for exported, all caps with underscores for internal
- Function and method names: camelCase for unexported, PascalCase for exported
- Interface names: PascalCase, often ending with '-er' suffix
- Struct names: PascalCase for exported, camelCase for unexported

### 2.2 Variables

Variables in Go are used to store and manipulate data. Here's a comprehensive overview of variables in Go:

#### Declaration and Initialization

1. Basic declaration:
   ```go
   var name string
   var age int
   ```

2. Declaration with initialization:
   ```go
   var name string = "John"
   var age int = 30
   ```

3. Short declaration (type inference):
   ```go
   name := "John"
   age := 30
   ```

4. Multiple declarations:
   ```go
   var (
       name string
       age  int
       isStudent bool
   )
   ```

#### Zero Values

Variables declared without an explicit initial value are given their zero value:
- Numeric types: `0`
- Boolean type: `false`
- String type: `""`
- Pointer types: `nil`

#### Constants

Constants are declared using the `const` keyword:
```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

#### Type Conversion

Go requires explicit type conversion:
```go
var x int = 10
var y float64 = float64(x)
```

#### Scope

- Package-level variables: Declared outside any function
- Local variables: Declared inside a function
- Block-level variables: Declared inside a block (e.g., if statement)

#### Naming Conventions

- Use camelCase for variable names
- Exported variables (accessible from other packages) start with an uppercase letter

#### Variable Shadowing

Inner blocks can declare variables with the same name as outer blocks, shadowing the outer variable:
```go
x := 10
if true {
    x := 20 // This x shadows the outer x
    fmt.Println(x) // Prints 20
}
fmt.Println(x) // Prints 10
```

#### Unused Variables

Go does not allow unused variables. The compiler will throw an error if a variable is declared but not used.

#### Blank Identifier

The blank identifier `_` can be used to ignore values:
```go
x, _ := someFunction() // Ignores the second return value
```

Additional examples:
- [Variable Examples](variables/main.go)

This overview covers the essential aspects of variables in Go, including declaration, initialization, scope, naming conventions, and special features like the blank identifier.

### 2.3 Data Types

Go has several built-in data types:

1. Basic Types:
   - Numeric Types:
     - Integers: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
     - Floating-point: float32, float64
     - Complex: complex64, complex128
   - Boolean Type: bool
   - String Type: string

2. Composite Types:
   - Array: Fixed-length sequence of elements
   - Slice: Dynamic-length sequence of elements
   - Map: Key-value pairs
   - Struct: User-defined type with named fields

3. Other Types:
   - Pointer: Stores the memory address of a value
   - Function: Functions are first-class citizens
   - Interface: Defines a set of method signatures
   - Channel: Used for communication between goroutines

### 2.4 Operators

Go provides various operators:

1. Arithmetic: `+`, `-`, `*`, `/`, `%`
2. Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
3. Logical: `&&`, `||`, `!`
4. Bitwise: `&`, `|`, `^`, `<<`, `>>`
5. Assignment: `=`, `+=`, `-=`, `*=`, `/=`, etc.
6. Address and Pointer: `&`, `*`

### 2.5 Control Structures

Go supports standard control structures:

1. if-else statements:
   ```go
   if condition {
       // code
   } else if anotherCondition {
       // code
   } else {
       // code
   }
   ```

2. for loops:
   ```go
   for i := 0; i < 10; i++ {
       // code
   }

   // while-like loop
   for condition {
       // code
   }

   // infinite loop
   for {
       // code
   }

   // range-based loop
   for index, value := range collection {
       // code
   }
   ```

3. switch statements:
   ```go
   switch variable {
   case value1:
       // code
   case value2, value3:
       // code
   default:
       // code
   }

   // tagless switch
   switch {
   case condition1:
       // code
   case condition2:
       // code
   default:
       // code
   }
   ```

4. defer statements:
   ```go
   defer function()
   ```

### 2.6 Composite Literals and Initializers

Composite literals provide a concise way to create and initialize composite types:

1. Array Literal:
   ```go
   primes := [5]int{2, 3, 5, 7, 11}
   ```

2. Slice Literal:
   ```go
   fruits := []string{"apple", "banana", "orange"}
   ```

3. Map Literal:
   ```go
   ages := map[string]int{
       "Alice": 30,
       "Bob":   25,
   }
   ```

4. Struct Literal:
   ```go
   type Person struct {
       Name string
       Age  int
   }

   p := Person{Name: "John", Age: 28}
   ```

5. Nested Composite Literals:
   ```go
   type Address struct {
       Street string
       City   string
   }

   type Employee struct {
       Name    string
       Address Address
   }

   emp := Employee{
       Name: "John Doe",
       Address: Address{
           Street: "123 Main St",
           City:   "Anytown",
       },
   }
   ```

### 2.7 Formatting Functions

Go provides a variety of formatting functions in the `fmt` package for printing and string manipulation:

1. Printing Functions:
   - `fmt.Print`: Prints its arguments with their default formats.
   - `fmt.Println`: Like Print, but adds spaces between arguments and a newline at the end.
   - `fmt.Printf`: Formats according to a format specifier and prints.

2. String Formatting Functions:
   - `fmt.Sprint`: Returns a string containing the default formats of its arguments.
   - `fmt.Sprintln`: Like Sprint, but adds spaces between arguments and a newline at the end.
   - `fmt.Sprintf`: Returns a string formatted according to a format specifier.

3. Formatted I/O:
   - `fmt.Fprintf`: Formats and writes to a specified io.Writer.
   - `fmt.Fscanf`: Scans formatted text from a specified io.Reader.

4. Scanning Functions:
   - `fmt.Scan`: Scans text read from standard input, storing successive space-separated values into successive arguments.
   - `fmt.Scanf`: Scans text read from standard input, parsing according to a format string.
   - `fmt.Scanln`: Like Scan, but stops scanning at a newline.

5. Formatting Directives:
   Formatting directives are used with `Printf`, `Sprintf`, and related functions to specify how to format values. Here are some common directives:

   - `%v`: The value in a default format
   - `%+v`: The value in a default format with field names for structs
   - `%#v`: A Go-syntax representation of the value
   - `%T`: A Go-syntax representation of the type of the value
   - `%t`: The word true or false (for boolean values)
   - `%d`: Base 10 integer
   - `%b`: Base 2 integer
   - `%o`: Base 8 integer
   - `%x`, `%X`: Base 16 integer, with lower-case/upper-case letters for a-f
   - `%f`, `%F`: Decimal point, no exponent
   - `%e`, `%E`: Scientific notation
   - `%s`: String
   - `%q`: Double-quoted string
   - `%p`: Pointer address
   - `%c`: The character represented by the corresponding Unicode code point

Example usage of formatting directives:

```go
num := 42
pi := 3.14159
name := "Gopher"

fmt.Printf("Integer: %d\n", num)
fmt.Printf("Float: %.2f\n", pi)
fmt.Printf("String: %s\n", name)
fmt.Printf("Boolean: %t\n", true)
fmt.Printf("Value: %v\n", num)
fmt.Printf("Pointer: %p\n", &num)
fmt.Printf("Type: %T\n", pi)
fmt.Printf("Quoted string: %q\n", name)
fmt.Printf("Hex: %#x\n", num)
```

Output:
```
Integer: 42
Float: 3.14
String: Gopher
Boolean: true
Value: 42
Pointer: 0xc0000b4008
Type: float64
Quoted string: "Gopher"
Hex: 0x2a
```

These formatting directives provide fine-grained control over how values are formatted in output strings. They are essential for creating well-formatted, readable output in Go programs.

## 3. Functions and Methods

### 3.1 Function Basics

Functions in Go are declared using the `func` keyword:

```go
func add(a, b int) int {
    return a + b
}
```

Multiple return values:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

Named return values:

```go
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return
}
```

### 3.2 Methods

Methods are functions associated with a type:

```go
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}
```

Additional examples:
- [Method Examples](Methods/main.go)

### 3.3 The `init` Function

The `init` function is automatically executed before the main function:

```go
func init() {
    // Initialization code
}
```

Characteristics of `init`:
- It doesn't take any arguments.
- It doesn't return any values.
- It's optional to define an `init` function.
- Multiple `init` functions can be defined within a single package.
- `init` functions are executed in the order they appear in the source file.
- They run after all variable declarations in the package have evaluated their initializers.

### 3.4 Variadic Functions

Variadic functions can accept a variable number of arguments:

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Usage
fmt.Println(sum(1, 2, 3))     // Output: 6
fmt.Println(sum(4, 5, 6, 7))  // Output: 22

numbers := []int{1, 2, 3, 4, 5}
fmt.Println(sum(numbers...))  // Output: 15
```

### 3.5 Anonymous Functions and Closures

Anonymous functions:

```go
func() {
    fmt.Println("I'm anonymous!")
}()

add := func(a, b int) int {
    return a + b
}
```

Closures:

```go
func counterFactory(start int) func() int {
    count := start
    return func() int {
        count++
        return count
    }
}

counter1 := counterFactory(0)
counter2 := counterFactory(10)

fmt.Println(counter1()) // Output: 1
fmt.Println(counter1()) // Output: 2
fmt.Println(counter2()) // Output: 11
```

Additional closure example:

```go
func debitCardFunction(balance int) func(int) int {
    return func(withdrawal int) int {
        if balance < withdrawal {
            return balance
        }
        balance -= withdrawal
        return balance
    }
}

debitCard := debitCardFunction(100)
fmt.Println(debitCard(20)) // Output: 80
fmt.Println(debitCard(30)) // Output: 50
```

Reference: [Closure Examples](closures/main.go)

### 3.6 Defer Functions

The `defer` keyword postpones the execution of a function until the surrounding function returns:

```go
func example() {
    defer fmt.Println("This prints last")
    fmt.Println("This prints first")
}
```

Multiple defers:

Deferred functions are executed in Last-In-First-Out (LIFO) order. When multiple defer statements are used, they are pushed onto a stack and executed in reverse order when the surrounding function returns.

```go
func multipleDefers() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    
    fmt.Println("Function body")
}

// Output:
// Function body
// Third defer
// Second defer
// First defer
```

In this example, the deferred functions are executed in reverse order of their declaration after the main function body completes.

## 4. Data Structures

### 4.1 Arrays and Slices

#### 4.1.1 Arrays

Arrays in Go are fixed-size sequences of elements of the same type. Key characteristics:

1. Declaration and Initialization:
```go
// Declaration with size
var arr [5]int

// Declaration with initialization
arr := [5]int{1, 2, 3, 4, 5}

// Size inference
arr := [...]int{1, 2, 3, 4, 5}

// Partial initialization
arr := [5]int{1, 2} // [1 2 0 0 0]
```

2. Memory Layout:
- Arrays are value types (copying an array creates a new copy)
- Contiguous memory allocation
- Fixed size at compile time
- Size is part of the type ([5]int and [6]int are different types)

```go
arr1 := [3]int{1, 2, 3}
arr2 := arr1        // Creates a copy
arr2[0] = 10       // Doesn't affect arr1
fmt.Println(arr1)  // [1 2 3]
fmt.Println(arr2)  // [10 2 3]
```

3. Multi-dimensional Arrays:
```go
// 2D array declaration
var matrix [3][4]int

// 2D array initialization
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

#### 4.1.2 Slices

Slices are dynamic, flexible views into arrays. They consist of:
- Pointer to underlying array
- Length (len)
- Capacity (cap)

1. Creation Methods:
```go
// From array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2 3 4]

// Using make
slice := make([]int, 3, 5)  // length 3, capacity 5

// Literal syntax
slice := []int{1, 2, 3}

// Zero value
var slice []int  // nil slice
```

2. Slice Header Structure:
```go
type SliceHeader struct {
    Data uintptr // Pointer to the underlying array
    Len  int    // Number of elements in the slice
    Cap  int    // Capacity (maximum number of elements)
}
```

3. Capacity Growth:
When appending to a slice beyond its capacity:
- If cap < 1024: new_cap = old_cap * 2
- If cap ≥ 1024: new_cap = old_cap + old_cap/4

```go
slice := make([]int, 0)
fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))

for i := 0; i < 10; i++ {
    slice = append(slice, i)
    fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))
}

// Output shows capacity doubling:
// Len: 0, Cap: 0
// Len: 1, Cap: 1
// Len: 2, Cap: 2
// Len: 3, Cap: 4
// Len: 4, Cap: 4
// Len: 5, Cap: 8
// ...
```

4. Slice Operations and Behavior:

a) Slicing Operation [low:high:max]:

low: starting index (inclusive)
high: ending index (exclusive)
max: maximum capacity index (exclusive)

```go
arr := [6]int{1, 2, 3, 4, 5, 6}

slice1 := arr[1:4]    // [2 3 4], len=3, cap=5
slice2 := arr[1:4:4]  // [2 3 4], len=3, cap=3

// Visualization:
// arr:    [1 2 3 4 5 6]
//            ↑     ↑ ↑
//            low   high max
```

b) Sharing Underlying Array:
```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:3]
slice2 := original[2:4]

slice1[1] = 10  // Affects original and slice2
fmt.Println(original)  // [1 2 10 4 5]
fmt.Println(slice2)   // [10 4]
```

5. Common Pitfalls and Solutions:

a) Unexpected Sharing:
```go
// Potential issue
data := []int{1, 2, 3, 4, 5}
slice := data[2:4]
newData := append(slice, 6)  // Might modify original data

// Solution: Use full slice expression
slice := data[2:4:4]  // Limits capacity
newData := append(slice, 6)  // Creates new array
```

b) Memory Leaks with Large Arrays:
```go
// Potential memory leak
largeArray := [1000000]int{}
slice := largeArray[1:5]  // Holds reference to large array

// Solution: Copy needed elements
slice := make([]int, 4)
copy(slice, largeArray[1:5])
```

6. Performance Considerations:

a) Pre-allocation for Known Size:
```go
// Inefficient
var slice []int
for i := 0; i < 10000; i++ {
    slice = append(slice, i)
}

// Efficient
slice := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    slice = append(slice, i)
}
```

b) Copying vs Referencing:
```go
// Full copy (new backing array)
newSlice := make([]int, len(originalSlice))
copy(newSlice, originalSlice)

// Reference (shares backing array)
reference := originalSlice[:]
```

7. Useful Patterns:

a) Stack Operations:
```go
// Push
stack = append(stack, value)

// Pop
n := len(stack) - 1
value := stack[n]
stack = stack[:n]

// Top
value := stack[len(stack)-1]
```

b) Queue Operations:
```go
// Enqueue
queue = append(queue, value)

// Dequeue (inefficient for large queues)
value := queue[0]
queue = queue[1:]

// Dequeue (efficient)
value := queue[0]
copy(queue, queue[1:])
queue = queue[:len(queue)-1]
```

8. Common Functions and Operations:

```go
// Length and Capacity
len(slice)  // Number of elements
cap(slice)  // Maximum capacity

// Append
slice = append(slice, 1, 2, 3)
slice = append(slice, otherSlice...)

// Copy
copy(dst, src)

// Clear
slice = slice[:0]  // Maintains capacity
slice = nil        // Releases memory

// Remove element
slice = append(slice[:i], slice[i+1:]...)

// Filter
filtered := slice[:0]
for _, x := range slice {
    if keepElement(x) {
        filtered = append(filtered, x)
    }
}
```

These concepts and examples cover the main aspects of arrays and slices in Go, including memory management, capacity growth, common operations, and best practices for various scenarios.

Additional examples:

- [Basic Array and Slice Operations](arrayAndSlices/main.go)
- [2D Slices](2dslices/main.go)

2D slices example:

```go
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols)
}
```

### 4.2 Maps

Maps are key-value data structures in Go that allow efficient lookup, insertion, and deletion operations. They are implemented as hash tables and provide an unordered collection of key-value pairs. Here's a brief overview of maps:

1. Declaration: `map[KeyType]ValueType`
2. Initialization: Can be done using `make()` or map literals
3. Operations: Adding, updating, deleting, and retrieving values
4. Concurrency: Not safe for concurrent use without additional synchronization
5. Performance: Average time complexity of O(1) for basic operations

Key features:
- Keys must be comparable (e.g., numbers, strings, structs with comparable fields)
- Values can be of any type
- Maps automatically grow as needed
- The zero value of a map is nil

### Maps as Pointers

In Go, maps are reference types, meaning that when you pass a map to a function, you are passing a reference to the original map, not a copy. This allows modifications made to the map within the function to affect the original map.

Example of modifying a map in a function:

```go
func updateMap(m map[string]int) {
    m["banana"] = 5 // Modifies the original map
}

func main() {
    m := map[string]int{
        "apple": 1,
        "banana": 2,
    }
    updateMap(m)
    fmt.Println(m["banana"]) // Output: 5
}
```

### How Maps Grow

Maps in Go automatically resize when the number of elements exceeds a certain threshold, which is determined by the load factor. When a map grows, Go allocates a new, larger underlying array and rehashes the existing key-value pairs into the new array. This process is handled automatically, and the programmer does not need to manage the resizing manually.

Example of adding elements to a map:

```go
m := make(map[string]int)
for i := 0; i < 100; i++ {
    m[fmt.Sprintf("key%d", i)] = i // The map grows as needed
}
```

### Example Usage

```go
m := map[string]int{
    "apple": 1,
    "banana": 2,
}

// Add or update
m["cherry"] = 3

// Delete
delete(m, "banana")

// Check existence
if value, exists := m["apple"]; exists {
    fmt.Println(value) // Output: 1
}
```

Additional examples:
- [Map Operations](Maps/main.go)

This example demonstrates various map operations, including declaration, initialization, adding elements, deleting elements, and iteration.

## 5. Pointers

### 5.1 Pointer Basics

Pointers store the memory address of a value:

```go
x := 10
ptr := &x
fmt.Println(*ptr) // Output: 10
*ptr = 20
fmt.Println(x) // Output: 20
```

### 5.2 Pointers to Structs

```go
type Person struct {
    Name string
    Age  int
}

func modifyPerson(p *Person) {
    p.Name = "Alice"
    p.Age = 30
}

person := Person{"Bob", 25}
modifyPerson(&person)
fmt.Println(person) // Output: {Alice 30}
```

Go provides automatic dereferencing of pointers to structs, but only in specific contexts. This feature is called "implicit dereferencing" and applies exclusively to method receivers. Here's a more detailed explanation:

1. Method Receivers:
   When a method is defined with a pointer receiver, Go allows you to call that method on either a pointer or a value of the struct type. The compiler automatically handles the dereferencing.

   ```go
   type Person struct {
       Name string
   }

   func (p *Person) SetName(name string) {
       p.Name = name  // p is automatically dereferenced
   }

   person := Person{}
   person.SetName("Alice")  // Go automatically takes the address of person
   ```

2. Function Parameters:
   This implicit dereferencing does not apply to regular function parameters. If a function expects a pointer, you must explicitly pass a pointer.

   ```go
   func UpdateName(p *Person, name string) {
       p.Name = name
   }

   person := Person{}
   UpdateName(&person, "Bob")  // Must explicitly take the address of person
   ```

3. Compilation Error:
   If you try to pass a value to a function expecting a pointer (or vice versa) without explicit conversion, you'll get a compilation error:

   ```go
   func UpdateName(p *Person, name string) {
       p.Name = name
   }

   person := Person{}
   UpdateName(person, "Charlie")  // Compilation error: cannot use person (type Person) as type *Person
   ```

This distinction is important to understand because it affects how you write and call methods versus functions in Go, and it's a key part of Go's approach to balancing convenience and explicitness in pointer usage.

Additional examples:
- [Pointer Examples](Pointer/main.go)
## 6. Structs

Structs in Go are composite types that group together variables under a single name.

### 6.1 Basic Struct Definition and Usage

```go
type Person struct {
    Name string
    Age  int
}

// Creating and initializing a struct
// Note: Each line in struct initialization must end with a comma, 
// including the last line. This makes version control diffs cleaner 
// and makes it easier to add/remove fields.
p1 := Person{
    Name: "Alice",  // comma required
    Age: 30,        // comma required even on last line
}

// This is harder to maintain:
p2 := Person{
    Name: "Bob",    // comma required
    Age: 25  // no comma makes it harder to add new fields
}

// Multi-field struct with proper formatting:
type Employee struct {
    Person: Person{
        Name: "John",     // comma required
        Age: 30,         // comma required
    },                   // comma required
    Address: Address{
        Street: "123 Main St",  // comma required
        City: "Anytown",        // comma required
    },                         // comma required
    Salary: 50000,            // comma required
}                            // no comma needed for closing brace
```

### 6.2 Struct Methods

Methods can be defined on structs:

```go
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

fmt.Println(p1.Greet())
```

### 6.3 Pointer Receivers

Methods can use pointer receivers to modify the struct:

```go
func (p *Person) Birthday() {
    p.Age++
}

p1.Birthday()
fmt.Println(p1.Age) // Increased by 1
```

### 6.4 Embedding and Composition

Go supports struct embedding for composition:

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Person  // Embedding Person
    Address // Embedding Address
    Salary  float64
}

emp := Employee{
    Person:  Person{Name: "John", Age: 30},
    Address: Address{Street: "123 Main St", City: "Anytown"},
    Salary:  50000,
}

fmt.Println(emp.Name) // Accessing embedded Person field
fmt.Println(emp.Street) // Accessing embedded Address field
```

### 6.5 Anonymous Fields

You can embed structs or other types without specifying a field name:

```go
type Manager struct {
    Employee
    Department string
}

mgr := Manager{
    Employee: Employee{
        Person:  Person{Name: "Jane", Age: 35},
        Address: Address{Street: "456 Elm St", City: "Othertown"},
        Salary:  75000,
    },
    Department: "IT",
}
```

### 6.6 Struct Tags

Struct tags provide metadata about struct fields:

```go
type User struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"-" validate:"required,min=8"`
}

// Using reflection to access tags
t := reflect.TypeOf(User{})
field, _ := t.FieldByName("Email")
fmt.Println(field.Tag.Get("json"))    // Output: email
fmt.Println(field.Tag.Get("validate")) // Output: required,email
```

### 6.7 Anonymous Structs

You can create one-off structs without defining a new type:

```go
point := struct {
    X, Y int
}{10, 20}
```

### 6.8 Comparing Structs

Structs are comparable if all their fields are comparable:

```go
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}
fmt.Println(p1 == p2) // Output: true
```

### 6.9 Struct Initialization with New

You can use the `new` function to create a pointer to a zeroed struct:

```go
p := new(Person)
p.Name = "David"
p.Age = 40
```

Additional examples:
- [Struct Examples](Structs/main.go)

This file demonstrates various aspects of struct usage in Go, including definition, methods, embedding, and more complex struct compositions.

## 7. Interfaces

### 7.1 Interface Basics

Interfaces in Go are a fundamental concept that define a set of method signatures. They provide a powerful way to achieve abstraction and polymorphism in Go programs. Unlike some other languages, Go interfaces are implemented implicitly, which allows for a high degree of flexibility and decoupling between packages.

Definition and Implementation:
An interface type is defined as a set of method signatures. For example:

```go
type Writer interface {
    Write([]byte) (int, error)
}
```

A type implicitly implements an interface if it defines all the methods specified by that interface. There's no need for explicit declaration of intent to implement an interface:

```go
type FileWriter struct {
    // ...
}
```
In this example, `FileWriter` implicitly implements the `Writer` interface because it has a `Write` method with the correct signature.

Interface Values:
An interface value consists of two components: a concrete value and a dynamic type. This allows for runtime polymorphism:

```go
var w Writer
w = FileWriter{}  // w now holds a FileWriter value
```

Empty Interface:
The empty interface `interface{}` is satisfied by all types, as it has no methods. It's often used to handle values of unknown type:

```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}
```

Type Assertions and Type Switches:
Go provides mechanisms to work with the concrete types behind interfaces:

Type Assertions allow you to extract the underlying value from an interface:
```go
fw, ok := w.(FileWriter)
if ok {
    // w holds a FileWriter
}
```

Type Switches determine the type of an interface value:
```go
switch v := w.(type) {
case FileWriter:
    fmt.Println("FileWriter:", v)
case *BufferedWriter:
    fmt.Println("BufferedWriter:", v)
default:
    fmt.Println("Unknown type")
}
```

Interface Composition:
Interfaces can be composed of other interfaces, allowing for more complex abstractions:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

Best Practices:
1. Keep interfaces small and focused on a single responsibility.
2. Use interfaces to define behavior, not data.
3. Accept interfaces, return concrete types in function signatures.
4. Use the `io.Reader` and `io.Writer` interfaces when dealing with streams of data.

Additional examples:
- [Interface Examples](Interfaces/main.go)
- [Shape Interface](shapeInterface/main.go)

Interfaces play a crucial role in Go's design philosophy, enabling loose coupling between components and facilitating easier testing and maintenance of code. They are extensively used in the standard library and are a key feature for writing flexible and reusable Go code.

### 7.2 Stringer Interface

The Stringer interface is used for custom string representations of types. It's particularly useful for printing custom types in a human-readable format.

```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

p := Person{Name: "Alice", Age: 30}
fmt.Println(p) // Output: Alice (30 years old)
```

- [Stringer Interface](stringerinterface/main.go)

By implementing the String() method, we provide a custom string representation for the Person struct. This is automatically used when the struct is printed or converted to a string.

### 7.3 Sorting in Go

Go provides multiple ways to sort collections, including `sort.Sort`, `sort.Slice`, and `slices.Sort`. Each method has its own use cases and advantages.

#### 7.3.1 sort.Interface

The `sort.Interface` is used for custom sorting of collections. It requires implementing three methods: `Len()`, `Less()`, and `Swap()`.

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
}

sort.Sort(ByAge(people))
fmt.Println(people)
```

This example demonstrates how to implement custom sorting for a slice of `Person` structs based on their age using `sort.Sort`.

#### 7.3.2 sort.Slice

`sort.Slice` is a more convenient way to sort slices, introduced in Go 1.8. It doesn't require implementing the `sort.Interface`:

```go
people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})

fmt.Println(people)
```

#### 7.3.3 slices.Sort (Go 1.21+)

In Go 1.21, the `slices` package was introduced, providing a type-safe and more efficient sorting method:

```go
import "slices"

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
}

slices.SortFunc(people, func(a, b Person) int {
    return a.Age - b.Age
})

fmt.Println(people)
```

#### 7.3.4 Differences and Use Cases

1. `sort.Sort`:
   - Requires implementing `sort.Interface`
   - Useful for custom types with complex sorting logic
   - More verbose but offers full control over sorting behavior

2. `sort.Slice`:
   - More convenient for one-off sorting operations
   - Doesn't require implementing `sort.Interface`
   - Less type-safe (uses `interface{}` internally)
   - Slightly less performant than `sort.Sort`

3. `slices.Sort` / `slices.SortFunc`:
   - Type-safe and generally more efficient
   - Available only in Go 1.21+
   - Preferred method for simple sorting operations in newer Go versions

Choose the method that best fits your use case:
- Use `sort.Sort` for custom types with complex sorting logic or when you need to reuse the sorting logic.
- Use `sort.Slice` for quick, one-off sorting operations in older Go versions.
- Use `slices.Sort` or `slices.SortFunc` for efficient, type-safe sorting in Go 1.21+.

This overview covers the main sorting methods in Go, their differences, and appropriate use cases. Is there any specific aspect of sorting in Go you'd like me to elaborate on?

Additional examples:
- [Sort Interface](sortinterface/main.go)

## 8. Concurrency

### 8.1 Goroutines

Goroutines are lightweight threads managed by the Go runtime. They are more efficient than OS threads and allow for concurrent execution of functions. Here's a brief comparison:

Goroutines vs OS threads vs threads in other languages:
1. Goroutines: Managed by Go runtime, lightweight, can run many concurrently
2. OS threads: Managed by the operating system, heavier, limited by system resources
3. Threads in other languages: Often map directly to OS threads, heavier than goroutines

Goroutines are lightweight because:
1. They have a smaller stack size (starting at 2KB)
2. They are multiplexed onto a small number of OS threads:
   This means that many goroutines can run on a single OS thread. The Go runtime scheduler manages this multiplexing, allowing it to run thousands or even millions of goroutines on just a few OS threads. This is more efficient than creating a new OS thread for each concurrent task, as OS threads are more resource-intensive.

3. Context switching between goroutines is faster than OS threads:
   Context switching is the process of storing the state of a thread or goroutine so that it can be resumed later, and then switching to run another thread or goroutine. For goroutines, this process is managed by the Go runtime and is much faster than OS-level context switches for several reasons:
   - Goroutine state is smaller and simpler than full thread state, so it's quicker to save and restore.
   - The Go scheduler doesn't need to switch to kernel mode to perform a context switch, which is a time-consuming operation for OS threads.
   - The Go scheduler can make more intelligent decisions about when to switch contexts, as it has more information about the goroutines it's managing than the OS has about threads.
4. Preemptive scheduling:
   The Go scheduler implements preemptive scheduling for goroutines, which enables:
   - Interruption of running goroutines at specific points to allow other goroutines to run
   - Since Go 1.14, asynchronous preemption that can interrupt even CPU-intensive tasks
   - Preemption at function calls and backward branches in loops to prevent CPU monopolization
   - Runtime checks that determine if a goroutine has been running too long and should yield
   - Improved application responsiveness and fair resource allocation among goroutines
   - Better system utilization compared to the pre-1.14 cooperative scheduling model

These factors contribute to making goroutines much more lightweight and efficient than OS threads, allowing Go programs to handle high levels of concurrency with relatively low overhead.

```go
func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

go printNumbers()
time.Sleep(time.Second)
```

### 8.2 Channels

Channels are used for communication between goroutines. There are two types of channels in Go:

1. Unbuffered Channels:
   - Created with `make(chan T)`
   - Synchronous: sender blocks until receiver is ready
   - Used when you need guaranteed delivery and synchronization

2. Buffered Channels:
   - Created with `make(chan T, capacity)`
   - Asynchronous: sender only blocks when buffer is full
   - Used when you want to decouple sender and receiver, or manage flow control

Differences between buffered and unbuffered channels:
- Unbuffered channels provide immediate handoff and synchronization
- Buffered channels allow for some decoupling and can improve performance in certain scenarios

Usage:
- Use unbuffered channels when you need strict synchronization between goroutines
- Use buffered channels when you want to allow some slack between sender and receiver, or to implement a semaphore-like behavior

Example:

```go
ch := make(chan int)
go func() {
    ch <- 42
}()
value := <-ch
fmt.Println(value)

// Buffered channel
bufCh := make(chan int, 2)
bufCh <- 1
bufCh <- 2
fmt.Println(<-bufCh, <-bufCh)

// Closing channels
close(ch)
```

Additional examples:
- [Channel Operations and Patterns](channels/main.go)
- [Concurrency Patterns](concurrencypatterns/)
- [Odd even](oddEven/main.go)
- [Odd even using channels and wg](oddevenusingchannelsandwg/main.go)
- [Odd even using just channels](oddevenusingjustchannels/main.go)
- [Odd even using wait group](oddevenusingwaitgroup/main.go)

This file demonstrates various channel operations, including creating channels, sending and receiving values, and using channels for communication between goroutines. It also shows how to use channels to implement a simple worker pool pattern.

### 8.3 Select Statement

The select statement is used to work with multiple channels:

```go
select {
case msg1, ok := <-ch1:
    if !ok {
        fmt.Println("ch1 is closed")
    } else {
        fmt.Println("Received from ch1:", msg1)
    }
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel operations ready")
}
```

Key points about select with closed channels:

1. When receiving from a channel, you can use the two-value form of channel receive to check if the channel is closed.
2. For a closed channel, the receive operation `<-ch` returns immediately with the zero value of the channel's type and `ok` set to `false`.
3. The select statement will choose a closed channel if no other cases are ready, allowing you to detect and handle closed channels.
4. Sending on a closed channel will cause a panic, so it's important to ensure a channel is open before sending.

Example with a closed channel:

```go
ch := make(chan int)
close(ch)

select {
case val, ok := <-ch:
    if !ok {
        fmt.Println("Channel is closed")
    } else {
        fmt.Println("Received:", val)
    }
default:
    fmt.Println("No value received")
}
// Output: Channel is closed
```

This behavior allows for graceful handling of channel closure in concurrent programs.

### 8.4 Synchronization Primitives

Go provides several synchronization primitives:

1. sync.Mutex:
   ```go
   var mu sync.Mutex
   mu.Lock()
   // Critical section
   mu.Unlock()
   ```

2. sync.RWMutex:
   ```go
   var rwmu sync.RWMutex
   rwmu.RLock()
   // Read operations
   rwmu.RUnlock()

   rwmu.Lock()
   // Write operations
   rwmu.Unlock()
   ```

3. sync.WaitGroup:
   ```go
   var wg sync.WaitGroup
   wg.Add(1)
   go func() {
       defer wg.Done()
       // Do work
   }()
   wg.Wait()
   ```

4. sync.Once:
   ```go
   var once sync.Once
   once.Do(func() {
       // This will only be executed once
   })
   ```

5. sync.Cond:
   ```go
   var mu sync.Mutex
   cond := sync.NewCond(&mu)

   // Wait for condition
   cond.L.Lock()
   for !condition {
       cond.Wait()
   }
   cond.L.Unlock()

   // Signal condition change
   cond.Signal()
   // or
   cond.Broadcast()
   ```

### 8.5 Concurrency Patterns

Here's a brief overview of some common concurrency patterns in Go:

1. Worker Pool Pattern:
```go
func workerPool(numWorkers int, jobs <-chan int, results chan<- int) {
    for i := 0; i < numWorkers; i++ {
        go func() {
            for job := range jobs {
                results <- processJob(job)
            }
        }()
    }
}
```

2. Generator Pattern:
```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}
```

3. Fan-Out, Fan-In Pattern:
Distributes work across multiple goroutines (fan-out) and then combines the results (fan-in).

```go
func fanOut(input <-chan int, numWorkers int) []<-chan int {
    outputs := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        outputs[i] = worker(input)
    }
    return outputs
}

func fanIn(channels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	multiplexed := make(chan Result)

	multiplex := func(c <-chan Result) {
		defer wg.Done()
		for result := range c {
			multiplexed <- result
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}
```

4. Pipeline Pattern:
A series of stages connected by channels, where each stage is a group of goroutines running the same function.

```go
func stage1(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * 2
        }
    }()
    return out
}

func stage2(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range n {
            out <- n + 1
        }
    }()
    return out
}

// Usage:
// input := generator(1, 2, 3)
// result := stage2(stage1(input))
```

5. Context Pattern:
   Uses the `context` package for managing cancellation, deadlines, and passing request-scoped values across API boundaries and between processes.
   Key features:
   - Cancellation: Allows canceling long-running operations
   - Deadlines: Sets time limits for operations
   - Values: Carries request-scoped data
   
   Basic usage:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   
   select {
   case <-ctx.Done():
       fmt.Println("Operation cancelled or timed out")
       return ctx.Err()
   case result := <-doOperation(ctx):
       return result
   }
   ```
   The context example demonstrates the use of the `context` package for managing goroutine lifecycles and cancellation. It shows how to create a context with a timeout and use it to control multiple goroutines.

6. Confinement:
   Restricts data access to a single goroutine, simplifying concurrent programming and avoiding race conditions. Two types:
   1. Lexical confinement: Data confined by scope (e.g., local variables).
   2. Ad-hoc confinement: Data confined by convention (e.g., only one goroutine accesses shared data).

   Example: [Confinement](concurrencypatterns/confinement/main.go)

   This example demonstrates confinement by:
   1. Each goroutine works on its own slice element (`&result[i]`).
   2. No shared mutable state between goroutines.
   3. `result` slice is pre-allocated, avoiding dynamic resizing.

   Mutex locks are avoided because:
   - Each goroutine writes to a distinct memory location.
   - No concurrent access to shared data structures.

   Without confinement, if using a shared slice, a mutex would be required.

7. Mutex for shared state:
   Uses mutual exclusion to protect shared data structures from concurrent access. It can be avoided by using confinement in above example.
   Examples: [Mutex Map](concurrencypatterns/mutex-map/), [Mutex Example](concurrencypatterns/mutex-example/)

8. Or-Done pattern:
   Allows for cancellation of multiple channels simultaneously. A separate re-usable function is created for this pattern. It is useful when you have multiple channels and want to cancel them all when one of them is done.
   Example: [Or Done](concurrencypatterns/or-done/)

Channel Management Best Practices:

1. When to close channels:
   - Required: For range loops and signaling completion to multiple goroutines
   - Not required: One-time communication, garbage-collected channels, synchronization-only channels

2. Common mistakes to avoid:
```go
// Never write to a closed channel
ch := make(chan int)
close(ch)
ch <- 1 // panic: send on closed channel

// Never close a channel twice
ch := make(chan int)
close(ch)
close(ch) // panic: close of closed channel
```

Examples:
- [Context Example](concurrencypatterns/context-example/main.go)
- [Generator](concurrencypatterns/generator/)
- [Confinement](concurrencypatterns/confinement/)
- [Prime Fan-in Fan-out](concurrencypatterns/prime-fan-in-fan-out/)
- [Mutex Map](concurrencypatterns/mutex-map/)
- [Mutex Example](concurrencypatterns/mutex-example/)
- [Prime Pipeline](concurrencypatterns/prime-pipeline/)
- [Or Done](concurrencypatterns/or-done/)

These patterns demonstrate various techniques for managing concurrency, from distributing work and combining results to protecting shared resources and handling cancellation.



## 9. Error Handling

### 9.1 The error Interface

In Go, errors are represented by the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

Any type that implements this interface can be used as an error. The most common way to create errors is using the `errors.New()` function:

```go
import "errors"

err := errors.New("something went wrong")
```

Example of using errors:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

### 9.2 Custom Errors

You can create custom error types by implementing the `error` interface:

```go
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func someFunction() error {
    return &MyError{Code: 404, Message: "Not Found"}
}

err := someFunction()
if err != nil {
    fmt.Println(err) // Output: Error 404: Not Found
}
```

### 9.3 Panic and Recover

Panic is used for unrecoverable errors. It stops the normal execution of the current goroutine and begins panicking:

```go
func doSomething() {
    panic("something went terribly wrong")
}

func main() {
    doSomething()
    fmt.Println("This line will not be executed")
}
```

`recover` is used to regain control of a panicking goroutine. It's only useful inside deferred functions:

```go
func mayPanic() {
    panic("a problem")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()
    mayPanic()
    fmt.Println("Never reached")
}
```

### 9.4 Defer

`defer` statements are often used with panics. Deferred functions are still executed in panic situations:

```go
func riskyFunction() {
    defer func() {
        fmt.Println("This will always execute")
    }()

    panic("Oops!")
}

func main() {
    riskyFunction()
}
```

### 9.5 fmt.Errorf and Error Wrapping

`fmt.Errorf` is used to create formatted error messages:

```go
name := "John"
age := 30
err := fmt.Errorf("invalid user: %s (age %d)", name, age)
fmt.Println(err)
// Output: invalid user: John (age 30)
```

It can also wrap errors using the `%w` verb (Go 1.13+):

```go
originalErr := errors.New("database connection failed")
wrappedErr := fmt.Errorf("failed to fetch user data: %w", originalErr)
```

### 9.6 Sentinel Errors

These are predefined errors used for specific conditions:

```go
import "io"

var ErrEOF = errors.New("EOF")

func readFile() error {
    // ...
    return io.EOF
}

if err := readFile(); err == io.EOF {
    fmt.Println("End of file reached")
}
```

### 9.7 Error Type Assertions and Type Switches

Useful for handling different types of errors:

```go
type NetworkError struct {
    Code int
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("Network error with code: %d", e.Code)
}

func handleError(err error) {
    switch e := err.(type) {
    case *NetworkError:
        fmt.Printf("Network error with code %d\n", e.Code) // Same output as fmt.Println(e)
    default:
        fmt.Println("Unknown error:", err)
    }
}
```

### 9.8 errors.Is and errors.As

Go 1.13 introduced `errors.Is` and `errors.As` to improve error handling, especially when dealing with wrapped errors. These functions make it easier to check for specific error types or values in an error chain.

#### errors.Is

`errors.Is` checks if a specific error value exists anywhere in the error chain.

Syntax:
```go
func Is(err, target error) bool
```

Use `errors.Is` when you want to compare an error to a sentinel error value.

Example:

```go
var ErrNotFound = errors.New("not found")

func fetchItem(id string) (Item, error) {
    // ... implementation ...
    return Item{}, fmt.Errorf("failed to fetch item: %w", ErrNotFound)
}

func main() {
    _, err := fetchItem("123")
    if errors.Is(err, ErrNotFound) {
        fmt.Println("The item was not found")
    } else {
        fmt.Println("An unknown error occurred:", err)
    }
}
```

In this example, `errors.Is` checks if `ErrNotFound` is anywhere in the error chain, even if it's wrapped.

#### errors.As

`errors.As` finds the first error in the error chain that matches the target type, and if so, sets the target to that error value.

Syntax:
```go
func As(err error, target interface{}) bool
```

Use `errors.As` when you need to check for a specific error type and access its fields or methods.

Example:

```go
type NetworkError struct {
    Code    int
    Message string
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("network error: %s (code: %d)", e.Message, e.Code)
}

func fetchData() error {
    // Simulating a network error
    return fmt.Errorf("failed to fetch data: %w", &NetworkError{Code: 500, Message: "Internal Server Error"})
}

func main() {
    err := fetchData()

    var netErr *NetworkError
    if errors.As(err, &netErr) {
        fmt.Printf("Network error occurred. Code: %d, Message: %s\n", netErr.Code, netErr.Message)
    } else {
        fmt.Println("An unknown error occurred:", err)
    }
}
```

In this example, `errors.As` checks if there's a `NetworkError` in the error chain. If found, it sets `netErr` to point to that error, allowing access to its fields.

#### Benefits over Type Assertions

1. Works with wrapped errors: Both functions work through entire error chains, not just the topmost error.
2. Nil-safety: Unlike type assertions, these functions handle nil errors gracefully.
3. Interface satisfaction: `errors.As` can find errors that implement an interface, not just concrete types.

Example demonstrating these benefits:

```go
type CustomError interface {
    CustomError() string
}

type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}

func (e *MyError) CustomError() string {
    return "This is a custom error: " + e.Msg
}

func someOperation() error {
    return fmt.Errorf("wrapped error: %w", &MyError{Msg: "something went wrong"})
}

func main() {
    err := someOperation()

    // Using errors.As with an interface
    var customErr CustomError
    if errors.As(err, &customErr) {
        fmt.Println(customErr.CustomError())
    }

    // Safe with nil errors
    var nilErr error
    fmt.Println(errors.Is(nilErr, io.EOF)) // false, no panic

    // Works through wrapped errors
    fmt.Println(errors.Is(err, &MyError{})) // true
}
```

This expanded explanation and examples demonstrate how `errors.Is` and `errors.As` provide powerful and flexible error handling capabilities in Go, especially when dealing with error wrapping and custom error types.

### 9.9 Multiple Return Values for Errors

Go's idiomatic way of returning errors along with results:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

### 9.10 Custom Error Types with Additional Methods

Errors can have methods beyond `Error()` for additional functionality:

```go
type ValidationError struct {
    Field string
    Err   error
}

func (v *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field %s: %v", v.Field, v.Err)
}

func (v *ValidationError) Unwrap() error {
    return v.Err
}
```

### 9.11 Implementing Errors as Constants

For errors that don't need to carry additional information:

```go
const ErrInvalidInput = Error("invalid input provided")

type Error string

func (e Error) Error() string {
    return string(e)
}
```

### 9.12 Context Package for Error Handling

The `context` package can be used to carry deadlines, cancellation signals, and other request-scoped values:

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

err := doSomethingSlowWithContext(ctx)
if err == context.DeadlineExceeded {
    fmt.Println("Operation timed out")
}
```

### 9.13 Pointer Usage in Error Handling

Pointers are often used in error handling, especially with `errors.As`:

```go
var netErr *NetworkError
if errors.As(err, &netErr) {
    fmt.Printf("Network error code: %d\n", netErr.Code)
}
```

fmt.Println(point.X, point.Y)

This allows for type safety, modification of the caller's variable, and flexibility in handling both pointer and non-pointer error types.

### 9.14 Unwrapping Errors

Unwrapping errors is the process of accessing the underlying error in a wrapped error chain:

```go
type WrapperError struct {
    Msg string
    Err error
}

func (w *WrapperError) Error() string {
    return fmt.Sprintf("%s: %v", w.Msg, w.Err)
}

func (w *WrapperError) Unwrap() error {
    return w.Err
}

// Using errors.Unwrap
originalErr := errors.New("original error")
wrappedErr := fmt.Errorf("wrapped: %w", originalErr)

unwrappedErr := errors.Unwrap(wrappedErr)
fmt.Println(unwrappedErr == originalErr) // true
```

### 9.15 When is the Error Method Called in fmt Print Statements

The `Error()` method of an error is called implicitly in several situations:

```go
err := errors.New("something went wrong")
fmt.Println(err)  // Calls err.Error() implicitly
fmt.Printf("Error occurred: %v\n", err)  // Calls err.Error()
fmt.Printf("Error message: %s\n", err)   // Calls err.Error()
fmt.Printf("Quoted error: %q\n", err)  // Calls err.Error() and quotes the result
```

For custom types implementing the `error` interface, the same rules apply:

```go
type MyError struct {
    Code int
    Message string
}

func (e MyError) Error() string {
    return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

err := MyError{Code: 404, Message: "Not Found"}
fmt.Println(err)  // Calls err.Error() implicitly
```

Note that the `Error()` method is not called when using reflection-based formatting verbs like `%#v`, and if a nil error is printed, it will output `<nil>` rather than calling any method.

This comprehensive overview covers the essential aspects of error handling in Go, including creating and using errors, custom error types, panic and recover, error wrapping and unwrapping, and best practices for error handling in various scenarios.

Example references:
- [Error Handling](errorhandling/main.go)

## 10. Advanced Topics

### 10.1 Reflection

Reflection allows inspection of types at runtime:

```go
t := reflect.TypeOf(someValue)
v := reflect.ValueOf(someValue)

if v.Kind() == reflect.Struct {
    for i := 0; i < v.NumField(); i++ {
        fmt.Printf("Field: %s\n", v.Type().Field(i).Name)
    }
}
```

Additional examples:
- [Reflection](reflection/main.go)

### 10.2 Generics

Go 1.18+ supports generics:

```go
// Basic generic function
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

// Usage
PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"a", "b", "c"})

// Generic function with multiple type parameters
func Map[T, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

// Generic struct
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Using comparable constraint
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Custom type constraint with ~
type Integer interface {
    ~int | ~int32 | ~int64
}

type Float interface {
    ~float32 | ~float64
}

// Combined numeric constraint
type Number interface {
    Integer | Float
}

// Works with custom integer types
type MyInt int
type MyFloat64 float64

func Sum[T Number](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// Example usage with custom types
func Example() {
    // Works with built-in types
    ints := []int{1, 2, 3}
    fmt.Println(Sum(ints)) // 6

    // Works with custom types too
    myInts := []MyInt{MyInt(1), MyInt(2), MyInt(3)}
    fmt.Println(Sum(myInts)) // 6

    myFloats := []MyFloat64{MyFloat64(1.1), MyFloat64(2.2)}
    fmt.Println(Sum(myFloats)) // 3.3
}

// Generic map type
type Cache[K comparable, V any] struct {
    data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        data: make(map[K]V),
    }
}
```

Key points about generics in Go:
- Use square brackets `[T any]` for type parameters
- `any` is an alias for `interface{}`
- `comparable` is a built-in constraint that allows `==` and `!=` operations
- Custom type constraints can be defined using interface types
- Multiple type parameters are supported `[K, V any]`
- Generic types can be used in structs, interfaces, and methods
- The `~` operator in constraints means "any type with underlying type". For example:
  - `~int` matches both `int` and any type defined as `type MyInt int`
  - Without `~`, the constraint would only match the exact type

Additional examples:
- [Generics](Generics/main.go)

### 10.3 Context Package

The context package is used for cancellation, deadlines, and request-scoped values:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case <-time.After(1 * time.Second):
    fmt.Println("overslept")
case <-ctx.Done():
    fmt.Println(ctx.Err()) // prints "context deadline exceeded"
}
```

func (fw FileWriter) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

## 11. Memory Management

### 11.1 new vs make

`new` and `make` are built-in functions for memory allocation, but they serve different purposes:

- `new` is used to allocate memory for a type and returns a pointer to a zero-initialized value of that type.
- `make` is specifically used to create slices, maps, and channels, and returns an initialized (not zeroed) value of the specified type.

```go
// new returns a pointer to a zero-initialized value
p := new(int)
fmt.Println(*p) // Output: 0

// make is used to create slices, maps, and channels
s := make([]int, 5, 10)  // slice with length 5 and capacity 10
m := make(map[string]int)  // empty map
ch := make(chan int, 5)  // buffered channel with capacity 5
```

Understanding the difference between `new` and `make` is crucial for proper memory allocation and initialization in Go programs.

### 11.2 Memory Management and Garbage Collection

Go uses a sophisticated memory management system with automatic garbage collection. Understanding how memory is managed and collected is crucial for writing efficient Go programs.

#### 11.2.1 Stack vs Heap Allocation

##### Stack Memory
- Fast allocation and deallocation
- Memory is automatically freed when function returns
- Size must be known at compile time
- Limited in size (typically a few MB)
- Thread-local (each goroutine has its own stack)

Examples of stack allocation:
```go
func stackExample() {
    // These will typically be allocated on the stack
    x := 42                    // Basic types
    y := [3]int{1, 2, 3}      // Small arrays
    z := struct{ name string }{"John"} // Small structs
}
```

##### Heap Memory
- Managed by garbage collector
- Flexible size
- Slower than stack allocation
- Shared across goroutines
- Used for values that outlive the function that creates them

Examples of heap allocation:
```go
func heapExample() *int {
    // These will typically be allocated on the heap
    x := new(int)       // Pointer types
    y := make([]int, 3) // Slices
    z := make(map[string]int) // Maps
    return x            // Returning address forces heap allocation
}
```

#### 11.2.2 Escape Analysis

Go's compiler performs escape analysis to determine whether a value can be allocated on the stack or must be allocated on the heap.

Common scenarios that cause heap allocation:
```go
// 1. Returning addresses of local variables
func createPointer() *int {
    x := 42
    return &x  // x escapes to heap
}

// 2. Storing references in interfaces
func createInterface() interface{} {
    x := 42
    return x  // x escapes to heap due to interface conversion
}

// 3. Slices or maps that might grow
func createSlice() []int {
    return make([]int, 0, 10)  // Usually heap allocated
}

// 4. Goroutine-accessed variables
func accessFromGoroutine() {
    x := 42
    go func() {
        fmt.Println(x)  // x escapes to heap
    }()
}
```

You can see escape analysis details using:
```bash
go build -gcflags="-m"
```

#### 11.2.3 Garbage Collection Process

Go uses a concurrent, tri-color mark-and-sweep collector:

1. **Mark Phase**
   - White: Unvisited objects
   - Grey: Visited but not scanned
   - Black: Visited and scanned
   ```go
   // Objects are initially white
   root := &Object{}      // Grey (visited)
   root.Next = &Object{}  // White (unvisited)
   // After scanning root, it becomes black
   // After scanning root.Next, both are black
   ```

2. **Write Barrier**
   - Ensures consistency during concurrent marking
   - Activated during garbage collection
   ```go
   // Write barrier example (conceptual)
   if writeBarrier.enabled {
       writeBarrier(ptr, val)
   } else {
       *ptr = val
   }
   ```

3. **Sweep Phase**
   - Reclaims memory from white objects
   - Runs concurrently with application

#### 11.2.4 GC Tuning

Control garbage collection behavior:

```go
// Force immediate garbage collection
runtime.GC()

// Set GOGC environment variable
// GOGC=50 means GC triggers when heap grows by 50%
os.Setenv("GOGC", "50")

// Get current memory statistics
var stats runtime.MemStats
runtime.ReadMemStats(&stats)
fmt.Printf("Heap size: %d\n", stats.HeapAlloc)
```

#### 11.2.5 Memory Optimization Techniques

1. **Object Pooling**
```go
var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func processData() {
    buf := pool.Get().([]byte)
    defer pool.Put(buf)
    // Use buf...
}
```

2. **Preallocate Slices**
```go
// Better
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    s = append(s, i)
}

// Worse (causes multiple reallocations)
var s []int
for i := 0; i < 1000; i++ {
    s = append(s, i)
}
```

3. **Avoid Memory Leaks**
```go
// Potential memory leak
type Cache struct {
    data map[string]*hugeStruct
}

// Better: Add cleanup method
func (c *Cache) Cleanup(maxAge time.Duration) {
    for k, v := range c.data {
        if time.Since(v.lastAccess) > maxAge {
            delete(c.data, k)
        }
    }
}
```

#### 11.2.6 GC Impact on Performance

Monitor GC performance:
```go
// Enable GC logging
debug.SetGCPercent(100)  // Default is 100
debug.SetMemoryLimit(1e9) // Set memory limit (Go 1.19+)

// Get GC statistics
var stats debug.GCStats
debug.ReadGCStats(&stats)
fmt.Printf("Last GC: %v\n", stats.LastGC)
fmt.Printf("Num GC: %d\n", stats.NumGC)
fmt.Printf("Pause Total: %v\n", stats.PauseTotal)
```

These concepts and examples provide a deeper understanding of Go's memory management and garbage collection system. Understanding these aspects helps in writing more efficient Go programs and troubleshooting memory-related issues.

## 12. Web Development

### 12.1 HTTP Server Basics

Creating a basic HTTP server:

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})
http.ListenAndServe(":8080", nil)
```

Reference: [Basic HTTP Server](webserver/basichttp.go)

### 12.2 HTTP Client

Using the `http.Get` function to make HTTP requests:

```go
resp, err := http.Get("https://example.com")
if err != nil {
    // Handle error
}
defer resp.Body.Close()
// Process the response
```

cw := ConsoleWriter{}
writeData(cw, []byte("Hello, World!"))

Reference: [HTTP Client Example](http/main.go)

### 12.3 Custom Writers

Implementing custom `io.Writer` for logging:

```go
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    return len(bs), nil
}

io.Copy(logWriter{}, resp.Body)
```

Reference: [Custom Writer Example](http/main.go)

### 12.4 Advanced HTTP Server

Using `http.ServeMux` for routing and handling different HTTP methods:

```go
mux := http.NewServeMux()
mux.HandleFunc("/", mainHandler)
mux.HandleFunc("POST /users", createUserHandler)
mux.HandleFunc("GET /users/{id}", getUserHandler)
http.ListenAndServe(":8080", mux)
```

Reference: [Advanced HTTP Server](webserver/httpadvance/main.go)

### 12.5 JSON Handling

Encoding and decoding JSON in HTTP handlers:

```go
// Decoding JSON from request body
var user User
json.NewDecoder(r.Body).Decode(&user)

// Encoding JSON to response
json.NewEncoder(w).Encode(user)
```

Reference: [JSON Handling Example](webserver/httpadvance/main.go)

### 12.6 URL Path Parameters

Extracting path parameters from URLs:

```go
id := r.PathValue("id")
intId, err := strconv.Atoi(id)
```

Reference: [Path Parameters Example](webserver/httpadvance/main.go)

### 12.7 Gorilla Mux Router

Using the Gorilla Mux router for more flexible routing:

```go
r := mux.NewRouter()
r.HandleFunc("/", homeHandler)
r.HandleFunc("/user/{id}", userHandler)
http.ListenAndServe(":8000", r)
```

Reference: [Gorilla Mux Example](webserver/gorillamux/main.go)

### 12.8 Extracting URL Variables with Gorilla Mux

Accessing URL variables in handler functions:

```go
vars := mux.Vars(r)
id := vars["id"]
```

Reference: [Gorilla Mux Variables Example](webserver/gorillamux/main.go)

These examples cover various aspects of web development in Go, including basic and advanced HTTP servers, HTTP clients, custom writers, JSON handling, URL routing, and using third-party routers like Gorilla Mux. Each topic includes a reference to the relevant Go file for more detailed implementation.

### 12.9 Gin Framework

Gin is a high-performance HTTP web framework written in Go. It provides a martini-like API with much better performance and features.

#### Installation
```bash
go get -u github.com/gin-gonic/gin
```

#### Basic Usage
```go
import "github.com/gin-gonic/gin"

func main() {
    // Create default gin router
    r := gin.Default()

    // Basic route
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    // Run on port 8080
    r.Run(":8080")
}
```

#### Route Parameters
```go
// URL Parameters
r.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(200, "Hello %s", name)
})

// Query Parameters
r.GET("/search", func(c *gin.Context) {
    query := c.DefaultQuery("q", "default search")
    c.String(200, "Search query: %s", query)
})
```

#### Request Body Binding
```go
type User struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

r.POST("/user", func(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "User created"})
})
```

#### Middleware
```go
// Custom middleware
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        
        // Set example variable
        c.Set("example", "12345")

        // before request
        c.Next()
        // after request

        latency := time.Since(t)
        log.Printf("Latency: %v", latency)
    }
}

// Using middleware
r.Use(Logger())
```

#### Groups and Versioning
```go
// API versioning
v1 := r.Group("/v1")
{
    v1.POST("/login", loginEndpoint)
    v1.POST("/submit", submitEndpoint)
    v1.POST("/read", readEndpoint)
}

v2 := r.Group("/v2")
{
    v2.POST("/login", loginEndpointV2)
    v2.POST("/submit", submitEndpointV2)
    v2.POST("/read", readEndpointV2)
}
```

#### File Upload
```go
r.POST("/upload", func(c *gin.Context) {
    file, _ := c.FormFile("file")
    
    // Save file
    c.SaveUploadedFile(file, "uploaded/"+file.Filename)
    
    c.String(200, "File %s uploaded!", file.Filename)
})
```

#### Serving Static Files
```go
// Serve single file
r.StaticFile("/favicon.ico", "./resources/favicon.ico")

// Serve directory
r.Static("/assets", "./assets")
```

#### Custom Validators
```go
type Booking struct {
    CheckIn  time.Time `form:"check_in" binding:"required,bookabledate"`
    CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn"`
}

func bookableDate(fl validator.FieldLevel) bool {
    date, ok := fl.Field().Interface().(time.Time)
    if ok {
        today := time.Now()
        if date.Unix() > today.Unix() {
            return true
        }
    }
    return false
}

func main() {
    r := gin.Default()
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("bookabledate", bookableDate)
    }
    // ... router setup
}
```

Additional examples:
- [Gin Basic Server](webserver/gin/basic/main.go)
- [Gin Advanced Features](webserver/gin/advanced/main.go)

The Gin framework provides a robust set of features for building web applications in Go, with excellent performance and a simple, expressive API. These examples cover the most common use cases, but Gin offers many more features for building complex web applications.

### 12.10 Middleware

Middleware in Go is a powerful concept, especially in the context of HTTP servers. It allows you to process requests and responses before they reach your main handler or after they've been processed by your handler.

#### Basic Structure of Middleware

The basic structure of middleware in Go is a function that takes an `http.Handler` and returns an `http.Handler`:

```go
func middlewareName(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Pre-processing logic
        next.ServeHTTP(w, r)
        // Post-processing logic
    })
}
```

#### Key Components

1. `http.Handler`: An interface with a single method:
   ```go
   type Handler interface {
       ServeHTTP(ResponseWriter, *Request)
   }
   ```

2. `http.HandlerFunc`: A type that allows regular functions to be used as HTTP handlers:
   ```go
   type HandlerFunc func(ResponseWriter, *Request)
   ```

3. `next.ServeHTTP(w, r)`: Calls the next handler in the chain.

#### Example: Logging Middleware

Here's an example of a simple logging middleware:

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
        log.Println("Finished processing request")
    })
}
```

#### Using Middleware

To use middleware in your Go HTTP server:

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)

    wrappedMux := loggingMiddleware(mux)
    log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome Home!")
}
```

#### Context in Middleware

Middleware is an excellent place to use and modify the request's context. Here's an example of middleware that adds a request ID to the context:

```go
func requestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := generateRequestID()
        ctx := context.WithValue(r.Context(), "requestID", requestID)
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
}
```

#### Chaining Middleware

You can chain multiple middleware functions:

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandler)

    wrappedMux := loggingMiddleware(requestIDMiddleware(mux))
    log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
```

In this setup, the request flows through `loggingMiddleware`, then `requestIDMiddleware`, before reaching the appropriate handler.

## 13. Database Integration

### 13.1 MongoDB with Go

MongoDB integration in Go is primarily done using the official MongoDB Go driver. Here's a comprehensive guide on working with MongoDB in Go applications:

#### Installation

First, install the MongoDB Go driver:
```bash
go get go.mongodb.org/mongo-driver/mongo
```

#### Basic Connection

```go
package main

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Connect to MongoDB
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return nil, err
    }
    
    // Ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }
    
    return client, nil
}
```

#### CRUD Operations

##### Define Model
```go
type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string            `bson:"name"`
    Email     string            `bson:"email"`
    Age       int               `bson:"age"`
    CreatedAt time.Time         `bson:"created_at"`
}
```

##### Create (Insert)
```go
// Insert One Document
func insertUser(client *mongo.Client, user User) (*mongo.InsertOneResult, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.InsertOne(ctx, user)
    if err != nil {
        return nil, err
    }
    return result, nil
}

// Insert Many Documents
func insertManyUsers(client *mongo.Client, users []interface{}) (*mongo.InsertManyResult, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.InsertMany(ctx, users)
    if err != nil {
        return nil, err
    }
    return result, nil
}
```

##### Read (Find)
```go
// Find One Document
func findUser(client *mongo.Client, filter bson.M) (*User, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    var user User
    err := collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// Find Many Documents
func findUsers(client *mongo.Client, filter bson.M) ([]User, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    
    var users []User
    if err = cursor.All(ctx, &users); err != nil {
        return nil, err
    }
    return users, nil
}
```

##### Update
```go
// Update One Document
func updateUser(client *mongo.Client, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return nil, err
    }
    return result, nil
}

// Update Many Documents
func updateManyUsers(client *mongo.Client, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.UpdateMany(ctx, filter, update)
    if err != nil {
        return nil, err
    }
    return result, nil
}
```

##### Delete
```go
// Delete One Document
func deleteUser(client *mongo.Client, filter bson.M) (*mongo.DeleteResult, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return nil, err
    }
    return result, nil
}

// Delete Many Documents
func deleteManyUsers(client *mongo.Client, filter bson.M) (*mongo.DeleteResult, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    result, err := collection.DeleteMany(ctx, filter)
    if err != nil {
        return nil, err
    }
    return result, nil
}
```

#### Advanced Operations

##### Aggregation Pipeline
```go
func aggregateUsers(client *mongo.Client, pipeline mongo.Pipeline) ([]bson.M, error) {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    cursor, err := collection.Aggregate(ctx, pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    
    var results []bson.M
    if err = cursor.All(ctx, &results); err != nil {
        return nil, err
    }
    return results, nil
}
```

##### Indexing
```go
func createIndex(client *mongo.Client, field string, unique bool) error {
    collection := client.Database("testdb").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    model := mongo.IndexModel{
        Keys:    bson.D{{field, 1}},
        Options: options.Index().SetUnique(unique),
    }
    
    _, err := collection.Indexes().CreateOne(ctx, model)
    return err
}
```

#### Usage Example
```go
func main() {
    // Connect to MongoDB
    client, err := connectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := client.Disconnect(context.Background()); err != nil {
            log.Fatal(err)
        }
    }()
    
    // Create a new user
    user := User{
        Name:      "John Doe",
        Email:     "john@example.com",
        Age:       30,
        CreatedAt: time.Now(),
    }
    
    // Insert the user
    result, err := insertUser(client, user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Inserted user ID: %v\n", result.InsertedID)
    
    // Find a user
    filter := bson.M{"email": "john@example.com"}
    foundUser, err := findUser(client, filter)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found user: %+v\n", foundUser)
}
```

#### Best Practices

1. Connection Management:
   - Always use context for timeouts
   - Properly close connections using defer
   - Implement connection pooling for production environments

2. Error Handling:
   - Check for specific MongoDB errors using `mongo.IsDuplicateKeyError(err)`
   - Implement proper retry mechanisms for transient failures
   - Use meaningful error wrapping for better debugging

3. Performance Optimization:
   - Use appropriate indexes
   - Implement pagination for large result sets
   - Use projection to limit returned fields
   - Consider using bulk operations for multiple documents

4. Security:
   - Use connection strings with authentication
   - Implement proper access control
   - Use TLS for production environments

This MongoDB integration guide provides a solid foundation for working with MongoDB in Go applications, covering basic to advanced operations with proper error handling and best practices.


## 14. Testing

### 14.1 Unit Testing

Writing and running tests:

```go
// math.go
func Add(a, b int) int {
    return a + b
}

// math_test.go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

Basic test commands:
```bash
# Run all tests in current package
go test

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run tests matching a pattern
go test -run "Test[A-Z].*"

# Run tests in all subdirectories
go test ./...
```

### 14.2 Benchmarking and Profiling

#### 14.2.1 Basic Benchmarking

Writing benchmarks:
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

Running benchmarks:
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkAdd

# Run benchmarks with custom iterations
go test -bench=. -benchtime=5s

# Show memory allocations
go test -bench=. -benchmem
```

#### 14.2.2 Profiling

1. CPU Profiling:
```bash
# Generate CPU profile
go test -cpuprofile=cpu.prof -bench=.

# Analyze with pprof
go tool pprof cpu.prof

# Generate web visualization (requires Graphviz)
go tool pprof -web cpu.prof
```

2. Memory Profiling:
```bash
# Generate memory profile
go test -memprofile=mem.prof -bench=.

# Analyze with pprof
go tool pprof mem.prof
```

3. Trace Profiling:
```bash
# Generate trace
go test -trace=trace.out -bench=.

# View trace
go tool trace trace.out
```

Common pprof interactive commands:
```bash
top           # Show top consumers
web           # Open graph visualization
list <func>   # Show line-by-line breakdown
peek          # Show parent-child relationships
quit          # Exit pprof
```

### 14.3 Coverage Testing

```bash
# Run tests with coverage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out

# Generate coverage report
go tool cover -func=coverage.out
```

### 14.4 Race Detection

```bash
# Run tests with race detector
go test -race

# Run specific test with race detector
go test -race -run TestAdd
```

### 14.5 Testify Package

The testify package provides enhanced testing capabilities with assertions, mocking, and suite support.

Installation:
```bash
go get github.com/stretchr/testify
```

#### 14.5.1 Assertions
```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
    // Simple assertions
    assert.Equal(t, 123, Calculate())
    assert.NotEqual(t, 456, Calculate())
    assert.True(t, IsValid())
    
    // Collection assertions
    assert.Contains(t, []string{"hello", "world"}, "hello")
    assert.Len(t, []int{1,2,3}, 3)
    
    // Error assertions
    err := SomeFunction()
    assert.NoError(t, err)
}
```

#### 14.5.2 Require Package
```go
import (
    "github.com/stretchr/testify/require"
)

func TestWithRequire(t *testing.T) {
    result := Setup()
    require.NotNil(t, result)  // Test stops here if result is nil
    require.Equal(t, expected, result.Value)
}
```

#### 14.5.3 Mocking
```go
type MockDB struct {
    mock.Mock
}

func (m *MockDB) Get(id string) (string, error) {
    args := m.Called(id)
    return args.String(0), args.Error(1)
}

func TestWithMock(t *testing.T) {
    mockDB := new(MockDB)
    mockDB.On("Get", "123").Return("data", nil)
    
    result, err := mockDB.Get("123")
    mockDB.AssertExpectations(t)
}
```

#### 14.5.4 Test Suites
```go
type ExampleTestSuite struct {
    suite.Suite
    db *Database
}

func (suite *ExampleTestSuite) SetupTest() {
    suite.db = NewDatabase()
}

func (suite *ExampleTestSuite) TestExample() {
    result := suite.db.Query()
    suite.Equal(expected, result)
}

func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(ExampleTestSuite))
}
```

Additional examples:
- [Test Examples](test/main.go)



## 15. Tools and Commands

### 15.1 Go Command

The `go` command is the primary tool for managing Go source code:

- `go run`: Compiles and runs a program
- `go build`: Compiles packages and dependencies
- `go test`: Runs tests
- `go get`: Downloads and installs packages and dependencies
- `go mod init`: Initializes a new module
- `go mod tidy`: Adds missing and removes unused modules

### 15.2 Formatting and Documentation

- `go fmt`: Formats Go source code
- `gofmt -s`: Simplifies code in addition to formatting
- `go doc`: Shows documentation for a package or symbol
- `godoc`: Starts a local documentation server

Additional examples:
- [Go Fix](gofix/main.go)

This example shows how to read the contents of a file using the `ioutil.ReadFile` function. Note that in more recent versions of Go, it's recommended to use `os.ReadFile` instead.

## 16. Best Practices and Patterns

### 16.1 Code Organization

- Use packages to organize code
- Follow the standard project layout
- Use meaningful names for packages, types, and functions
- Keep package names short and lowercase

### 16.2 Performance Optimization

- Use profiling tools (pprof) to identify bottlenecks
- Optimize algorithms and data structures
- Use sync.Pool for frequently allocated and deallocated objects
- Consider using sync.Map for concurrent map access instead of mutex-protected maps

### 16.3 Concurrency Patterns

- Use channels for communication, mutexes for state
- Implement the fan-out, fan-in pattern for parallel processing
- Use context for cancellation and timeouts
- Implement worker pools for managing concurrent tasks

This comprehensive guide covers all major aspects of Go programming, from basic syntax to advanced concepts and best practices. It includes detailed explanations, code examples, and practical tips to help developers write efficient and idiomatic Go code.

## 17. File Operations

Go provides robust support for file operations through various packages like `os`, `io`, `io/ioutil`, and `bufio`. Here's a comprehensive guide to file operations:

### 17.1 Reading Files

#### Small Files (< 1MB)
```go
// Method 1: os.ReadFile (Simplest)
content, err := os.ReadFile("small.txt")  // Returns []byte
if err != nil {
    log.Fatal(err)
}

// Method 2: io.ReadAll
file, err := os.Open("small.txt")         // Returns *os.File
if err != nil {
    log.Fatal(err)
}
defer file.Close()
content, err := io.ReadAll(file)          // Returns []byte

// Method 3: bufio.Scanner (line by line)
file, _ := os.Open("small.txt")           // Returns *os.File
defer file.Close()
scanner := bufio.NewScanner(file)         // Returns *bufio.Scanner
for scanner.Scan() {
    line := scanner.Text()                // Returns string
    // Process line
}
```

#### Large Files (> 1MB)
```go
// Method 1: bufio.Scanner (Memory efficient)
file, _ := os.Open("large.txt")           // Returns *os.File
defer file.Close()
scanner := bufio.NewScanner(file)         // Returns *bufio.Scanner
for scanner.Scan() {
    line := scanner.Text()                // Returns string
    // Process line
}

// Method 2: Buffered reading
file, _ := os.Open("large.txt")           // Returns *os.File
defer file.Close()
buffer := make([]byte, 1024)              // Returns []byte
for {
    n, err := file.Read(buffer)           // Returns (int, error)
    if err == io.EOF {
        break
    }
    // Process buffer[:n]
}
```

### 17.2 Writing Files

#### Simple Writing
```go
// Method 1: Write entire file
data := []byte("Hello, World!")
err := os.WriteFile("file.txt", data, 0644)

// Method 2: Create and write
file, err := os.Create("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
file.Write([]byte("Hello"))
```

#### Buffered Writing
```go
file, err := os.Create("large.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

writer := bufio.NewWriter(file)
for i := 0; i < 1000; i++ {
    writer.WriteString("Line of text\n")
}
writer.Flush()
```

### 17.3 Updating a File

#### Append to an Existing File
```go
// Adds text to the end of the file
file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("\nAppended text")
if err != nil {
    log.Fatal(err)
}
```

#### Overwrite Entire File
```go
// Replaces all content with new content
err := os.WriteFile("example.txt", []byte("New content"), 0644)
if err != nil {
    log.Fatal(err)
}
```

#### Replace Specific Text (Small Files)
```go
// Read the file
content, err := os.ReadFile("example.txt")   // Returns []byte
if err != nil {
    log.Fatal(err)
}

// Replace text and write back
newContent := strings.Replace(string(content), "old text", "new text", -1)
err = os.WriteFile("example.txt", []byte(newContent), 0644)
if err != nil {
    log.Fatal(err)
}
```

#### Process Line by Line (Large Files)
```go
// Create a temporary file
tempFile, err := os.CreateTemp("", "temp-*.txt")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(tempFile.Name())

// Open original file
original, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer original.Close()

// Process line by line
scanner := bufio.NewScanner(original)
writer := bufio.NewWriter(tempFile)
for scanner.Scan() {
    line := scanner.Text()
    // Example: Replace "old" with "new" in each line
    updatedLine := strings.Replace(line, "old", "new", -1)
    writer.WriteString(updatedLine + "\n")
}
writer.Flush()

// Replace original with updated file
os.Rename(tempFile.Name(), "example.txt")
```

### 17.4 JSON Operations

#### Small JSON Files
```go
// Reading: Direct Unmarshal
content, _ := os.ReadFile("small.json")
var data interface{}
json.Unmarshal(content, &data)

// Writing: Marshal and Write
data := map[string]string{"name": "John"}
jsonData, _ := json.Marshal(data)
os.WriteFile("output.json", jsonData, 0644)
```

#### Large JSON Files
```go
// Reading: Streaming Decoder
file, _ := os.Open("large.json")
defer file.Close()
decoder := json.NewDecoder(file)
for decoder.More() {
    var item interface{}
    err := decoder.Decode(&item)
    // Process item
}

// Writing: Streaming Encoder
file, _ := os.Create("large.json")
defer file.Close()
encoder := json.NewEncoder(file)
for item := range items {
    encoder.Encode(item)
}
```

### 17.5 Additional File Operations

```go
// Check if file exists
_, err := os.Stat("file.txt")
if os.IsNotExist(err) {
    fmt.Println("File does not exist")
}

// Rename file
err := os.Rename("old.txt", "new.txt")
if err != nil {
    log.Fatal(err)
}

// Delete file
err := os.Remove("file.txt")
if err != nil {
    log.Fatal(err)
}
```

### Best Practices
1. Always close files using `defer file.Close()`
2. Check for errors after each operation
3. Consider memory constraints when choosing a method
4. Use buffered operations for large files
5. Use appropriate buffer sizes for your use case
6. Handle errors appropriately

Reference: [File Operations](fileOperations/main.go)
These examples cover the basic file operations in Go. Remember to handle errors appropriately and close files when you're done with them. The `defer` keyword is particularly useful for ensuring files are closed properly.

# GO-Study

## Table of Contents

1. [Introduction to Go](#1-introduction-to-go)
2. [Go Basics](#2-go-basics)
3. [Functions and Methods](#3-functions-and-methods)
4. [Data Structures](#4-data-structures)
5. [Pointers](#5-pointers)
6. [Interfaces](#6-interfaces)
7. [Concurrency](#7-concurrency)
8. [Error Handling](#8-error-handling)
9. [Advanced Topics](#9-advanced-topics)
10. [Memory Management](#10-memory-management)
11. [Web Development](#11-web-development)
12. [Database Integration](#12-database-integration)
13. [Testing](#13-testing)
14. [Tools and Commands](#14-tools-and-commands)
15. [Best Practices and Patterns](#15-best-practices-and-patterns)

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

### 2.2 Data Types

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

### 2.3 Operators

Go provides various operators:

1. Arithmetic: `+`, `-`, `*`, `/`, `%`
2. Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
3. Logical: `&&`, `||`, `!`
4. Bitwise: `&`, `|`, `^`, `<<`, `>>`
5. Assignment: `=`, `+=`, `-=`, `*=`, `/=`, etc.
6. Address and Pointer: `&`, `*`

### 2.4 Control Structures

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

### 2.5 Composite Literals and Initializers

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

## 4. Data Structures

### 4.1 Arrays and Slices

Arrays have a fixed size, while slices are dynamic:

```go
// Array
arr := [3]int{1, 2, 3}

// Slice
slice := []int{1, 2, 3, 4, 5}
slice = append(slice, 6)

// Create a slice from an array
arrSlice := arr[1:3]

// Make a slice with initial length and capacity
s := make([]int, 5, 10)
```

Additional examples:

- [Basic Array and Slice Operations](arrayAndSliceREADME copy.mds/main.go)
- [2D Slices](2dslices/main.go)

2D slices example:

```go
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols)
}
```

### 4.2 Maps

Maps are key-value data structures:

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
    fmt.Println(value)
}
```

Additional examples:
- [Map Operations](Maps/main.go)

This example demonstrates various map operations, including declaration, initialization, adding elements, deleting elements, and iteration.

### 4.3 Structs

Structs are user-defined types:

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name)

// Anonymous struct
point := struct {
    x, y int
}{10, 20}
```

Additional examples:
- [Struct and Method Examples](Structs/main.go)

This file showcases the definition and use of structs in Go, as well as how to define methods on structs. It includes examples of both value receivers and pointer receivers.

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

## 6. Interfaces

### 6.1 Interface Basics

Interfaces define a set of method signatures:

```go
type Writer interface {
    Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
    n, err := fmt.Println(string(data))
    return n, err
}

func writeData(w Writer, data []byte) {
    w.Write(data)
}

cw := ConsoleWriter{}
writeData(cw, []byte("Hello, World!"))
```

Additional examples:
- [Interface Examples](Interfaces/main.go)

### 6.2 Stringer Interface

The Stringer interface is used for custom string representations:

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

### 6.3 Sort.Interface

The sort.Interface is used for custom sorting:

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
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

## 7. Concurrency

### 7.1 Goroutines

Goroutines are lightweight threads managed by the Go runtime:

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

### 7.2 Channels

Channels are used for communication between goroutines:

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

This file demonstrates various channel operations, including creating channels, sending and receiving values, and using channels for communication between goroutines. It also shows how to use channels to implement a simple worker pool pattern.

### 7.3 Select Statement

The select statement is used to work with multiple channels:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel operations ready")
}
```

### 7.4 Synchronization Primitives

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

### 7.5 Concurrency Patterns

- [Context Example](concurrencypatterns/context-example/main.go)
- [Generator](concurrencypatterns/generator/)
- [Confinement](concurrencypatterns/confinement/)
- [Prime Fan-in Fan-out](concurrencypatterns/prime-fan-in-fan-out/)
- [Mutex Map](concurrencypatterns/mutex-map/)
- [Mutex Example](concurrencypatterns/mutex-example/)
- [Prime Pipeline](concurrencypatterns/prime-pipeline/)
- [Or Done](concurrencypatterns/or-done/)

The context example demonstrates the use of the `context` package for managing goroutine lifecycles and cancellation. It shows how to create a context with a timeout and use it to control multiple goroutines.

## 8. Error Handling

### 8.1 Error Interface

Go uses the `error` interface for error handling:

```go
type error interface {
    Error() string
}

if err != nil {
    fmt.Println("Error:", err)
    return
}
```

### 8.2 Custom Errors

You can create custom error types:

```go
type MyError struct {
    message string
}

func (e *MyError) Error() string {
    return e.message
}

func doSomething() error {
    return &MyError{"something went wrong"}
}
```

### 8.3 Panic and Recover

Panic is used for unrecoverable errors, and recover can catch panics:

```go
func mayPanic() {
    panic("a problem")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    mayPanic()
}
```

### 8.4 Error Wrapping

Go 1.13+ introduced error wrapping:

```go
if err != nil {
    return fmt.Errorf("failed to process data: %w", err)
}

// Unwrapping
if errors.Is(err, io.EOF) {
    // handle EOF
}

var myErr *MyError
if errors.As(err, &myErr) {
    // handle MyError
}
```

## 9. Advanced Topics

### 9.1 Reflection

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

### 9.2 Generics

Go 1.18+ supports generics:

```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

// Usage
PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"a", "b", "c"})
```

### 9.3 Context Package

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

## 10. Memory Management

### 10.1 new vs make

`new` and `make` are built-in functions for memory allocation:

```go
// new returns a pointer to a zero-initialized value
p := new(int)
fmt.Println(*p) // Output: 0

// make is used to create slices, maps, and channels
s := make([]int, 5, 10)
m := make(map[string]int)
ch := make(chan int, 5)
```

### 10.2 Garbage Collection

Go uses a concurrent mark-and-sweep garbage collector. While it's mostly automatic, you can influence it:

```go
runtime.GC() // Force a garbage collection

// Set GOGC environment variable to control GC frequency
// GOGC=50 means GC will run when heap size is 50% larger than after the previous GC
```

## 11. Web Development

### 11.1 HTTP Server Basics

Creating a basic HTTP server:

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})
http.ListenAndServe(":8080", nil)
```

### 11.2 Routing and Handlers

Using the `gorilla/mux` router:

```go
r := mux.NewRouter()
r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
http.ListenAndServe(":8080", r)
```

### 11.3 Middleware

Creating middleware:

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

r.Use(loggingMiddleware)
```

### 11.4 Templates

Using HTML templates:

```go
tmpl, err := template.ParseFiles("template.html")
if err != nil {
    log.Fatal(err)
}
tmpl.Execute(w, data)
```

### 11.5 RESTful APIs

Creating a RESTful API:

```go
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    // Fetch user from database
    user := User{ID: id, Name: "John Doe"}
    json.NewEncoder(w).Encode(user)
}
```

Additional examples:
- [Basic Web Server](webserver/gorillamux/main.go)

This example uses the Gorilla Mux router to create a more feature-rich web server compared to the standard library's `http.ServeMux`.

## 12. Database Integration

Using the `database/sql` package with PostgreSQL:

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

rows, err := db.Query("SELECT id, name FROM users WHERE id = $1", 1)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()
```

## 13. Testing

### 13.1 Unit Testing

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

// Run tests
// go test
```

### 13.2 Benchmarking

Writing and running benchmarks:

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

// Run benchmarks
// go test -bench=.
```

Additional examples:
- [Test Examples](test/main.go)

This file contains various examples of Go language features, including switch statements, type assertions, and generics. While not strictly a test file, it demonstrates how different language constructs can be used and tested.

## 14. Tools and Commands

### 14.1 Go Command

The `go` command is the primary tool for managing Go source code:

- `go run`: Compiles and runs a program
- `go build`: Compiles packages and dependencies
- `go test`: Runs tests
- `go get`: Downloads and installs packages and dependencies
- `go mod init`: Initializes a new module
- `go mod tidy`: Adds missing and removes unused modules

### 14.2 Formatting and Documentation

- `go fmt`: Formats Go source code
- `gofmt -s`: Simplifies code in addition to formatting
- `go doc`: Shows documentation for a package or symbol
- `godoc`: Starts a local documentation server

Additional examples:
- [Go Fix](gofix/main.go)

This example shows how to read the contents of a file using the `ioutil.ReadFile` function. Note that in more recent versions of Go, it's recommended to use `os.ReadFile` instead.

## 15. Best Practices and Patterns

### 15.1 Code Organization

- Use packages to organize code
- Follow the standard project layout
- Use meaningful names for packages, types, and functions
- Keep package names short and lowercase

### 15.2 Performance Optimization

- Use profiling tools (pprof) to identify bottlenecks
- Optimize algorithms and data structures
- Use sync.Pool for frequently allocated and deallocated objects
- Consider using sync.Map for concurrent map access instead of mutex-protected maps

### 15.3 Concurrency Patterns

- Use channels for communication, mutexes for state
- Implement the fan-out, fan-in pattern for parallel processing
- Use context for cancellation and timeouts
- Implement worker pools for managing concurrent tasks

This comprehensive guide covers all major aspects of Go programming, from basic syntax to advanced concepts and best practices. It includes detailed explanations, code examples, and practical tips to help developers write efficient and idiomatic Go code.
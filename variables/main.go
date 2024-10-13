package main

import "fmt"

var packageVar = "I'm a package-level variable"

func main() {
    localVar := "I'm a local variable"
    fmt.Println(packageVar)
    fmt.Println(localVar)
	x := 10

    // Re-assignment (allowed)
    x = 20

    // Re-declaration (allowed because y is new)
    x, y := 30, 40

    // This would be an error:
    // x := 50  // Cannot re-declare x alone

    // This is allowed (re-assignment of x and y)
    x, y = 60, 70
	fmt.Println(x, y)

    // More examples of variable declaration and initialization
    var a int // Declaration without initialization
    var b int = 5 // Declaration with initialization
    c := 10 // Short variable declaration
    var d, e int = 15, 20 // Multiple variable declaration
    f, g := 25, "hello" // Multiple variable declaration with type inference
    var (
        h int
        i string
        j bool
    ) // Block declaration
    const k = 100 // Constant declaration
    
    fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
}
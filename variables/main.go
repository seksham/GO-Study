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
}
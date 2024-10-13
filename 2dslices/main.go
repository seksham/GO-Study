package main

import "fmt"

func main() {
    // Declare a 2D slice
    var matrix [][]int

    // Initialize the 2D slice with 3 rows and 4 columns
    rows, cols := 3, 4
    matrix = make([][]int, rows)
    for i := range matrix {
        matrix[i] = make([]int, cols)
    }

    // Populate the matrix
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            matrix[i][j] = i*cols + j
        }
    }

    // Print the matrix
    fmt.Println("2D Slice (Matrix):")
    for _, row := range matrix {
        fmt.Println(row)
    }
    // Output:
    // 2D Slice (Matrix):
    // [0 1 2 3]
    // [4 5 6 7]
    // [8 9 10 11]

    // Accessing elements
    fmt.Printf("Element at row 1, column 2: %d\n", matrix[1][2])
    // Output: Element at row 1, column 2: 6

    // Modifying an element
    matrix[2][3] = 100
    fmt.Println("After modification:")
    for _, row := range matrix {
        fmt.Println(row)
    }
    // Output:
    // After modification:
    // [0 1 2 3]
    // [4 5 6 7]
    // [8 9 10 100]
}
package main

import "fmt"

func main() {
	rows, columns := 3, 4
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
	}
	fmt.Println(matrix)
	// for i:= 0; i<rows; i++{
	// 	for j := 0; j<columns; j++{
	// 		matrix[i][j] = j+columns*i
	// 	}
	// }
	for i, row := range matrix{
		for j, _ := range row{
			row[j] = i*columns+j
		}
	}
	
	fmt.Println(matrix)
	 // Print the matrix
	 fmt.Println("2D Slice (Matrix):")
	 for i, row := range matrix {
		 fmt.Println(i,row)
	 }
}

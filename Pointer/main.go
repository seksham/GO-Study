package main
 
import "fmt"
 
func main() {
 name := "bill"
 
 namePointer := &name
 fmt.Println(namePointer)  // Output: 0xc000010230 (memory address of 'name')
 fmt.Println(&namePointer) // Output: 0xc00000e028 (memory address of 'namePointer')
 printPointer(namePointer)
}
 
func printPointer(namePointer *string) {
 fmt.Println(namePointer)  // Output: 0xc000010230 (same as in main())
 fmt.Println(&namePointer) // Output: 0xc00000e038 (different address, local to the function)
}

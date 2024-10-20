package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		opp := n - 1 - i
		runes[i], runes[opp] = runes[opp], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseString("Hello saksham"))
	fmt.Println(reverseString("你好，世界"))
	fmt.Println(reverseString("こんにちは、世界"))
	fmt.Println(reverseString("Γειά σου Κόσμε"))
	fmt.Println(reverseString("Hello 🌍🌎🌏"))
	fmt.Println(reverseString("!@#$%^&*()_+"))

}

package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

func wordFrequency(s string) map[string]int {
	frequencyMap := make(map[string]int)
	words := strings.Fields(strings.ToLower(s))
	for _, word := range words {
		match, err := regexp.Compile("[^A-Za-z]")
		if err == nil {
			word = string(match.ReplaceAll([]byte(word), []byte("")))
		}
		frequencyMap[word]++
		fmt.Println(word)
	}
	return frequencyMap
}

func countWords(paragraph string) map[string]int {
	wordCount := make(map[string]int)

	// Convert to lowercase and split into words
	words := strings.Fields(strings.ToLower(paragraph))

	for _, word := range words {
		// Remove non-letter characters
		cleanWord := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) {
				return r
			}
			return -1
		}, word)

		if cleanWord != "" {
			wordCount[cleanWord]++
		}
	}

	return wordCount
}

func main() {
	paragraph := `This is a random paragraph with repeated words. words can appear multiple times, and punctuation,
	 such as commas, periods, and exclamation marks, can also be included. This random sentence serves as an example. A dog is a dog's enemy`
	f := wordFrequency(paragraph)
	fmt.Println(f)
	f2 := countWords(paragraph)
	fmt.Println(f2)
	fmt.Println(reflect.DeepEqual(f, f2))
}

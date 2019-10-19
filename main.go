package main

import (
	"fmt"
	"strings"
	"sort"
)

func SortWord(word string) string{
	var vowels []string
	var consonants []string
	var print []string

	words := strings.Split(word, "")

	for _, char := range words {
		switch char {
			case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
				vowels = append(vowels, char)
			default:
				consonants = append(consonants, char)
		}
	}

	sort.Strings(vowels)
	sort.Strings(consonants)

	print = append(vowels, consonants...)
	result := strings.Join(print, "")

	return result
}

func main() {
	var word string

	fmt.Print("Masukkan Kata: ")
	fmt.Scanln(&word)

	fmt.Printf("Hasil: %s", SortWord(word))
}
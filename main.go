package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	var word string
	var vowels []string
	var consonants []string
	var print []string

	fmt.Print("Masukkan Kata: ")
	fmt.Scanln(&word)

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

	fmt.Printf("Hasil: %s", result)
}
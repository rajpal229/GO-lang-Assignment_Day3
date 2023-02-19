package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	var v_count, c_count int
	for i, v := range str {
		if unicode.IsLetter(v) {
			vowels := "aeiou"
			if strings.Contains(vowels, string(v)) {
				v_count = v_count + 1
				fmt.Printf("%v at index %v is a vowel\n", string(v), i)
			} else {
				c_count = c_count + 1
				fmt.Printf("%v at index %v is a consonant\n", string(v), i)
			}
		}
	}
	fmt.Printf("Vowel count = %d & Consonant count = %d", v_count, c_count)
}

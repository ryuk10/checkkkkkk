package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}

	s1 := map[rune]int{}
	s2 := map[rune]int{}
	uniqueCounter := 0

	for _, char := range str1 {
		s1[char]++
	}
	for _, char := range str2 {
		s2[char]++
	}

	for char, num := range s1 {
		if num == 1 && s2[char] == 0 {
			uniqueCounter++
		}
	}
	for char, num := range s2 {
		if num == 1 && s1[char] == 0 {
			uniqueCounter++
		}
	}
	return uniqueCounter
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

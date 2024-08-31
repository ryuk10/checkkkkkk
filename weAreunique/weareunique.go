package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) []string {
	m := make(map[rune]int)

	if str1 == "" || str2 == "" {
		return nil
	}

	

	
	
}

func main() {
	fmt.Println(WeAreUnique("foo", "vboo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "daef"))
}

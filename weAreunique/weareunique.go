package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) []string {
	
	var unique []string

	if str1 == "" || str2 == "" {
		return nil
	}

	for i := 0; i <= len(str1)-1; i++ {
		for j := 0; j <= len(str2)-1; j++ {
			if str1[i] == str2[j] {
				
			}
		}
	}
}

func main() {
	fmt.Println(WeAreUnique("foo", "vboo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "daef"))
}

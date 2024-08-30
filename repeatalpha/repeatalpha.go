package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	res := ""
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			for j := 0; j <= int(s[i])-int('A'); j++ {
				res += string(s[i])
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			for j := 0; j <= int(s[i])-int('a'); j++ {
				res += string(s[i])
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

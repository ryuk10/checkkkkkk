package main

import "fmt"

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}
	res := ""
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == ' ' {
			res = s[:i] + "\n"
			break
		}
	}
	return res
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

package main

import "fmt"

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}
	res := ""
	finres := ""
	for i := 0; i <= len(s)-1; i++ {
		if s[i] != ' ' {
			res += string(s[i])
		}
		if i < len(s)-1 && s[i] != ' ' && s[i+1] == ' ' {
			break
		}
	}
	finres = res + "\n"
	return finres
}

func main() {
	fmt.Print(FirstWord("       hello       there"))
	fmt.Print(FirstWord("hej"))
	fmt.Print(FirstWord("      hello     .........  bye"))
}

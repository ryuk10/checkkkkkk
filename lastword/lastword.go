package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum  "))
	fmt.Print(LastWord("     "))
}

func LastWord(s string) string {
	runes := []rune{}     //
	strings := []string{} //["this" , "" , "" , "..." "" "" "" "is" ]
	for i := range s {
		if s[i] != ' ' {
			runes = append(runes, rune(s[i]))
		}
		if s[i] == ' ' || i == len(s)-1 {
			if string(runes) != "" {
				strings = append(strings, string(runes))
				runes = []rune{}
			}
		}
	}
	if len(strings) == 0 {
		return "\n"
	}
	return strings[len(strings)-1] + "\n"
}

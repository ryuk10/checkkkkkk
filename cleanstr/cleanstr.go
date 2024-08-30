package main

import (
	"fmt"
	"os"
)

func main() {
	// args := os.Args[1:]

	// if len(args) != 1 {
	// 	fmt.Println()
	// 	return
	// }

	// input := args[0]

	// str := ""

	// for i := 0; i <= len(input)-1; i++ {
	// 	if string(input[i]) != " " {
	// 		str += string(input[i])
	// 	}
	// 	if i != 0 && i < len(input)-1 && string(input[i]) == " " && string(input[i+1]) != " " {
	// 		str += " "
	// 	}
	// }
	// fmt.Println(str)

	// runes := []rune{}
	str := []string{}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println()
		return
	}
	input := args[0]

	for i := 0; i <= len(input)-1; i++ {
		if input[i] != ' ' {
			str = append(str, string(input[i]))
		} else if i != 0 && i < len(input)-1 && input[i] == ' ' && input[i+1] != ' ' {
			str = append(str, " ")
		}
	}
	for _, r := range str {
		fmt.Print(r)
	}
	fmt.Println()
}

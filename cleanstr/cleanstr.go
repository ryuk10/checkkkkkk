package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println()
		return
	}

	input := args[0]

	str := ""

	for i := 0; i <= len(input)-1; i++ {
		if string(input[i]) != " " {
			str += string(input[i])
		}
		if i != 0 && i < len(input)-1 && string(input[i]) == " " && string(input[i+1]) != " " {
			str += " "
		}
	}
	fmt.Println(str)
}

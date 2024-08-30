package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	inputstring := args[0]
	char := args[1]
	rep := args[2]

	for _, r := range char {
		for i := 0; i <= len(inputstring)-1; i++ {
			if rune(inputstring[i]) == r {
				fmt.Print(rep)
				continue
			}
			fmt.Print(string(inputstring[i]))
		}
		fmt.Println()
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// str := []string{}

	// args := os.Args[1:]

	// if len(args) != 1 {
	// 	return
	// }
	// input := args[0]
	// count := 0

	// for i := 0; i <= len(input)-1; i++ {

	// 	if input[i] != ' ' {
	// 		count++
	// 		str = append(str, string(input[i]))
	// 	}
	// 	if count == 0 {
	// 		continue
	// 	}
	// 	if i < len(input)-1 && input[i+1] != ' ' && input[i] == ' ' {
	// 		str = append(str, "   ")
	// 	}

	// }

	// for _, r := range str {
	// 	fmt.Print(r)
	// }
	// fmt.Println()

	res := ""
	fin := 0

	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	input := args[0]

	for i := 0; i <= len(input)-1; i++ {
		if input[i] != ' ' {
			fin++
			res += string(input[i])
		}
		if fin == 0 {
			continue
		}
		if i < len(input)-1 && input[i+1] != ' ' && input[i] == ' ' {
			res += "   "
		}
	}
	for _, r := range res {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

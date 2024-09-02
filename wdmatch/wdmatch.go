package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	a := args[0]
	b := args[1]
	j := 0

	res := ""

	for i := 0; i <= len(a)-1; i++ {
		for j <= len(b)-1 {
			if a[i] == b[j] && i < len(a) && j < len(b) {
				res += string(a[i])
				break
			}
			j++
		}
	}

	if len(a) > len(res) {
		return
	} else {
		for _, r := range res {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

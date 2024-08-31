package main

import (
	"fmt"
)

func Itoa(n int) string {
	var sl []string
	res := ""
	m := ""

	if n < 0 {
		m = "-"
		n *= -1
	}

	for n > 0 {
		mod := n % 10
		sl = append(sl, string(mod+'0'))
		n = n / 10
	}
	
	
	

	for i := len(sl) - 1; i >= 0; i-- {
		res += sl[i]
	}
	return m + res
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

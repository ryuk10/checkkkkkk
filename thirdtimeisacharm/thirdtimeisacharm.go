package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	result := ""
	for i := 0; i <= len(str)-1; i++ {
		if i+2 > len(str)-1 {
			break
		}
		result += string(str[i+2])
		i += 2
	}

	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

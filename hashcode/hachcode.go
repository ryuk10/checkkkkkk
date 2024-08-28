package main

import (
	"fmt"
)

func HashCode(dec string) string {
	res := ""
	result := 0
	for _, r := range dec {
		result = (int(r) + len(dec)) % 127
		if result >= 0 && result <= 127 {
			res += string(result)
		}
	}

	return res
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BACé"))
	fmt.Println(HashCode("Hello World"))
}

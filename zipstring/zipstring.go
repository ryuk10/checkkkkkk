package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllasssssss"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!          "))
	// fmt.Println(ZipString("heeeff"))
}

func ZipString(s string) string {
	result := ""
	count := 1
	for i := 0; i <= len(s)-1; i++ {
		if i < len(s)-1 && s[i] == s[i+1] {
			count++
			continue
		}
		result += strconv.Itoa(count) + string(s[i])
		count = 1

	}
	return result
}

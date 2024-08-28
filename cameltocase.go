package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	for i := 0; i <= len(s)-1; i++ {
		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
			continue
		}

		if s[i] >= 'a' && s[i] <= 'z' {
			if i != len(s)-1 && s[i+1] >= 65 && s[i+1] <= 90 {
				s = s[:i+1] + "_" + s[i+1:]
			}
		}
	}
	return s
}

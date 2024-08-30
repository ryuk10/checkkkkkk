package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) string {
	var d []string
	// var unique []string
	res := ""

	if str1 == "" || str2 == "" {
		return ""
	}

	for i := 0 ; i <= len(str1)-1 ; i++ {
		for j := 0 ; j <= len(str2)-1 ; j++ {
			if str1[i] == str2[j] {
				d = append(d, string(str1[i]))
			}else {
				res = string(str1[i])
			}
		}
	}
	for i := 0 ; i <= len(str2)-1 ; i++ {
		for j := 0 ; j <= len(str1)-1 ; j++ {
			if str2[i] == str1[j] {
				d = append(d, string(str2[i]))
			}else {
				res = string(str2[i])
			}
		}
	}
	return d[1]
	
}

func main() {
	fmt.Println(WeAreUnique("foo", "vboo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "daef"))
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	result := ""
	count := 1
	for i := 0 ; i<= len(s)-1 ; i++ {
		if s[i] == s[i+1] {
			count++
			
		}else {

		}
	}
	

}
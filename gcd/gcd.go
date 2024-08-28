package main

import (
	"fmt"
	
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}


func Gcd(a, b uint) uint {

	if a < b {
		for i := a ; i > 0 ; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
	} else {
		for i := b ; i > 0 ; i-- {
			if b%i == 0 && a%i == 0 {
				return i
			}
		}
	}
	return 0

}
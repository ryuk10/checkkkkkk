package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}


func FindPrevPrime(nb int) int {
    if nb <= 2 {
        return 0
    }

    for i := nb; i >= 2; i-- {
        if isPrime(i) {
            return i
        }
    }

    return 0
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
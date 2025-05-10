package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is a negative number")
	} else if num < 10 {
		fmt.Println(num, "is one digit number")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

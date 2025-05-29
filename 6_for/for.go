package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	j := 1
	for j < 3 {
		fmt.Println(j)
		j = j + 1
	}

	fmt.Println("----")

	for k := range "Five" {
		println(k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range "SIXSIX" {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}

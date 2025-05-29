package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[3] = 100
	fmt.Println("Array's first version", a)
	fmt.Println("Array's fourth element", a[3])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//compiler count
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Indexed Initialization
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b) //[100 0 0 400 500]
}

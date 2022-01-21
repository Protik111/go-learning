package main

import "fmt"

//public variable
var fullName string = "Rafiur Rahman Protik"

func main() {
	fmt.Println("Hello Variable")

	var counts int
	fmt.Printf("Type of counts is : %T \n", counts)

	var name = "protik"
	fmt.Println("My name is :", name)

	lastname := "rahman"
	fmt.Println("lastname is", lastname)

	fmt.Println("Full Name is", fullName)
	fmt.Printf("Full name type is %T \n", fullName)
}

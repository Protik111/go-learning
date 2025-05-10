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

	var num int = 10
	fmt.Println(num)

	var num2 = 20
	fmt.Println(num2)

	a := 30.50
	fmt.Println(a)

	title := "Learning Go!"
	fmt.Println(title)

	//changing variable
	//name := "Protik"
	name = "Rafiur Rahman Protik"
	//name = false       //can't change the data type while changing the value
	fmt.Println(name)

	const p = 100
	fmt.Println(p)
}

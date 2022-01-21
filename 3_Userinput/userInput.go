package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Start taking user input")

	reader := bufio.NewReader(os.Stdin)

	//comma ok // comma error

	input, _ := reader.ReadString('\n')
	fmt.Println("The user input is :", input)
	//fmt.Printf("The type of input is : %T \n", input)

	newNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number after added", newNumber+10)
	}

}

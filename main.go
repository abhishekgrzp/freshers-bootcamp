package main

import (
	"fmt"
)

func main() {
	//First program Hello World
	fmt.Println("hello world")

	defer fmt.Println("defer statement executes at last.")

	//Variables and data type
	var isBool, isTrue bool = true, false
	isBool = false
	isTrue = true
	fmt.Println("isBool:", isBool)
	fmt.Println("isTrue:", isTrue)

	withoutVar := 14
	fmt.Println("short assignment:", withoutVar)

	runeVar := 'A'
	fmt.Println("ASCII value for A:", runeVar)

	//For loop
	for i := 0; i < 3; i++ {
		fmt.Println("number", i)
	}

	//while loop
	sum := 1
	for sum < 6 {
		sum++
	}
	fmt.Println("sum", sum)
}

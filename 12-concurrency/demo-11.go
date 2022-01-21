package main

import (
	"fmt"
	"time"
)

/*
	Assignment:
		Print "Hello" & "World" simultaneously
		IMPORTANT : Do NOT move the loop outside the "print" function
*/

func main() {
	print("Hello")
	print("World")
	fmt.Println("End of main!")
}

func print(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}
}

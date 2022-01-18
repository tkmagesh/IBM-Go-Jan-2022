package main

import "fmt"

func main() {
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is even\n", no)
		} else {
			fmt.Printf("%d is odd\n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is even\n", no)
	} else {
		fmt.Printf("%d is odd\n", no)
	}
	fmt.Println(no)
}

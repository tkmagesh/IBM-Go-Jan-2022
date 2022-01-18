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

	//for
	//variation - 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum = %v\n", sum)

	//variation - 2 (using like a while loop)
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Printf("numSum = %v\n", numSum)

	//variation - 3 (infinite loop)
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Printf("x = %v\n", x)
}

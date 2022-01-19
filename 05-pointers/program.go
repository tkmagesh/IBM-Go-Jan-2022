package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(no, noPtr)

	//deferencing (address -> value)
	var x = *noPtr
	fmt.Println(x, noPtr)

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	var n1, n2 = 10, 20
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n")
	swap( /*  */ )
	fmt.Printf("After swapping n1 = %d and n2 = %d\n")

}

func increment(no *int) {
	*no++
}

func swap( /*  */ ) {

}

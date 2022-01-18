package main

import "fmt"

func main() {
	//var x int

	//var x int = 100

	//type inference
	//var x = 100
	//var x = "Hi there"

	x := 100 /* ":=" is applicable only in function  and not for package level variables */
	//fmt.Println(x)
	fmt.Printf("value of x is %v of type %T\n", x, x)

	/*
		var n1 int
		var n2 int
		var result int
		var str string
		n1 = 10
		n2 = 20
		result = n1 + n2
		str = "n1(%d) + n2(%d) = %d\n"
	*/
	/*
		var n1, n2, result int
		var str string
		n1 = 10
		n2 = 20
		result = n1 + n2
		str = "n1(%d) + n2(%d) = %d\n"
	*/

	/*
		var (
			n1, n2, result int
			str            string
		)
		n1 = 10
		n2 = 20
		result = n1 + n2
		str = "n1(%d) + n2(%d) = %d\n"
	*/

	/*
		var (
			n1, n2 int    = 10, 20
			result int    = n1 + n2
			str    string = "n1(%d) + n2(%d) = %d\n"
		)
	*/

	/*
		var (
			n1, n2 = 10, 20
			result = n1 + n2
			str    = "n1(%d) + n2(%d) = %d\n"
		)
	*/

	/*
		var (
			n1, n2, str = 10, 20, "n1(%d) + n2(%d) = %d\n"
			result      = n1 + n2
		)
	*/

	/*
		n1, n2 := 10, 20
		result := n1 + n2
		str := "n1(%d) + n2(%d) = %d\n"
	*/

	n1, n2, str := 10, 20, "n1(%d) + n2(%d) = %d\n"
	result := n1 + n2

	fmt.Printf(str, n1, n2, result)
}

package main

import "fmt"

func main() {
	//Array
	fmt.Println("Arrays")
	//var nos [5]int
	//var nos [5]int = [5]int{4, 1, 2, 5, 3}
	nos := [5]int{4, 1, 2, 5, 3}
	fmt.Println(nos)

	fmt.Println("Iterating an array (using indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	/*
		fmt.Println("Iterating an array (using range)")
		for idx, val := range nos {
			fmt.Printf("nos[%d] = %d\n", idx, val)
		}
	*/

	fmt.Println("Iterating an array (using range)")
	for _, val := range nos {
		fmt.Printf("%d\n", val)
	}

	fmt.Println("Copying an array")
	newNos := nos
	nos[2] = 100
	fmt.Println(newNos, nos)
}

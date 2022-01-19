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

	fmt.Println()
	fmt.Println("Slice")
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	//var products []string
	var products []string = make([]string, 3, 10)
	fmt.Println(products, len(products), cap(products))
	products[0] = "Pen"
	products[1] = "Pencil"
	products[2] = "Marker"
	fmt.Println(products, len(products), cap(products))
	/*
		products = append(products, "Pencil", "Marker", "Duster")
		fmt.Println(products, len(products), cap(products))
		products = append(products, "Mouse")
		fmt.Println(products, len(products), cap(products))
		products = append(products, "Scribble-pad", "Stylus")
		fmt.Println(products, len(products), cap(products))
		products = append(products, "Printer")
		fmt.Println(products, len(products), cap(products)) */

	fmt.Println()
	fmt.Println("Map")
	//var productRanks map[string]int = make(map[string]int)
	/*
		productRanks := map[string]int{}
		productRanks["Pen"] = 4
	*/
	productRanks := map[string]int{
		"Pen":    4,
		"Pencil": 1,
		"Marker": 5,
	}
	productRanks["Printer"] = 3
	productRanks["Scribble-pad"] = 2
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("Rank of %q is %d\n", key, val)
	}

	fmt.Println("Cheching if a key exists")
	productToSearch := "Dummy"
	if val, exists := productRanks[productToSearch]; exists {
		fmt.Printf("Rank of %q is %d\n", productToSearch, val)
	} else {
		fmt.Printf("%q does not exist\n", productToSearch)
	}

	fmt.Println("Deleting from a map")
	delete(productRanks, "Pen")
	fmt.Println("productRanks after deleting Pen", productRanks)
}

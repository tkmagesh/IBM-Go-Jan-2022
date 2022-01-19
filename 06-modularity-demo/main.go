package main

/*
import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Printf("Op Count = %d\n", calculator.GetOpCount())
	fmt.Println(utils.IsPrime(97))
}
*/

import (
	"fmt"
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Printf("Op Count = %d\n", calc.GetOpCount())
	fmt.Println(utils.IsPrime(97))
}

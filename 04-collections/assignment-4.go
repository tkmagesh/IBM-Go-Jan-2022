package main

import "fmt"

var operations map[int]func(int, int) int = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	var n1, n2, userChoice, result int
	for {
		userChoice = getUserChoice()
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		n1, n2 = getOperands()
		result = performOperation(userChoice, n1, n2)
		fmt.Println("result = ", result)
	}
}

func performOperation(userChoice, n1, n2 int) (result int) {
	if operation, exists := operations[userChoice]; exists {
		result = operation(n1, n2)
		return
	}
	return
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the values :")
	fmt.Scanf("%d %d\n", &n1, &n2)
	return
}

func getUserChoice() (userChoice int) {
	fmt.Println("Enter the choice:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanf("%d\n", &userChoice)
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

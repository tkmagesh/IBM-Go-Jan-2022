package main

import (
	"fmt"
	"structs-demo/models"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	//var product Product
	//var product Product = Product{101, "Pen", 10, 100, "Stationary"}
	/*
		var product Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	//var product Product = Product{Id: 100, Name: "Pen", Cost: 10}
	//var product = Product{Id: 100, Name: "Pen", Cost: 10}
	product := Product{Id: 100, Name: "Pen", Cost: 10}
	/*
		var p2 = Product{Id: 100, Name: "Pen", Cost: 10}
		fmt.Println(product == p2)
		fmt.Println(ToString(p2))
	*/

	/*
		var productPtr *Product = &product
		//fmt.Println((*productPtr).Id)
		fmt.Println(productPtr.Id)
	*/

	fmt.Println(ToString(&product))

	/* emp := &Employee{
		Id:   5001,
		Name: "Magesh",
		Org: Organization{
			Name: "IBM",
			City: "Bengaluru",
		},
	} */
	//emp := Employee{}
	//emp := new(Employee) //returns a pointer to the Employee object
	/*
		emp := &Employee{}
		fmt.Printf("%#v\n", emp)
		fmt.Println(emp.Org.City)
	*/
	emp := models.NewEmployee(5001, "Magesh", "IBM", "Bengaluru")
	fmt.Printf("%#v\n", emp)

}

func ToString(product *Product) string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%f, Units=%d, Category=%q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

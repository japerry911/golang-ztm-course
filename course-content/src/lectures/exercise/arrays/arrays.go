//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func main() {
	products := [4]Product{
		{name: "xbox", price: 350},
		{name: "ps4", price: 500},
		{name: "gamecube", price: 175},
	}

	fmt.Println(products[2])
	fmt.Println(len(products))

	totalCost := 0
	for i := 0; i < len(products); i++ {
		product := products[i]
		totalCost += product.price
	}
	fmt.Println(totalCost)

	products[3] = Product{name: "computer", price: 2000}
	fmt.Println(products[3])
}

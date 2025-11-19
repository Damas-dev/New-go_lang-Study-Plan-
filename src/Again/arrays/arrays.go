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

// ~Products must include the price and the name
type Product struct {
	name  string
	price int
}

func totalCost(shoppingList [4]Product) {
	var totalCost int
	for i := 0; i < len(shoppingList); i++ {

		product := shoppingList[i]
		totalCost += product.price
	}
	println("totalCost:", totalCost)
}

func main() {

	fmt.Println("this is cool")

	//? Using an array, create a shopping list with enough room for 4 products
	//* Insert 3 products into the array
	shoppingList := [4]Product{
		{"eggs", 12000},
		{"10fruits", 20000},
		{"3spoons", 6000},
	}
	//* Print to the terminal:
	//~ The last item on the list
	fmt.Println(shoppingList[len(shoppingList)-1])
	//~The total number of items
	fmt.Println("total produts in my shopping list", len(shoppingList))
	//~The total cost of the items

	totalCost(shoppingList)
	shoppingList[3].name = "bag"
	shoppingList[3].price = 30000

	fmt.Println(shoppingList[len(shoppingList)-1])

	totalCost(shoppingList)

}

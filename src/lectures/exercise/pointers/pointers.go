//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
const (
	active   = true
	inactive = false
)

type Item struct {
	name  string
	state bool
}

// * Create functions to activate and deactivate security tags using pointers
func activate(i *Item) {
	i.state = active

}

func deactivate(i *Item) {
	i.state = inactive

}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(i *[]Item) {
	for _, v := range *i {
		v.state = inactive
		fmt.Println(v.name, ":inactive")
	}

}

func main() {

	//  - Create at least 4 items, all with active security tags

	car1 := Item{"car1", active}
	car2 := Item{"car2", active}
	car3 := Item{"car3", active}
	car4 := Item{"car4", active}
	car5 := Item{"car5", active}

	//  - Store them in a slice or array
	items := []Item{
		car1,
		car2,
		car3,
		car4,
		car5,
	}

	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&car5)
	fmt.Println(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(&items)
	fmt.Println(items)

	activate(&car3)

	fmt.Println(items)

}

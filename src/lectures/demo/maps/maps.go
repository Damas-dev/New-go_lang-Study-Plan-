package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is maps")
	shopingList := make(map[string]int)

	shopingList["eggs"] = 11
	shopingList["bread"] += 1
	shopingList["caps"] = 3

	fmt.Println(shopingList)

	delete(shopingList, "caps")

	fmt.Println("caps are deleted, new list", shopingList)

	shopingList["eggs"] += 1

	fmt.Println("their", shopingList["eggs"], "eggs in the shopping list")

	fmt.Println(shopingList)

	fmt.Println("\n---what I need to bay---")

	fmt.Println("I need to bay:", shopingList["eggs"], "eggs &", shopingList["bread"])

	fmt.Println("\nDo I need caps?")

	caps, found := shopingList["caps"]
	if found {
		fmt.Println(" yap I need:", caps, "caps")
	} else {
		fmt.Println(" nope")
	}

	totalAmount := 0
	for _, amount := range shopingList {
		totalAmount += amount

	}
	fmt.Println("the total amount of my shopingList:", totalAmount)
}

package main

import "fmt"

func main() {
	fmt.Println("RANGES")

	names := []string{"Damas", "joseph", "Wambura"}
	fmt.Println("=== index & names ===")
	for i, name := range names {
		fmt.Println(i, name)

		for _, ch := range name {
			fmt.Printf("%q\n", ch)
		}

	}
}

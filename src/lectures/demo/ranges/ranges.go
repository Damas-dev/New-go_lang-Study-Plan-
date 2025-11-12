package main

import "fmt"

func main() {
	fmt.Println("this is range")

	names := []string{"Damas", "joseph", "range", "dovemica"}

	for i, name := range names {
		fmt.Println(i, name)

		for _, ch := range name {
			fmt.Printf(" %q\n", ch)
		}
	}
}

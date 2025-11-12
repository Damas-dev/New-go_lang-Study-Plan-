package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("----", title, "----")
	for i := 0; i < len(slice); i++ {
		item := slice[i]
		fmt.Println(item)
	}
}
func main() {

	route := []string{"a", "b", "c", "d"}

	printSlice("route1", route)

	route = append(route, "e")

	printSlice("route2", route)

	fmt.Println()
	fmt.Println(route[1], "visited")
	fmt.Println(route[2], "visited")

	route = route[2:]
	printSlice("remaining locations", route)

}

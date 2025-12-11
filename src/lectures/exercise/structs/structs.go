package main

import "fmt"

type Codinates struct {
	x, y int
}

type rectPoints struct {
	a, b Codinates
}

func width(points rectPoints) int {
	return points.b.y - points.a.y
}

func lenght(points rectPoints) int {
	return points.b.x - points.a.x
}

func area(points rectPoints) int {
	return width(points) * lenght(points)
}

func perimeter(points rectPoints) int {
	return width(points)*2 + lenght(points)*2
}

func printInfo(points rectPoints) {
	fmt.Println("===== Rectangle Area and Perimeter =====")
	fmt.Println("--- given point a(x,y) `top right` & point b(x,y) `bottom left` ---")
	fmt.Println("given", "a(", points.a.x, ",", points.a.y, ")  &  b(", points.b.x, ",", points.b.y, ")")

	fmt.Println("Area:", area(points))
	fmt.Println("Perimeter:", perimeter(points))

}

func main() {
	codPoints := rectPoints{a: Codinates{0, 7}, b: Codinates{4, 17}}

	printInfo(codPoints)
}

//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// ~ Create a rectangle structure containing its coordinates
type Rectangle struct {
	lenght float64
	width  float64
}

// ~ Using functions, calculate the area and perimeter of a rectangle,
//   - Print the results to the terminal
//   - The functions must use the rectangle structure as the function parameter
func rectangleArea(rec Rectangle) float64 {
	return (rec.lenght * rec.width)
}

func rectanglePerimeter(rec Rectangle) float64 {
	return (rec.lenght*2 + rec.width*2)
}

//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal

func doubleTheValue(rec Rectangle) (float64, float64) {
	return rec.lenght * 2, rec.width * 2

}

func main() {

	rectangle1 := Rectangle{3, 4}

	fmt.Println("Rectangle1 Area:", rectangleArea(rectangle1))

	fmt.Println("Rectangle1 Perimeter:", rectanglePerimeter(rectangle1))

	rectangle1.lenght, rectangle1.width = doubleTheValue(rectangle1)
	fmt.Println(rectangle1)

	fmt.Println("Rectangle1 Area:", rectangleArea(rectangle1))

	fmt.Println("Rectangle1 Perimeter:", rectanglePerimeter(rectangle1))

}

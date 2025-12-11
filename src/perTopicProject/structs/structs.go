package main

//* this project can now calculate the area of  all trinagles
//! it can only calulate the perimeter and hypotenuse of  a right angled ,isosceles tringle and equilateral triangle
//! it can neither  compute the  perimeter nor hypotenuse of any other irregular triagle

import (
	"fmt"
	"math"
)

type Codinates struct {
	x, y int
}

// struct of points that make up a triangle
type Tripoints struct {
	a, b, c Codinates
}

// height:= by - ay/ ay - by which over is bigger
func height(point Tripoints) int {

	a, b := point.a.y, point.b.y
	var height int

	//! if all codinates are negative
	if a < 0 && b < 0 {
		if a < 0 {
			a *= -1
		}
		if b < 0 {
			b *= -1
		}
		if a > b {
			height = a - b
		} else {
			height = b - a
		}
	} else if (a < 0 && b > 0) || (a > 0 && b < 0) { //! if one codinate is negative
		if a < 0 {
			a *= -1
		}
		if b < 0 {
			b *= -1
		}
		height = a + b

	} else { //! if all coddinate are positive
		if a > b {
			height = a - b
		} else {
			height = b - a
		}
	}

	return height
}

// bc is base:= bx - cx/ cx - bx which over is bigger
func base(point Tripoints) int {

	c, b := point.c.x, point.b.x
	var base int

	//! if all codinates are negative
	if c < 0 && b < 0 {
		if c < 0 {
			c *= -1
		}
		if b < 0 {
			b *= -1
		}
		if c > b {
			base = c - b
		} else {
			base = b - c
		}
	} else if (c < 0 && b > 0) || (c > 0 && b < 0) { //! if one codinate is negative
		if c < 0 {
			c *= -1
		}
		if b < 0 {
			b *= -1
		}
		base = c + b

	} else { //! if all coddinate are positive
		if c > b {
			base = c - b
		} else {
			base = b - c
		}
	}

	return base
}
func area(points Tripoints) float64 {
	h, b := float64(height(points)), float64(base(points))
	area := (h * b) / 2
	return area
}

func hypotenuse(points Tripoints) float64 {
	h := float64(height(points))
	b := float64(base(points))
	if points.a.x != points.b.x { //! if the triangle is no a right angled triangle
		b = b / 2
	}
	sqH := h * h // square of height
	sqB := b * b

	hypotenuse := math.Sqrt(sqB + sqH)

	return hypotenuse

}

func perimeter(points Tripoints) float64 {
	h := float64(height(points))
	b := float64(base(points))
	hypo := hypotenuse(points)
	var p float64                 //! p == perimeter
	if points.a.x == points.b.x { //! if a right angled triangle
		p = h + b + hypo
	} else {
		p = b + hypo*2
	}

	return p
}

func printTriangleInfo(points Tripoints) {
	fmt.Println("\n========================= Triangle Infomation =========================")

	fmt.Println("\nGiven .........")
	fmt.Println("a (", points.a.x, ",", points.a.y, ")\nb (", points.b.x, ",", points.b.y, ") \nc (", points.c.x, ",", points.c.y, ")")
	fmt.Println("---point a(x,y) , point b(x,y)  & point c(x,y) making a triangle---")

	fmt.Println("\noutput .........")
	fmt.Println("base:", base(points))
	fmt.Println("height:", height(points))
	fmt.Println("hypotenuse:", hypotenuse(points))
	fmt.Println("Perimeter:", perimeter(points))
	fmt.Println("Area:", area(points))

}

func main() {
	triangle1 := Tripoints{a: Codinates{0, 3}, b: Codinates{0, 0}, c: Codinates{4, 0}}

	printTriangleInfo(triangle1)

}

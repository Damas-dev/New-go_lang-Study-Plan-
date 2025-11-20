package main

import "fmt"

//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts// * Create a function to print out the contents of the assembly line

func printslice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		item := slice[i]
		fmt.Println(item)
	}
}

func main() {
	// 1. Create a slice with initial values
	marks := []int{10, 20, 30, 40, 50, 55, 88, 94}

	// 2. Add two more numbers
	marks = append(marks, 60, 70)

	// 3. Remove the 3rd element (value 30)
	index := 2 // 0-based index
	marks = append(marks[:index], marks[index+1:]...)

	// 4. Print final slice
	fmt.Println("Final slice:", marks)

	// Print  length
	fmt.Println("Length:", len(marks))

	// Print capacity
	fmt.Println("Capacity:", cap(marks))

	assemblyLine := []string{"p1", "p2", "p3"}

	printslice("assembly line1", assemblyLine)

	//~Add two new parts to the line

	assemblyLine = append(assemblyLine, "p4", "p5")

	printslice("new assembly line", assemblyLine)

	//~Slice the assembly line so it contains only the two new parts

	assemblyLine = assemblyLine[3:]
	printslice("Unassembled parts", assemblyLine)

}

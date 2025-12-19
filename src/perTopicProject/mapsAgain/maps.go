//--Summary:
//  Write a program to manage student exam marks in a class.
//
//--Scenario:
//  A teacher wants to store students and their exam marks,
//  then calculate useful information from them.
//
//--Requirements:
//* Store student names and their marks in a map
//* Add at least 5 students with different marks
//* Create a function to:
//  - Print all students and their marks
//  - Calculate and print the class average
//  - Find and print:
//      • Highest mark and the student who scored it
//      • Lowest mark and the student who scored it
//* Update a student’s mark and show the updated results

package main

import "fmt"

// * Create a function to:
//   - Print all students and their marks
//   - Calculate and print the class average
//   - Find and print:
//   - Highest mark and the student who scored it
//   - Lowest mark and the student who scored it
func printReport(students map[string]int) {
	total := 0
	count := len(students)

	highestMark := -1
	lowestMark := 101

	var topStudent string
	var lowStudent string

	fmt.Println("Student Marks...")
	for name, mark := range students {
		fmt.Println(name, ":", mark)

		total += mark

		if mark > highestMark {
			highestMark = mark
			topStudent = name
		}

		if mark < lowestMark {
			lowestMark = mark
			lowStudent = name
		}
	}

	average := float64(total) / float64(count)

	fmt.Println("\nClass Average:", average)
	fmt.Println("Highest Mark:", highestMark, "by", topStudent)
	fmt.Println("Lowest Mark:", lowestMark, "by", lowStudent)
	fmt.Println("--------------------------------")
}

func main() {
	//* Store student names and their marks in a map
	//* Add at least 5 students with different marks
	students := map[string]int{
		"Adroit": 78,
		"Dajo":   92,
		"John":   65,
		"Mama":   85,
		"Damas":  90,
	}

	// Display initial report
	printReport(students)

	//* Update a student’s mark and show the updated results
	fmt.Println("\nUpdating John's mark...")
	students["John"] = 75

	// Display updated report
	printReport(students)
}

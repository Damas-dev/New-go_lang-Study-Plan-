package main

import (
	"fmt"
)

func max(slice []int) int {
	max := slice[0]
	for _, value := range slice {
		if value >= max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println("this is the  maps demo")

	// a map of school to it's respective classes
	school := make(map[string]int)

	school["form1a"] = 36
	school["form1b"] = 32
	school["form2a"] = 36
	school["form2b"] = 36

	school["form3a"] = 36
	school["form3b"] = 30
	school["form4a"] = 31
	school["form4b"] = 10

	fmt.Println(school)
	// keep all form4 in one class
	school["form4a"] += school["form4b"]

	delete(school, "form4b")
	fmt.Println(school)

	school["form4"] = school["form4a"]
	delete(school, "form4a")
	for class, total := range school {
		fmt.Println(class, total)
	}

	// finding a class(form4a)
	form4a, found := school["form4a"]
	fmt.Print("is their form4a? ")
	if found {
		fmt.Println("they are", form4a, "in form4a")
	} else {
		fmt.Println("nope")
	}
	var totalStudentInSchool = 0
	for _, numStudentInclass := range school {
		totalStudentInSchool += numStudentInclass
	}

	fmt.Println("their are", totalStudentInSchool, " students in school")

	listStudents := []int{}

	for _, total := range school {
		listStudents = append(listStudents, total)
	}
	fmt.Println(listStudents)

	for class, total := range school {
		if total == max(listStudents) {
			fmt.Println("class:", class, "has the highest number of students:", total)
		}
	}

}

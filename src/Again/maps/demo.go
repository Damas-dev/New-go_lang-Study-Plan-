package main

import (
	"fmt"
)

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
	fmt.Println(school)

	// finding a class
	form4a, found := school["form4a"]

	fmt.Println("=== are their students===")
	if found {
		fmt.Println("their are", form4a, "in form4a")
	} else {
		fmt.Println("nope")
	}
	var totalstudent = 0
	for _, noStudent := range school {
		totalstudent += noStudent
	}

	fmt.Println("their are", totalstudent, " students in school")

}

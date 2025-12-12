package main

import "fmt"

type Student struct {
	name  string
	score int
}

func bestStudents(students [10]Student) {
	max := students[0].score
	for i := 1; i < len(students); i++ {
		std := students[i]
		if std.score >= max {
			max = std.score
		}
	}
	fmt.Println("\nBest Student(s).........")
	for j := 0; j < len(students); j++ {
		std := students[j]
		if std.score == max {

			fmt.Println("Name:", std.name, "Scored:", std.score)

		}

	}
	fmt.Println("..........................................")

}

func classStatus(students [10]Student) {
	var avg, total, classAvg int
	var totalSub int //! toal students submited
	allStudents := len(students)
	for i := 0; i < allStudents; i++ {
		std := students[i]
		total += std.score

		if std.name != "" {
			totalSub++
		}
	}
	avg = total / totalSub //! this is the current average of the stunded stubmited
	classAvg = total / allStudents

	fmt.Println("\nClass Resurts.........")
	fmt.Println(totalSub, "Students Submited the exam")
	fmt.Println(allStudents-totalSub, "Students Didn't Submit")
	fmt.Println("Current Average:", avg)
	fmt.Println("Class Average:", classAvg)
	fmt.Println("..........................................")

}
func printResults(students [10]Student) {
	classStatus(students)
	bestStudents(students)

}

func main() {
	students := [10]Student{
		{"Damas", 95},
		{"Dan", 80},
		{"Domi", 75},
		{"Dog", 70},
		{"Dante", 65},
		{"Jose", 95},
	}
	printResults(students)
	fmt.Println("\n......remaing 4 student submited......")
	students[6] = Student{"Dajo", 95}
	students[7] = Student{"Doja", 85}
	students[8] = Student{"Dj", 82}
	students[9] = Student{"Jj", 73}

	printResults(students)
}

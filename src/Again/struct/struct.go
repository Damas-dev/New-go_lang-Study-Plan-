package main

import "fmt"

// structure for student and his marks
type Student struct {
	firstName string
	lastName  string
	math      int
	english   int
	kiswahili int
	chemistry int
}

func average(stud Student) int {
	sum := stud.math + stud.english + stud.kiswahili + stud.chemistry

	return sum / 4

}

func status(stud Student) {

	switch av := average(stud); {
	case av >= 80:
		fmt.Printf("Name%q%q\nAverage:%c\nGread:A\nStatus:PASS", stud.firstName, stud.lastName, av)
	case av >= 65:
		fmt.Printf("Name%q%q\nAverage:%c\nGread:B\nStatus:PASS", stud.firstName, stud.lastName, av)
	case av >= 45:
		fmt.Printf("Name%q%q\nAverage:%c\nGread:C\nStatus:PASS", stud.firstName, stud.lastName, av)
	case av >= 30:
		fmt.Printf("Name%q%q\nAverage:%c\nGread:D\nStatus:Supplementary ", stud.firstName, stud.lastName, av)

	default:
		fmt.Printf("Name%q%q\nAverage:%c\nGread:F\nStatus:Discontinue", stud.firstName, stud.lastName, av)
	}
}

func main() {
	fmt.Println("this is some stract")

	damas := Student{
		firstName: "Damas",
		lastName:  "Wambura",
		math:      85,
		english:   85,
		kiswahili: 85,
		chemistry: 85,
	}

	status(damas)
}

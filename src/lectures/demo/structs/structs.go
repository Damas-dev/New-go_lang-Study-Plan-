package main

import "fmt"

type Student struct {
	name     string
	id       int
	hasAbook bool
}
type Libary struct {
	inLibary Student
}

func main() {

	damas := Student{"Damas", 1, false}

	fmt.Println(damas)

	var (
		jose = Student{"Joseph", 2, false}
		win  = Student{"Winfrida", 3, false}
	)

	fmt.Println(jose, win)

	jose.hasAbook = true
	win.hasAbook = true

	if jose.hasAbook {
		fmt.Println(jose.name, "has a book form the library")
	} else {
		fmt.Println(jose.name, "has no book form the library")

	}
	if win.hasAbook {
		fmt.Println(win.name, "has a book form the library")
	} else {
		fmt.Println(win.name, "has no book form the library")
	}
	if damas.hasAbook {
		fmt.Println(damas.name, "has a book form the library")
	} else {
		fmt.Println(damas.name, "has no book form the library")
	}

	lib1 := Libary{damas}

	fmt.Println(lib1.inLibary.name, "is in libray1")

}

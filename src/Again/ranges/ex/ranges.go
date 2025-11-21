package main

import "fmt"

func main() {
	fmt.Println("===== This is sign Game =====")

	//~ names := []string{"Damas", "joseph", "Wambura"}

	fName := "Damas"
	mName := "joseph"
	sirName := "Wambura"

	names := make([]string, 0, 3)

	names = append(names, fName, mName, sirName)

	fmt.Println(names)

	//! fmt.Println(names[0][1], names[1][1], names[2]) //= 97 111 wambura

	//~ for _, ch := range names[0] {
	//~ 	fmt.Printf("%q", ch)
	//~ 	break
	//~ }

	//~ for _, ch := range names[1] {
	//~ 	fmt.Printf("%q", ch)
	//~ 	break
	//~ }

	//~ for _, ch := range names[2] {
	//~ 	fmt.Printf("%q", ch)

	//~ }

	for i, name := range names {
		for _, ch := range name {
			fmt.Printf("%q", ch)
			if i == 0 || i == 1 {
				break
			}
		}
	}

}

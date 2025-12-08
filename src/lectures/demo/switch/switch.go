package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap item")

		fallthrough // allow the exechution of the next case
	case p < 10:
		fmt.Println("Modelete price")
	case p < 100:
		fmt.Println("Expenceve item")
	default:
		fmt.Println("Very expenceve item")

	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy sitting")
	case Business:
		fmt.Println("Business sitting")
	case FirstClass:
		fmt.Println("first Class sitting")
	default:
		fmt.Println("other sittings")
	}
}

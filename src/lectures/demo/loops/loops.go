package main

import "fmt"

func main() {

	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum:", sum)
	}

	for sum >= 10 {
		fmt.Println("<Sum:", sum)
		sum -= 5

	}

}

package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	//!q = quiz

	q1, q2, q3 := 7, 8, 9

	const passMark float32 = 7

	if q1 > q2 && q1 > q3 {
		fmt.Printf("you scored the highest in quiz1: %v\n", q1)
	} else if q2 > q3 {
		fmt.Printf("you scored the highest in quiz2: %v\n", q2)
	} else {
		fmt.Printf("you scored the highest in quiz3: %v\n", q3)

	}

	// cheak if the student has passed or not

	if average(q1, q2, q3) >= passMark {
		fmt.Println("status: PASS")
	} else {
		fmt.Println("status: FAIL")
	}

}

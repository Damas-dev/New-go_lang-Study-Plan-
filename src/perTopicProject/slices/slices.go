package main

import "fmt"

// * Ascending order using remove and append
func printAscending(nums []int) {
	var sorted []int

	for len(nums) > 0 {
		minIndex := 0

		//~ find smallest element
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}

		//~ add smallest to new slice
		sorted = append(sorted, nums[minIndex])

		//~ remove smallest from original slice
		nums = append(nums[:minIndex], nums[minIndex+1:]...)
	}

	// print result
	for i := 0; i < len(sorted); i++ {
		fmt.Print(sorted[i], " ")

	}
	fmt.Println("\n------------")
}

// * Descending order using remove and append logic
func printDescending(nums []int) {
	var sorted []int

	for len(nums) > 0 {
		maxIndex := 0

		//~ find largest element
		for i := 1; i < len(nums); i++ {
			if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
		}

		//~ add largest to new slice
		sorted = append(sorted, nums[maxIndex])

		//~ remove largest from original slice
		nums = append(nums[:maxIndex], nums[maxIndex+1:]...)
	}

	// print result
	for i := 0; i < len(sorted); i++ {
		fmt.Print(sorted[i], " ")
	}
	fmt.Println("\n------------")
}

func main() {
	arr := [6]int{5, 2, 9, 1, 6, 3}
	nums := arr[:]

	fmt.Println("Ascending order...")
	printAscending(nums)

	// reset array
	arr = [6]int{5, 2, 9, 1, 6, 3}
	nums = arr[:]

	fmt.Println("Descending order...")
	printDescending(nums)

}

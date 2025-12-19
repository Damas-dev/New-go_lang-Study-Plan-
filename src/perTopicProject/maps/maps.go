//--Summary:
//  Write a program to track user account statuses.
//
//--Requirements:
//* Create a function to print account statistics, including:
//  - Total number of users
//  - Number of users for each status (Active, Suspended, Banned, Pending)
//
//* Store the existing slice of usernames in a map
//* Default all users to `Active`
//
//* Perform the following status changes and display account info:
//  - display account info
//  - change `john` to `Suspended`
//  - change `alice` to `Banned`
//  - display account info
//  - change all users to `Pending`
//  - display account info

package main

import "fmt"

// user status
const (
	Active    = "-Active"
	Suspended = "-Suspended"
	Banned    = "-Banned"
	Pending   = "-Pending"
)

//* Create a function to print account statistics, including:
//  - Total number of users
//  - Number of users for each status (Active, Suspended, Banned, Pending)

func printaccountStatistics(users map[string]string) {
	fmt.Println("#users:", len(users))
	for name, status := range users {
		fmt.Println(name, status)

	}
	stat := make(map[string]int)
	for _, status := range users {
		switch status {
		case Active:
			stat[Active] += 1
		case Suspended:
			stat[Suspended] += 1
		case Banned:
			stat[Banned] += 1
		case Pending:
			stat[Pending] += 1
		default:
			fmt.Println("panic: error")
		}
	}
	fmt.Println(stat[Active], "Users are Active")
	fmt.Println(stat[Suspended], "Users are Suspended")
	fmt.Println(stat[Banned], "Users are Banned")
	fmt.Println(stat[Pending], "Users Pending")

}

func main() {
	users := []string{"Damas", "alice", "Adroit", "john", "Octo", "piko"}

	// user status map
	usersStatus := make(map[string]string)
	//* Store the existing slice of usernames in a map
	//* Default all users to `Active`
	for _, name := range users {
		usersStatus[name] = Active
	}

	//* Perform the following status changes and display account info:
	//  - display account info
	printaccountStatistics(usersStatus)
	//  - change `john` to `Suspended`
	usersStatus["john"] = Suspended
	//  - change `alice` to `Banned`
	usersStatus["alice"] = Banned

	//  - display account info
	printaccountStatistics(usersStatus)
	//  - change all users to `Pending`
	for name := range usersStatus {
		usersStatus[name] = Pending
	}
	//  - display account info
	printaccountStatistics(usersStatus)

}

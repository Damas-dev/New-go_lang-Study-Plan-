package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func cheakIfCleaned(rooms [5]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println("Room", room.name, "is cleaned")
		} else {
			fmt.Println("Room", room.name, "is Not cleaned")
		}
	}
}

func main() {

	rooms := [...]Room{
		{name: "Office1"},
		{name: "Office2"},
		{name: "Office3"},
		{name: "Office4"},
		{name: "Office5"},
	}

	cheakIfCleaned(rooms)

	fmt.Println("After cleaning........")
	rooms[1].cleaned = true
	rooms[3].cleaned = true
	rooms[0].cleaned = true

	cheakIfCleaned(rooms)

}

package main

import "fmt"

type Counter struct {
	hits int
}

func incremet(counter *Counter) {
	counter.hits += 1
	fmt.Println("counter:", counter.hits)
}

func replaceString(old *string, new string, count *Counter) {
	*old = new
	incremet(count)
}

func main() {

	count := Counter{}

	hello := "hello"
	world := "world!"
	name := "Damas"

	fmt.Println(hello, world)

	replaceString(&hello, "hi", &count)

	fmt.Println(hello, world)

	phrase := []string{hello, world}

	replaceString(&phrase[1], name, &count)

	fmt.Println(phrase)

}

package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var people []Person
	//people := []Person{}
	people = make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		p := Person{
			"Gus",
			"Gomez",
			13,
		}
		people = append(people, p)
	}
}

package main

import "fmt"

func prefixer(input string) func(string) string {
	return func(in string) string {
		return input + " " + in
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

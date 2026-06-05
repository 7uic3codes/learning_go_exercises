package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	// The program wants to add from 0-9
	// but we declare total and reassign in the block of for
	// so it is just the i variable assigned to total
}

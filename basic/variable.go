package main

import "fmt"

func main() {
	name := "a"
	fmt.Println(name)

	name = "b"
	fmt.Println(name)

	var (
		firstName = "Es"
		lastName = "Teh"
	)

	fmt.Println(firstName, lastName)
}
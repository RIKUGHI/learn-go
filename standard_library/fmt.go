package main

import "fmt"

func main() {
	firstName := "Bambang"
	lastName := "Cahyo"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Bambang"
	// middleName = "Cahya"
	lastName = "Di"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
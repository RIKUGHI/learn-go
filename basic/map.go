package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Bambang"
	// person["address"] = "Jakarta"

	person := map[string]string{
		"name": "Bambang",
		"address": "Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(person["nter"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Bambang"
	book["nter"] = "salah"

	fmt.Println(book)
	delete(book, "nter")
	fmt.Println(book)
}
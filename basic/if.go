package main

import "fmt"

func main() {
	name := "Bambang"

	if name == "Bambang" {
		fmt.Println("Hello Bambang")
		} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi")
	}

	if length := len(name); length>5 {
		
		fmt.Println("Telalu panjang", length)
	} else {
		fmt.Println("Nama benar")
	}
}
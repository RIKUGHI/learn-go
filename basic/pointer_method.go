package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bambang := Man{"Bambang"}
	bambang.Married()

	fmt.Println(bambang.Name)
}
package main

import (
	"basic/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Bambang")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Hi")) // tidak bisa diakses
}
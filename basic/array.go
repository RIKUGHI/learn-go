package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "a"
	names[1] = "b"
	names[2] = "c"

	fmt.Println(names[0], names[1], names[2])
	
	var values = [...]int{
		1,2,1,5,6,7,
	}

	fmt.Println(values, len(values))

	var test = [3]string{"a","b","c"}
	
	fmt.Println(test)
}
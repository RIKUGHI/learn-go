package main

import "fmt"

func main() {
	type NoKTP string

	var ktpMe NoKTP = "aaa"

	var contoh string = "bbb"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpMe)
	fmt.Println(contohKtp)
}
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Bambang Cahyo", "Bambang"))
	fmt.Println(strings.Split("Bambang Cahyo", " "))
	fmt.Println(strings.ToLower("Bambang Cahyo"))
	fmt.Println(strings.ToUpper("Bambang Cahyo"))
	fmt.Println(strings.Trim("      Bambang Cahyo      ", " "))
	fmt.Println(strings.ReplaceAll("Bambang Cahyo Bambang Cahyo", "Bambang", "Budi"))
}
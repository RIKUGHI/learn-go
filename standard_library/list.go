package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Bambang")
	data.PushBack("Cahyo")
	data.PushBack("Nter")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // bambang

	next := head.Next() // cahyo
	fmt.Println(next.Value)

	next = next.Next() // nter
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
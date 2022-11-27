package main

import (
	"data_structures/linked_list"
	"fmt"
)

func main() {
	list := &linked_list.SingleLinkedList[int]{}
	list.AddFirst(2)
	list.AddFirst(1)
	fmt.Println("Hello World")
	fmt.Println(list.Len())
    fmt.Println(list)
}

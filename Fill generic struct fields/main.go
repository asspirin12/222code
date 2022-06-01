package main

import "fmt"

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func main() {
	node := Node[any]{}
	fmt.Println(node)
}

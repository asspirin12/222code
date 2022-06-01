package main

type T[A any, B any] struct {
	field A
}

func main() {
	var _ T[]    // List of type parameters must not be empty
	var _ T[int] // Got 1 arguments but 2 type parameters
	var _ T[int, string]
	var _ T[int, string, int] // Got 3 arguments but 2 type parameters
}

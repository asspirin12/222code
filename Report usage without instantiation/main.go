package main

type Foo[T any] interface {
	Bork() T
}

type Bar interface {
	GetFoo() Foo
}

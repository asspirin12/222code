package main

import (
	"fmt"
	"io"
)

func handle[T io.Closer | Flusher](t T) {
	err := t.Flush()
	if err != nil {
		fmt.Println("failed to flush: ", err.Error())
	}

	err = t.Close()
	if err != nil {
		fmt.Println("failed to close: ", err.Error())
	}
}

type Flusher interface {
	Flush() error
}
